package client

import "github.com/unlaunch/go-sdk/unlaunchio/util/logger"

// UnlaunchClientConfig ...
type UnlaunchClientConfig struct {
	PollingInterval      int
	MetricsFlushInterval int
	MetricsQueueSize     int
	HTTPTimeout          int
	Host                 string
	LoggerConfig         *logger.Options
}

// DefaultConfig ...
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      15000,
		HTTPTimeout:          3000,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 15000,
		MetricsQueueSize:     1000,
		LoggerConfig:         nil,
	}
}
