package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"log"
	"os"
)

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
		LoggerConfig:         &logger.Options{
			BaseLogger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix),
			Level: "DEBUG",
		},
	}
}
