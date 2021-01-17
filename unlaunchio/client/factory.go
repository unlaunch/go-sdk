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
	logger logger.LoggerInterface
}

type configMinimumValues struct {
	minPollingInterval int
	minHttpTimeout int
	minMetricsFlushInterval int
	minMetricsQueueSize int
}


var prodConfigMinValues = &configMinimumValues{
	minPollingInterval: 6000,
	minHttpTimeout: 1000,
	minMetricsFlushInterval: 45000,
	minMetricsQueueSize: 500,
}

var debugConfigMinValues = &configMinimumValues{
	minPollingInterval: 15000,
	minHttpTimeout: 1000,
	minMetricsFlushInterval: 15000,
	minMetricsQueueSize: 10,
}


func normalizeConfig(cfg *UnlaunchClientConfig, m *configMinimumValues) *UnlaunchClientConfig {
	var res *UnlaunchClientConfig

	if cfg == nil {
		res = &UnlaunchClientConfig{}
	} else {
		res = cfg
	}

	if cfg.PollingInterval < m.minPollingInterval {
		res.PollingInterval = m.minPollingInterval
	}

	if cfg.HTTPTimeout < 1000 {
		res.HTTPTimeout = m.minHttpTimeout
	}

	if cfg.MetricsFlushInterval < m.minMetricsFlushInterval {
		res.MetricsFlushInterval = m.minMetricsFlushInterval
	}

	if cfg.MetricsQueueSize < m.minMetricsQueueSize {
		res.MetricsQueueSize = m.minMetricsQueueSize
	}

	res.LoggerConfig = cfg.LoggerConfig

	return res
}


// NewUnlaunchClientFactory is a factory
func NewUnlaunchClientFactory(SDKKey string, cfg *UnlaunchClientConfig) (*UnlaunchFactory, error) {

	if SDKKey == "" {
		return nil, errors.New("the SDK Key cannot be null")
	}


	if cfg == nil {
		cfg = DefaultConfig()
	}

	var c *UnlaunchClientConfig
	if strings.HasPrefix(SDKKey, "prod") {
		c = normalizeConfig(cfg, prodConfigMinValues)
	} else {
		c = normalizeConfig(cfg, debugConfigMinValues)
	}

	logging := logger.NewLogger(c.LoggerConfig)

	return &UnlaunchFactory{
		sdkKey: SDKKey,
		cfg:    c,
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

	hc := util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger)

	return &UnlaunchClient{
		sdkKey:          f.sdkKey,
		pollingInterval: f.cfg.PollingInterval,
		httpTimeout:     f.cfg.HTTPTimeout,
		FeatureStore: service.NewHTTPFeatureStore(
			hc,
			f.cfg.PollingInterval,
			f.logger),
		eventsRecorder: eventsRecorder,
		eventsCountAggregator: eventsCounts,
		logger: f.logger,
		evaluator: engine.NewEvaluator(),
	}

}
