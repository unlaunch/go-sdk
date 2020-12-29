package client

// Main Unlaunch Client

type UnlaunchClient struct {
	SDKKey          string
	PollingInterval int
	HTTPTimeout     int
}

func (c *UnlaunchClient) GetVariation(identity string, feature string, attributes *map[string]interface{}) string {
	return "control"
}
