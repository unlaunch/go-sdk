package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// UnlaunchFactory ...
type UnlaunchFactory struct {
	sdkKey string
	cfg    *UnlaunchClientConfig
	logger logger.Interface
}

// NewUnlaunchClientFactory is a factory
func NewUnlaunchClientFactory(SDKKey string, cfg *UnlaunchClientConfig) (*UnlaunchFactory, error) {

	if SDKKey == "" {

	}

	if cfg == nil {
		cfg = DefaultConfig()
	}

	logging := logger.NewLogger(cfg.LoggerConfig)

	return &UnlaunchFactory{
		sdkKey: SDKKey,
		cfg:    cfg,
		logger: logging,
	}, nil
}

// Client ...
func (f *UnlaunchFactory) Client() *UnlaunchClient {

	eventsRecorder := api.NewHTTPEventsRecorder(
		f.sdkKey,
		f.cfg.Host,
		"/api/v1/impressions",
		f.cfg.MetricsFlushInterval,
		f.cfg.HTTPTimeout,
		f.cfg.MetricsQueueSize,
		"impressions",
		f.logger)
	return &UnlaunchClient{
		sdkKey:          f.sdkKey,
		pollingInterval: f.cfg.PollingInterval,
		httpTimeout:     f.cfg.HTTPTimeout,
		FeatureStore: service.NewHTTPFeatureStore(
			f.sdkKey,
			f.cfg.Host,
			f.cfg.HTTPTimeout,
			f.cfg.PollingInterval,
			f.logger),
		eventsRecorder: eventsRecorder,
		logger: f.logger,
	}

}
