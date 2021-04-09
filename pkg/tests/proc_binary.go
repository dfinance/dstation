package tests

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"path"
	"strings"
	"time"
)

const (
	DvmBinaryStartTimeout = 2 * time.Second
)

// BinaryCmdOption defines functional arguments for NewBinaryCmd.
type BinaryCmdOption func(*BinaryCmd) error

// BinaryCmd keeps process metadata.
type BinaryCmd struct {
	cmd          string
	args         []string
	proc         *exec.Cmd
	printLogs    bool
}

func (c *BinaryCmd) String() string {
	return fmt.Sprintf("binary %s %s", c.cmd, strings.Join(c.args, " "))
}

// Start starts a new process in background.
func (c *BinaryCmd) Start() error {
	if c.proc != nil {
		return fmt.Errorf("%s: process already started", c)
	}
	c.proc = exec.Command(c.cmd, c.args...)

	if c.printLogs {
		stdoutPipe, err := c.proc.StdoutPipe()
		if err != nil {
			return fmt.Errorf("%s: getting stdout pipe: %w", c, err)
		}
		stderrPipe, err := c.proc.StderrPipe()
		if err != nil {
			return fmt.Errorf("%s: getting stderr pipe: %w", c, err)
		}

		c.startLogWorker(FmtInfColorPrefix, stdoutPipe)
		c.startLogWorker(FmtWrnColorPrefix, stderrPipe)
	}

	if err := c.proc.Start(); err != nil {
		return fmt.Errorf("%s: starting process: %w", c, err)
	}

	return nil
}

// Stop sends SIGKILL to a running process.
func (c *BinaryCmd) Stop() error {
	if c.proc == nil {
		return nil
	}

	if err := c.proc.Process.Kill(); err != nil {
		return fmt.Errorf("%s: stop failed: %w", c, err)
	}

	return nil
}

// startLogWorker starts the logs pipe logger (to the Host stdout).
func (c *BinaryCmd) startLogWorker(msgFmtPrefix string, pipe io.ReadCloser) {
	msgFmt := msgFmtPrefix + "%s: %s" + FmtColorEndLine
	buf := bufio.NewReader(pipe)

	go func() {
		for {
			line, _, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					return
				}

				fmt.Printf("%s: broken pipe: %v", c, err.Error())
				return
			}

			fmt.Printf(msgFmt, path.Base(c.cmd), line)
		}
	}()
}

// NewBinaryCmd creates a new BinaryCmd (no Start).
func NewBinaryCmd(cmd string, options ...BinaryCmdOption) (*BinaryCmd, error) {
	c := &BinaryCmd{
		cmd:     cmd,
	}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, fmt.Errorf("%s: option apply failed: %w", c, err)
		}
	}

	return c, nil
}

// BinaryWithArgs sets process arguments.
func BinaryWithArgs(args ...string) BinaryCmdOption {
	return func(c *BinaryCmd) error {
		c.args = args
		return nil
	}
}

// BinaryWithConsoleLogs redirects process stdout / stderr to the current stdout.
func BinaryWithConsoleLogs(enabled bool) BinaryCmdOption {
	return func(c *BinaryCmd) error {
		c.printLogs = enabled
		return nil
	}
}

// NewBinaryDVMWithNetTransport creates and starts a new DVM BinaryCmd with TCP connection.
func NewBinaryDVMWithNetTransport(basePath, connectPort, dsServerPort string, printLogs bool, args ...string) (*BinaryCmd, error) {
	cmdArgs := []string{
		"http://127.0.0.1:" + connectPort,
		"http://127.0.0.1:" + dsServerPort,
	}
	cmdArgs = append(cmdArgs, args...)

	c, err := NewBinaryCmd(path.Join(basePath, "dvm"), BinaryWithArgs(cmdArgs...), BinaryWithConsoleLogs(printLogs))
	if err != nil {
		return nil, err
	}

	if err := c.Start(); err != nil {
		return nil, err
	}
	time.Sleep(DvmBinaryStartTimeout)

	return c, nil
}

// NewBinaryDVMWithUDSTransport creates and starts a new DVM BinaryCmd with UDS connection.
func NewBinaryDVMWithUDSTransport(basePath, socketsDir, connectSocketName, dsSocketName string, printLogs bool, args ...string) (*BinaryCmd, error) {
	cmdArgs := []string{
		"ipc:/" + path.Join(socketsDir, connectSocketName),
		"ipc:/" + path.Join(socketsDir, dsSocketName),
	}
	cmdArgs = append(cmdArgs, args...)

	c, err := NewBinaryCmd(path.Join(basePath, "dvm"), BinaryWithArgs(cmdArgs...), BinaryWithConsoleLogs(printLogs))
	if err != nil {
		return nil, err
	}

	if err := c.Start(); err != nil {
		return nil, err
	}

	if err := WaitForFileExists(path.Join(socketsDir, connectSocketName), DvmBinaryStartTimeout); err != nil {
		return nil, fmt.Errorf("%s: waiting for UDS server to start-up: %v", c.Start(), err)
	}

	return c, nil
}
