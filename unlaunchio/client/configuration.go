package client

import "github.com/unlaunch/go-sdk/unlaunchio/util/logger"

// UnlaunchClientConfig ...
type UnlaunchClientConfig struct {
	pollingInterval int
	httpTimeout     int
	loggerConfig    *logger.Options
}

// DefaultConfig ...
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		pollingInterval: 30,
		httpTimeout:     30,
		loggerConfig:    nil,
	}
}
