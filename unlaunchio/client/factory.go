package client

type UnlaunchFactory struct {
	Client *UnlaunchClient
}

func NewUnlaunchClientFactory(config *UnlaunchClientConfig) *UnlaunchFactory {
	client := &UnlaunchClient{
		SDKKey:          config.SDKKey,
		PollingInterval: config.PollingInterval,
		HTTPTimeout:     config.HTTPTimeout,
	}

	return &UnlaunchFactory{
		Client: client,
	}

}
