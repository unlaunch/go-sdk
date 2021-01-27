package client

import (
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"strings"
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
		return nil, errors.New("the SDK Key cannot be null")
	}

	var c *UnlaunchClientConfig
	c = normalizeConfigValues(cfg, strings.HasPrefix(SDKKey, "prod"))

	logging := logger.NewLogger(c.LoggerConfig)

	return &UnlaunchFactory{
		sdkKey: SDKKey,
		cfg:    c,
		logger: logging,
	}, nil
}

// Client ...
func (f *UnlaunchFactory) Client() Client {

	// TODO: Create and pass HTTP client instead of sdkey key, host
	// like eventsCount

	if f.cfg.OfflineMode {
		f.logger.Info("offline mode ", f.cfg.OfflineMode)
		return &OfflineClient{
			logger: f.logger,
		}
	}

	eventsRecorder := api.NewHTTPEventsRecorder(
		false,
		util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger),
		"/api/v1/impressions",
		f.cfg.MetricsFlushInterval,
		f.cfg.MetricsQueueSize,
		"impressions",
		f.logger)

	eventsCounts := api.NewEventsCountAggregator(
		util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger),
		"/api/v1/events",
		f.cfg.MetricsFlushInterval,
		f.cfg.MetricsQueueSize,
		f.logger)

	hc := util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger)

	return &SimpleClient{
		FeatureStore: service.NewHTTPFeatureStore(
			hc,
			f.cfg.PollingInterval,
			f.logger),
		eventsRecorder:        eventsRecorder,
		eventsCountAggregator: eventsCounts,
		logger:                f.logger,
		evaluator:             engine.NewEvaluator(f.logger),
	}

}
