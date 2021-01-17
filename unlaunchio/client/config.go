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

// DefaultConfig for our users. Not used anywhere internally
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      15000,
		HTTPTimeout:          3000,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 15000,
		MetricsQueueSize:     100,
		LoggerConfig: &logger.LogOptions{
			Level:    "INFO",
			Colorful: true,
		},
	}
}

type configMinimumValues struct {
	minPollingInterval      int
	minHTTPTimeout          int
	minMetricsFlushInterval int
	minMetricsQueueSize     int
	host                    string
}

var prodConfigMinValues = &configMinimumValues{
	minPollingInterval:      60000,
	minHTTPTimeout:          1000,
	minMetricsFlushInterval: 45000,
	minMetricsQueueSize:     500,
	host:                    "https://api.unlaunch.io",
}

var debugConfigMinValues = &configMinimumValues{
	minPollingInterval:      15000,
	minHTTPTimeout:          1000,
	minMetricsFlushInterval: 15000,
	minMetricsQueueSize:     10,
	host:                    "https://api.unlaunch.io",
}

func normalizeConfigValues(cfg *UnlaunchClientConfig, m *configMinimumValues) *UnlaunchClientConfig {
	var res *UnlaunchClientConfig

	if cfg == nil {
		res = &UnlaunchClientConfig{}
	} else {
		res = cfg
	}

	if res.PollingInterval < m.minPollingInterval {
		res.PollingInterval = m.minPollingInterval
	}

	if res.HTTPTimeout < m.minHTTPTimeout {
		res.HTTPTimeout = m.minHTTPTimeout
	}

	if res.MetricsFlushInterval < m.minMetricsFlushInterval {
		res.MetricsFlushInterval = m.minMetricsFlushInterval
	}

	if res.MetricsQueueSize < m.minMetricsQueueSize {
		res.MetricsQueueSize = m.minMetricsQueueSize
	}

	if res.Host == "" {
		res.Host = "https://api.unlaunch.io"
	}

	if res.LoggerConfig == nil {
		res.LoggerConfig = &logger.LogOptions{
			Level:    "INFO",
			Colorful: true,
		}
	}

	return res
}
