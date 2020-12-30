package client

import "github.com/unlaunch/go-sdk/unlaunchio/util/logger"

// UnlaunchClient Main Unlaunch Client
type UnlaunchClient struct {
	sdkKey          string
	pollingInterval int
	httpTimeout     int
	logger          logger.Interface
}

// GetVariation ...
func (c *UnlaunchClient) GetVariation(feature string, identity string, attributes *map[string]interface{}) string {
	if feature == "" {
		c.logger.Error("feature key cannot be empty")
		return "control"
	}

	if identity == "" {
		c.logger.Error("identity key cannot be empty")
		return "control"
	}
	return "control"
}

func (c *UnlaunchClient) BlockUntilReady(timeout uint32) error {
	return nil
}
