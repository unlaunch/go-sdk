package client

import (
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// UnlaunchFactory ...
type UnlaunchFactory struct {
	sdkKey string
	cfg    *UnlaunchClientConfig
	logger logger.LoggerInterface
}

// NewUnlaunchClientFactory is a factory
func NewUnlaunchClientFactory(SDKKey string, cfg *UnlaunchClientConfig) (*UnlaunchFactory, error) {

	if SDKKey == "" {
		return nil, errors.New("the SDK Key cannot be null")
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

	// TODO: Create and pass HTTP client instead of sdkey key, host
	// like eventsCount
	eventsRecorder := api.NewHTTPEventsRecorder(
		f.sdkKey,
		f.cfg.Host,
		"/api/v1/impressions",
		f.cfg.HTTPTimeout,
		f.cfg.MetricsFlushInterval,
		f.cfg.MetricsQueueSize,
		"impressions",
		f.logger)


	eventsCounts := api.NewEventsCountAggregator(
		util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger),
		"/api/v1/events",
		f.cfg.MetricsFlushInterval,
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
		eventsCountAggregator: eventsCounts,
		logger: f.logger,
		evaluator: engine.NewEvaluator(),
	}

}
