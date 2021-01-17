package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// UnlaunchClientConfig ...
type UnlaunchClientConfig struct {
	PollingInterval      int
	MetricsFlushInterval int
	MetricsQueueSize     int
	HTTPTimeout          int
	Host                 string
	LoggerConfig         *logger.LogOptions
}

// DefaultConfig ...
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      15000,
		HTTPTimeout:          3000,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 15000,
		MetricsQueueSize:     1000,
		LoggerConfig:         &logger.LogOptions{
			Level: "INFO",
			Colorful: true,
		},
	}
}

func defaultProductionConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      30000,
		HTTPTimeout:          4000,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 30000,
		MetricsQueueSize:     100,
		LoggerConfig:         &logger.LogOptions{
			Level: "INFO",
			Colorful: true,
		},
	}
}

func defaultNonProductionConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      15000,
		HTTPTimeout:          4000,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 15000,
		MetricsQueueSize:     5,
		LoggerConfig:         &logger.LogOptions{
			Level: "INFO",
			Colorful: true,
		},
	}
}
