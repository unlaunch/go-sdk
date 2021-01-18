package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"time"
)

// UnlaunchClientConfig contains configuration parameters to fine-tune the client
type UnlaunchClientConfig struct {
	PollingInterval      time.Duration
	MetricsFlushInterval time.Duration
	MetricsQueueSize     int
	HTTPTimeout          time.Duration
	Host                 string
	OfflineMode          bool
	LoggerConfig         *logger.LogOptions
}

// DefaultConfig for our users. Not used anywhere internally
func DefaultConfig() *UnlaunchClientConfig {
	return &UnlaunchClientConfig{
		PollingInterval:      60 * time.Second,
		HTTPTimeout:          10 * time.Second,
		Host:                 "https://api.unlaunch.io",
		MetricsFlushInterval: 45 * time.Second,
		MetricsQueueSize:     500,
		OfflineMode:          false,
		LoggerConfig: 		  nil,
	}
}

type configValues struct {
	pollingInterval      time.Duration
	httpTimeout          time.Duration
	metricsFlushInterval time.Duration
	metricsQueueSize     int
	host                 string
}

// config minimums
var minValues = &configValues{
	pollingInterval:      15 * time.Second,
	httpTimeout:          1 * time.Second,
	metricsFlushInterval: 10 * time.Second,
	metricsQueueSize:     10,
	host:                 "https://api.unlaunch.io",
}

func normalizeConfigValues(cfg *UnlaunchClientConfig, prod bool) *UnlaunchClientConfig {
	var res *UnlaunchClientConfig

	if cfg == nil {
		if prod {
			return DefaultConfig()
		} else {
			return &UnlaunchClientConfig{
				PollingInterval:      15 * time.Second,
				HTTPTimeout:          10 * time.Second,
				Host:                 "https://api.unlaunch.io",
				MetricsFlushInterval: 15 * time.Second,
				MetricsQueueSize:     20,
				OfflineMode:          false,
				LoggerConfig: 		  nil,
			}
		}
	} else {
		res = cfg
	}

	// make sure that no setting is set below its minimum value or is wrong
	if res.PollingInterval  < minValues.pollingInterval {
			res.PollingInterval = minValues.pollingInterval
	}

	if res.HTTPTimeout < minValues.httpTimeout {
		res.HTTPTimeout = minValues.httpTimeout
	}

	if res.MetricsFlushInterval < minValues.metricsFlushInterval {
		res.MetricsFlushInterval = minValues.metricsFlushInterval
	}

	if res.MetricsQueueSize < minValues.metricsQueueSize {
		res.MetricsQueueSize = minValues.metricsQueueSize
	}

	if res.Host == "" {
		res.Host = "https://api.unlaunch.io"
	}


	return res
}
