package tests

import (
	"fmt"
	"os"
)

const (
	EnvDvmIntegUse            = "DN_DVM_INTEG_TESTS_USE"             // defines which DVM runner is used: "binary", "docker","" (not integ tests)
	EnvDvmIntegDockerRegistry = "DN_DVM_INTEG_TESTS_DOCKER_REGISTRY" // Docker runner: registry
	EnvDvmIntegDockerTag      = "DN_DVM_INTEG_TESTS_DOCKER_TAG"      // Docker runner: DVM tag
	EnvDvmIntegBinaryPath     = "DN_DVM_INTEG_TESTS_BINARY_PATH"     // Binary runner: directory path containing DVM binary (dvm), if empty - $PATH is used
	//
	EnvDvmIntegUseDocker = "docker"
	EnvDvmIntegUseBinary = "binary"
	//
	TestErrFmt = "Launching DVM over %s with %s transport"
)

func LaunchDVMWithNetTransport(connectPort, dsServerPort string, printLogs bool, args ...string) (stopFunc func() error, retErr error) {
	transportLabel := "Net"

	if ok, registry, tag, errMsg := dvmDockerLaunchEnvParams(transportLabel); ok {
		container, err := NewDockerDVMWithNetTransport(registry, tag, connectPort, dsServerPort, printLogs, args...)
		if err != nil {
			retErr = fmt.Errorf("%s: %w", errMsg, err)
			return
		}

		stopFunc = func() error {
			return container.Stop()
		}

		return
	}

	if ok, path, errMsg := dvmBinaryLaunchEnvParams(transportLabel); ok {
		cmd, err := NewBinaryDVMWithNetTransport(path, connectPort, dsServerPort, printLogs, args...)
		if err != nil {
			retErr = fmt.Errorf("%s: %w", errMsg, err)
			return
		}

		stopFunc = func() error {
			return cmd.Stop()
		}

		return
	}

	retErr = fmt.Errorf("Docker / Binary DVM launch option not specified: %s", os.Getenv(EnvDvmIntegUse))

	return
}

func LaunchDVMWithUDSTransport(socketsDir, connectSocketName, dsSocketName string, printLogs bool, args ...string) (stopFunc func() error, retErr error) {
	transportLabel := "UDS"

	if ok, registry, tag, errMsg := dvmDockerLaunchEnvParams(transportLabel); ok {
		container, err := NewDockerDVMWithUDSTransport(registry, tag, socketsDir, connectSocketName, dsSocketName, printLogs, args...)
		if err != nil {
			retErr = fmt.Errorf("%s: %w", errMsg, err)
			return
		}

		stopFunc = func() error {
			return container.Stop()
		}

		return
	}

	if ok, cmdPath, errMsg := dvmBinaryLaunchEnvParams(transportLabel); ok {
		cmd, err := NewBinaryDVMWithUDSTransport(cmdPath, socketsDir, connectSocketName, dsSocketName, printLogs, args...)
		if err != nil {
			retErr = fmt.Errorf("%s: %w", errMsg, err)
			return
		}

		stopFunc = func() error {
			return cmd.Stop()
		}

		return
	}

	retErr = fmt.Errorf("Docker / Binary DVM launch option not specified: %s", os.Getenv(EnvDvmIntegUse))

	return
}

func dvmDockerLaunchEnvParams(transportLabel string) (enabled bool, registry, tag, errMsg string) {
	if os.Getenv(EnvDvmIntegUse) != EnvDvmIntegUseDocker {
		return
	}
	enabled = true
	registry = os.Getenv(EnvDvmIntegDockerRegistry)
	tag = os.Getenv(EnvDvmIntegDockerTag)
	errMsg = fmt.Sprintf(TestErrFmt, EnvDvmIntegUseDocker, transportLabel)

	return
}

func dvmBinaryLaunchEnvParams(transportLabel string) (enabled bool, path, errMsg string) {
	if os.Getenv(EnvDvmIntegUse) != EnvDvmIntegUseBinary {
		return
	}
	enabled = true
	path = os.Getenv(EnvDvmIntegBinaryPath)
	errMsg = fmt.Sprintf(TestErrFmt, EnvDvmIntegUseBinary, transportLabel)

	return
}
