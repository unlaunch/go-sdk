package client

import "github.com/unlaunch/go-sdk/unlaunchio/util/logger"

// UnlaunchClientConfig ...
type UnlaunchClientConfig struct {
	pollingInterval int
	httpTimeout     int
	host 			string
	loggerConfig    *logger.Options
}

// DefaultConfig ...
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		pollingInterval: 15000,
		httpTimeout:     3000,
		host:			"https://api.unlaunch.io",
		loggerConfig:    nil,
	}
}
