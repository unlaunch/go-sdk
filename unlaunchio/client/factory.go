package client

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"

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
	logger logger.Interface
}

var sync0Complete atomic.Value

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

	if sync0Complete.Load() == nil {
		sync0Complete.Store(true) // we preemptively marked it as done
		client := f.sync0()
		if client == nil { // regular server sync
			return f.regularServerSync()
		} else {
			return client
		}

	} else {
		return f.regularServerSync()
	}
}

func (f *UnlaunchFactory) regularServerSync() Client {

	eventsRecorder := api.NewHTTPEventsRecorder(
		false,
		util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger, false),
		"/api/v1/impressions",
		f.cfg.MetricsFlushInterval,
		f.cfg.MetricsQueueSize,
		"impressions",
		f.logger)

	eventsCounts := api.NewEventsCountAggregator(
		util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger, false),
		"/api/v1/events",
		f.cfg.MetricsFlushInterval,
		f.cfg.MetricsQueueSize,
		f.logger)

	hc := util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger, false)

	return &SimpleClient{
		FeatureStore: service.NewHTTPFeatureStore(
			hc,
			f.cfg.PollingInterval,
			f.logger,
			false,
			nil),
		eventsRecorder:        eventsRecorder,
		eventsCountAggregator: eventsCounts,
		logger:                f.logger,
		evaluator:             engine.NewEvaluator(f.logger),
	}
}

func (f *UnlaunchFactory) sync0() Client {
	hc := util.NewHTTPClient(f.sdkKey, f.cfg.Host, f.cfg.HTTPTimeout, f.logger, true)

	res, err := hc.Get("https://app-qa-unlaunch-io-master-flags.s3-us-west-1.amazonaws.com")

	if res == nil {
		return nil
	}

	if err != nil {
		f.logger.Error("[HTTP GET] error reading body", err)
		return nil
	}

	f.logger.Debug(fmt.Sprintf("[HTTP GET] data: %d", res))

	return &SimpleClient{
		FeatureStore: service.NewHTTPFeatureStore(
			hc,
			f.cfg.PollingInterval,
			f.logger,
			true,
			res),
		logger:    f.logger,
		evaluator: engine.NewEvaluator(f.logger),
	}

	// return res
}
