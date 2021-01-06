package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/store"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// UnlaunchClient Main Unlaunch Client
type UnlaunchClient struct {
	sdkKey          string
	pollingInterval int
	httpTimeout     int
	FeatureStore    store.FeatureStore
	logger          logger.Interface
}

// Variation ...
func (c *UnlaunchClient) Variation(
	featureKey string,
	identity string,
	attributes *map[string]interface{},
	) string {
	if featureKey == "" {
		c.logger.Error("feature key cannot be empty")
		return "control"
	}

	if identity == "" {
		c.logger.Error("identity key cannot be empty")
		return "control"
	}

	feature, err := c.FeatureStore.GetFeature(featureKey)

	if err != nil {
		c.logger.Error("error retrieving flag: ", err)
		return "control"
	}

	ulFeature, err := engine.Evaluate(feature, identity, attributes)

	if err != nil {
		c.logger.Error("error evaluating flag: ", err)
		return "control"
	}

	 c.logger.Debug("flag evaluation reason: ", ulFeature.EvaluationReason)

	return ulFeature.Variation.Key
}

func (c *UnlaunchClient) BlockUntilReady(timeout uint32) error {
	c.FeatureStore.Ready()
	return nil
}
