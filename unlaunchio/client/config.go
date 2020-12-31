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
		pollingInterval: 30,
		httpTimeout:     30,
		host:			"https://api.unlaunch.io",
		loggerConfig:    nil,
	}
}
