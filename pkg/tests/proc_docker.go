package tests

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	docker "github.com/fsouza/go-dockerclient"
)

const (
	DvmDockerStartTimeout = 5 * time.Second
)

// DockerContainerOption defines functional arguments for NewDockerContainer.
type DockerContainerOption func(*DockerContainer) error

// DockerContainer keeps Docker container metadata.
type DockerContainer struct {
	dClient          *docker.Client
	dContainer       *docker.Container
	dOptions         docker.CreateContainerOptions
	printLogs        bool
	logsStreamStopCh chan bool
}

func (c *DockerContainer) String() string {
	return "container " + c.dOptions.Config.Image
}

// Start starts a new container with timeout.
func (c *DockerContainer) Start(startTimeout time.Duration) error {
	if c.dClient != nil {
		return fmt.Errorf("%s: already started", c.String())
	}

	client, err := docker.NewClientFromEnv()
	if err != nil {
		return fmt.Errorf("%s: connecting to docker: %w", c.String(), err)
	}

	container, err := client.CreateContainer(c.dOptions)
	if err != nil {
		return fmt.Errorf("%s: creating container: %w", c.String(), err)
	}

	if err := client.StartContainer(container.ID, nil); err != nil {
		return fmt.Errorf("%s: starting container: %w", c.String(), err)
	}

	if c.printLogs {
		var stdoutBuf, stderrBuf bytes.Buffer
		opts := docker.LogsOptions{
			Container:    container.ID,
			OutputStream: &stdoutBuf,
			ErrorStream:  &stderrBuf,
			Follow:       true,
			Stdout:       true,
			Stderr:       true,
		}

		c.logsStreamStopCh = make(chan bool)
		go func() {
			if err := client.Logs(opts); err != nil {
				fmt.Printf("%s: logs worker: %v\n", c.String(), err)
			}
		}()

		c.startLogWorker(FmtInfColorPrefix, &stdoutBuf, c.logsStreamStopCh)
		c.startLogWorker(FmtWrnColorPrefix, &stderrBuf, c.logsStreamStopCh)
	}

	// Wait for container to start
	timeoutCh := time.NewTimer(startTimeout).C
	for {
		inspectContainer, err := client.InspectContainerWithOptions(docker.InspectContainerOptions{ID: container.ID})
		if err != nil {
			return fmt.Errorf("%s: wait for container to start: %w", c, err)
		}
		if inspectContainer.State.Running {
			break
		}

		select {
		case <-timeoutCh:
			return fmt.Errorf("%s: wait for container to start: timeout reached (%v)", c, startTimeout)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Wait for all TCP port to be reachable
	portReports := make(map[docker.Port]string)
	for p := range c.dOptions.Config.ExposedPorts {
		portReports[p] = "not checked"
	}
	for {
		cnt := len(portReports)
		for p, status := range portReports {
			if status == "OK" {
				cnt--
				continue
			}

			if err := PingTcpAddress("127.0.0.1:"+p.Port(), 500*time.Millisecond); err != nil {
				portReports[p] = err.Error()
			} else {
				portReports[p] = "OK"
				cnt--
				continue
			}

			select {
			case <-timeoutCh:
				reports := make([]string, 0, len(portReports))
				for p, status := range portReports {
					reports = append(reports, fmt.Sprintf("%s: %s", p.Port(), status))
				}

				return fmt.Errorf(
					"%s: wait for container TCP ports to be rechable: timeout reached (%v): %s",
					c, startTimeout, strings.Join(reports, ", "),
				)
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
		if cnt == 0 {
			break
		}
	}

	c.dClient = client
	c.dContainer = container

	return nil
}

// Stop stops a running container.
func (c *DockerContainer) Stop() error {
	if c.dClient == nil {
		return fmt.Errorf("%s: not started", c)
	}

	if c.logsStreamStopCh != nil {
		close(c.logsStreamStopCh)
	}
	err := c.dClient.RemoveContainer(docker.RemoveContainerOptions{
		ID:    c.dContainer.ID,
		Force: true,
	})
	if err != nil {
		return fmt.Errorf("%s: removing container: %w", c, err)
	}

	return nil
}

// startLogWorker starts the logs buffer logger (to the Host stdout).
func (c *DockerContainer) startLogWorker(msgFmtPrefix string, buf *bytes.Buffer, stopCh chan bool) {
	logCh := make(chan string)
	msgFmt := msgFmtPrefix + "%s: %s" + FmtColorEndLine

	go func() {
		for msg := range logCh {
			fmt.Printf(msgFmt, c, msg)
		}
	}()

	go func() {
		defer close(logCh)

		for {
			select {
			case <-stopCh:
				return
			default:
				line, err := buf.ReadString('\n')
				if err != nil && err != io.EOF {
					return
				}

				line = strings.TrimSpace(line)
				if line != "" {
					logCh <- line
				}
			}
		}
	}()
}

// NewDockerContainer creates a new DockerContainer object (no Start).
func NewDockerContainer(options ...DockerContainerOption) (*DockerContainer, error) {
	c := DockerContainer{}

	c.dOptions = docker.CreateContainerOptions{
		Config:     &docker.Config{},
		HostConfig: &docker.HostConfig{},
	}

	for _, options := range options {
		if err := options(&c); err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// DockerWithCreds sets container registry, name and tag.
func DockerWithCreds(registry, name, tag string) DockerContainerOption {
	return func(c *DockerContainer) error {
		c.dOptions.Config.Image = fmt.Sprintf("%s/%s:%s", registry, name, tag)
		return nil
	}
}

// DockerWithCmdArgs sets container cmd arguments.
func DockerWithCmdArgs(cmdArgs []string) DockerContainerOption {
	return func(c *DockerContainer) error {
		c.dOptions.Config.Cmd = cmdArgs
		return nil
	}
}

// DockerWithVolume mounts Docker volume to a container.
func DockerWithVolume(hostPath, containerPath string) DockerContainerOption {
	return func(c *DockerContainer) error {
		c.dOptions.HostConfig.VolumeDriver = "bind"
		c.dOptions.HostConfig.Binds = append(
			c.dOptions.HostConfig.Binds,
			fmt.Sprintf("%s:%s", hostPath, containerPath),
		)

		return nil
	}
}

// DockerWithTcpPorts bridges TCP ports from the Host to a container.
func DockerWithTcpPorts(tcpPorts []string) DockerContainerOption {
	return func(c *DockerContainer) error {
		ports := make(map[docker.Port]struct{}, len(tcpPorts))
		portBindings := make(map[docker.Port][]docker.PortBinding, len(tcpPorts))
		for _, p := range tcpPorts {
			dPort := docker.Port(fmt.Sprintf("%s/tcp", p))

			ports[dPort] = struct{}{}
			portBindings[dPort] = []docker.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: dPort.Port(),
				},
			}
		}

		c.dOptions.Config.ExposedPorts = ports
		c.dOptions.HostConfig.PortBindings = portBindings

		return nil
	}
}

// DockerWithHostNetwork sets the Host address and network mode.
func DockerWithHostNetwork() DockerContainerOption {
	return func(c *DockerContainer) error {
		_, mode, err := HostMachineDockerUrl()
		if err != nil {
			return err
		}

		c.dOptions.HostConfig.NetworkMode = mode

		return nil
	}
}

// DockerWithUser sets custom UID and GID for a container.
func DockerWithUser() DockerContainerOption {
	return func(c *DockerContainer) error {
		userUid, userGid := os.Getuid(), os.Getgid()
		if userUid < 0 {
			return fmt.Errorf("invalid user UID: %d", userUid)
		}
		if userGid < 0 {
			return fmt.Errorf("invalid user GID: %d", userGid)
		}

		c.dOptions.Config.User = fmt.Sprintf("%d:%d", userUid, userGid)

		return nil
	}
}

// DockerWithConsoleLogs enables container logs print to the Host stdout.
func DockerWithConsoleLogs(enabled bool) DockerContainerOption {
	return func(c *DockerContainer) error {
		c.printLogs = enabled
		return nil
	}
}

// NewDockerDVMWithNetTransport creates and starts a new DVM DockerContainer with TCP connection.
func NewDockerDVMWithNetTransport(registry, tag, connectPort, dsServerPort string, printLogs bool, args ...string) (*DockerContainer, error) {
	if registry == "" || tag == "" {
		return nil, fmt.Errorf("registry / tag: not specified")
	}

	hostUrl, _, _ := HostMachineDockerUrl()
	dsServerAddress := fmt.Sprintf("%s:%s", hostUrl, dsServerPort)
	cmdArgs := []string{"./dvm", "http://0.0.0.0:" + connectPort, dsServerAddress}
	if len(args) > 0 {
		cmdArgs = append(cmdArgs, strings.Join(args, " "))
	}

	container, err := NewDockerContainer(
		DockerWithCreds(registry, "dfinance/dvm", tag),
		DockerWithCmdArgs(cmdArgs),
		DockerWithTcpPorts([]string{connectPort}),
		DockerWithHostNetwork(),
		DockerWithConsoleLogs(printLogs),
	)
	if err != nil {
		return nil, fmt.Errorf("creating DVM container over Net: %v", err)
	}

	if err := container.Start(DvmDockerStartTimeout); err != nil {
		return nil, fmt.Errorf("starting DVM container over Net: %v", err)
	}

	return container, nil
}

// NewDockerDVMWithUDSTransport creates and starts a new DVM DockerContainer with UDS connection.
func NewDockerDVMWithUDSTransport(registry, tag, volumePath, vmFileName, dsFileName string, printLogs bool, args ...string) (*DockerContainer, error) {
	const defVolumePath = "/tmp/dn-uds"

	if registry == "" || tag == "" {
		return nil, fmt.Errorf("registry / tag: not specified")
	}

	vmFilePath := path.Join(defVolumePath, vmFileName)
	dsFilePath := path.Join(defVolumePath, dsFileName)

	// one '/' is omitted on purpose
	cmdArgs := []string{"./dvm", "ipc:/" + vmFilePath, "ipc:/" + dsFilePath}
	if len(args) > 0 {
		cmdArgs = append(cmdArgs, strings.Join(args, " "))
	}

	container, err := NewDockerContainer(
		DockerWithCreds(registry, "dfinance/dvm", tag),
		DockerWithCmdArgs(cmdArgs),
		DockerWithVolume(volumePath, defVolumePath),
		DockerWithUser(),
		DockerWithConsoleLogs(printLogs),
	)
	if err != nil {
		return nil, fmt.Errorf("creating DVM container over UDS: %v", err)
	}

	if err := container.Start(DvmDockerStartTimeout); err != nil {
		return nil, fmt.Errorf("starting DVM container over UDS: %v", err)
	}

	if err := WaitForFileExists(path.Join(volumePath, vmFileName), DvmDockerStartTimeout); err != nil {
		return nil, fmt.Errorf("creating DVM container over UDS: %v", err)
	}

	return container, nil
}

// HostMachineDockerUrl returns the Host address and network mode (OS-dependant).
func HostMachineDockerUrl() (hostUrl, hostNetworkMode string, err error) {
	switch runtime.GOOS {
	case "darwin", "windows":
		hostUrl, hostNetworkMode = "http://host.docker.internal", ""
	case "linux":
		hostUrl, hostNetworkMode = "http://localhost", "host"
	default:
		err = fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return
}
