package client

import (
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"runtime/debug"
	"time"
)

// UnlaunchClient Main Unlaunch Client
type UnlaunchClient struct {
	sdkKey          string
	pollingInterval int
	httpTimeout     int
	FeatureStore    service.FeatureStore
	eventsRecorder  *api.EventsRecorder
	eventsCountAggregator *api.EventsCountAggregator
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
				Feature:                featureKey,
				Variation:              "control",
				VariationConfiguration: nil,
				EvaluationReason:       "SDK panicked. check logs.",
			}
		}
	}()

	ulf := processFlag(featureKey, identity, attributes, c)

	if ulf.Variation != "" || ulf.Variation != "control" {
		c.eventsCountAggregator.Record(featureKey, ulf.Variation)


		event := &dtos.Event{
			CreatedTime:  time.Now().UTC().Unix() * 1000,
			Key:          featureKey,
			Type: "IMPRESSION",
			Properties:   nil,
			Sdk:          "Go",
			SdkVersion:   "0.0.1",
			Impression:   dtos.Impression{
				FlagKey:          featureKey,
				UserId:           identity,
				VariationKey:     ulf.Variation,
				EvaluationReason: ulf.EvaluationReason,
				MachineName:      "UNKNOWN",
			},
		}

		c.eventsRecorder.Record(event)
	}

	return ulf
}

func (c *UnlaunchClient) BlockUntilReady(timeout uint32) error {
	c.FeatureStore.Ready()
	return nil
}

func (c *UnlaunchClient) Shutdown() {
	c.FeatureStore.Stop()
	c.eventsRecorder.Shutdown()
	c.eventsCountAggregator.Shutdown()
}


func processFlag(
	featureKey string,
	identity string,
	attributes map[string]interface{},
	c *UnlaunchClient) *dtos.UnlaunchFeature {
	if featureKey == "" {
		c.logger.Error("feature key cannot be empty")
		return &dtos.UnlaunchFeature{
			Feature:                "",
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "feature key was empty string. You must provide the key of the feature flag to evaluate",
		}

	}

	if identity == "" {
		c.logger.Error("identity key cannot be empty")
		return &dtos.UnlaunchFeature{
			Feature:                featureKey,
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "identity (id) was empty string. You must provide a unique value per user",
		}
	}

	feature, err := c.FeatureStore.GetFeature(featureKey)

	if err != nil {
		c.logger.Error("error retrieving flag: ", err)
		return &dtos.UnlaunchFeature{
			Feature:                featureKey,
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "feature flag was not found in memory",
		}
	}

	ulFeature, err := engine.Evaluate(feature, identity, &attributes)

	if err != nil {
		c.logger.Error("error evaluating flag: ", err)
		return &dtos.UnlaunchFeature{
			Feature:                featureKey,
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "there was an error evaluating flag. see logs.",
		}
	}

	c.logger.Debug("flag evaluation reason: ", ulFeature.EvaluationReason)

	return ulFeature
}
