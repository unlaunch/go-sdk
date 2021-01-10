package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/store"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"runtime/debug"
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
func (c *UnlaunchClient) Feature(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) *dtos.UnlaunchFeature {
	return c.evaluateFlag(featureKey, identity, attributes)
}

// Variation ...
func (c *UnlaunchClient) Variation(
	featureKey string,
	identity string,
	attributes map[string]interface{},
	) string {

	return c.evaluateFlag(featureKey, identity, attributes).Variation
}


// Variation ...
func (c *UnlaunchClient) evaluateFlag(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) (ul *dtos.UnlaunchFeature) {

	// Guard function if SDK panics to return control
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error("SDK is panicking. Error", r, "\n", string(debug.Stack()), "\n")
			ul = &dtos.UnlaunchFeature{
				Feature: featureKey,
				Variation: "control",
				VariationConfiguration: nil,
				EvaluationReason: "SDK panicked. check logs.",
			}
		}
	}()

	return processFlag(featureKey, identity, attributes, c)
}

func (c *UnlaunchClient) BlockUntilReady(timeout uint32) error {
	c.FeatureStore.Ready()
	return nil
}


func processFlag(
	featureKey string,
	identity string,
	attributes map[string]interface{},
	c *UnlaunchClient) *dtos.UnlaunchFeature {
	if featureKey == "" {
		c.logger.Error("feature key cannot be empty")
		return &dtos.UnlaunchFeature{
			Feature: "",
			Variation: "control",
			VariationConfiguration: nil,
			EvaluationReason: "feature key was empty string. You must provide the key of the feature flag to evaluate",
		}

	}

	if identity == "" {
		c.logger.Error("identity key cannot be empty")
		return &dtos.UnlaunchFeature{
			Feature: featureKey,
			Variation: "control",
			VariationConfiguration: nil,
			EvaluationReason: "identity (id) was empty string. You must provide a unique value per user",
		}

	}

	feature, err := c.FeatureStore.GetFeature(featureKey)

	if err != nil {
		c.logger.Error("error retrieving flag: ", err)
		return &dtos.UnlaunchFeature{
			Feature: featureKey,
			Variation: "control",
			VariationConfiguration: nil,
			EvaluationReason: "feature flag was not found in memory",
		}
	}

	ulFeature, err := engine.Evaluate(feature, identity, &attributes)

	if err != nil {
		c.logger.Error("error evaluating flag: ", err)
		return &dtos.UnlaunchFeature{
			Feature: featureKey,
			Variation: "control",
			VariationConfiguration: nil,
			EvaluationReason: "there was an error evaluating flag. see logs.",
		}
	}

	c.logger.Debug("flag evaluation reason: ", ulFeature.EvaluationReason)

	return ulFeature
}