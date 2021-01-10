package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// UnlaunchFactory ...
type UnlaunchFactory struct {
	sdkKey         string
	cfg            *UnlaunchClientConfig
	logger         logger.Interface
}

// NewUnlaunchClientFactory is a factory
func NewUnlaunchClientFactory(SDKKey string, cfg *UnlaunchClientConfig) (*UnlaunchFactory, error) {

	if SDKKey == "" {

	}

	if cfg == nil {
		cfg = DefaultConfig()
	}

	logging := logger.NewLogger(cfg.loggerConfig)

	return &UnlaunchFactory{
		sdkKey: SDKKey,
		cfg:    cfg,
		logger: logging,
	}, nil
}

// Client ...
func (f *UnlaunchFactory) Client() *UnlaunchClient {
	return &UnlaunchClient{
		sdkKey:          f.sdkKey,
		pollingInterval: f.cfg.pollingInterval,
		httpTimeout:     f.cfg.httpTimeout,
		FeatureStore:    api.NewHTTPFeatureStore(f.sdkKey, f.cfg.host, f.cfg.httpTimeout, f.cfg.pollingInterval, f.logger),
		eventsRecorder: api.NewHTTPEventsRecorder(f.sdkKey, f.cfg.host, "/api/v1/events", f.cfg.httpTimeout, f.logger),
		logger:          f.logger,
	}

}
