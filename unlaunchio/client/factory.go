package client

import "github.com/unlaunch/go-sdk/unlaunchio/util/logger"

// UnlaunchFactory ...
type UnlaunchFactory struct {
	sdkKey string
	config *UnlaunchClientConfig
	logger logger.Interface
}

// NewUnlaunchClientFactory is a factory
func NewUnlaunchClientFactory(SDKKey string, cfg *UnlaunchClientConfig) (*UnlaunchFactory, error) {

	if cfg == nil {
		cfg = DefaultConfig()
	}

	logging := logger.NewLogger(cfg.loggerConfig)

	return &UnlaunchFactory{
		sdkKey: SDKKey,
		config: cfg,
		logger: logging,
	}, nil
}

func (f *UnlaunchFactory) Client() *UnlaunchClient {
	return &UnlaunchClient{
		sdkKey:          f.sdkKey,
		pollingInterval: f.config.pollingInterval,
		httpTimeout: f.config.httpTimeout,
		logger: f.logger,
	}

}
