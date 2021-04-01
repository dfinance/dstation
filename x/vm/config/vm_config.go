package config

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	tmOs "github.com/tendermint/tendermint/libs/os"
)

const (
	VmConfigFile      = "vm.toml" // Default VM config file name
	DefaultConfigsDir = "config"

	// VM connection defaults
	DefaultVMAddress  = "tcp://127.0.0.1:50051"
	DefaultDataListen = "tcp://127.0.0.1:50052"

	// VM request retry defaults
	DefaultMaxAttempts = 0
	DefaultReqTimeout  = 0
)

// VMConfig defines virtual machine connection config.
type VMConfig struct {
	// Virtual machine address to connect from Cosmos SDK
	Address string `mapstructure:"vm_address"`
	// Node's data server address to listen for connections from VM
	DataListen string `mapstructure:"vm_data_listen"` // node's data server listen address

	// Retry policy: maximum retry attempts (0 - infinity)
	MaxAttempts uint `mapstructure:"vm_retry_max_attempts"`
	// Retry policy: request timeout per attempt [ms] (0 - infinite, no timeout)
	ReqTimeoutInMs uint `mapstructure:"vm_retry_req_timeout_ms"`
}

// DefaultVMConfig returns VMConfig with defaults.
func DefaultVMConfig() VMConfig {
	return VMConfig{
		Address:        DefaultVMAddress,
		DataListen:     DefaultDataListen,
		MaxAttempts:    DefaultMaxAttempts,
		ReqTimeoutInMs: DefaultReqTimeout,
	}
}

// WriteVMConfig writes VM config file in configuration directory.
func WriteVMConfig(configPath string, vmConfig VMConfig) {
	var buffer bytes.Buffer
	if err := configTemplate.Execute(&buffer, vmConfig); err != nil {
		panic(err)
	}

	tmOs.MustWriteFile(configPath, buffer.Bytes(), 0644)
}

// ReadVMConfig reads VM config file from configuration directory.
func ReadVMConfig(homeDir string) VMConfig {
	configFilePath := filepath.Join(homeDir, DefaultConfigsDir, VmConfigFile)

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		defConfig := DefaultVMConfig()
		WriteVMConfig(configFilePath, defConfig)
		return defConfig
	}

	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := VMConfig{}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
