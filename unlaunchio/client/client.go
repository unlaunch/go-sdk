package client

import (
	"errors"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"runtime/debug"
	"time"
)

// SimpleClient is the main interface for interacting with Unlaunch
type SimpleClient struct {
	FeatureStore          service.FeatureStore
	eventsRecorder        api.EventsRecorder
	eventsCountAggregator api.EventsCountAggregator
	logger                logger.Interface
	shutdown              bool
	evaluator             engine.Evaluator
}

// Feature ...
func (c *SimpleClient) Feature(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) *dtos.UnlaunchFeature {
	return c.processFlagEvaluation(featureKey, identity, attributes)
}

// IsShutdown returns true if the client has already shutdown
func (c *SimpleClient) IsShutdown() bool {
	return c.shutdown
}

// Variation evaluates and returns the variation (variation key) for this feature. Variations are defined using the
// Unlaunch console at https://app.unlaunch.io
func (c *SimpleClient) Variation(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) string {
	return c.processFlagEvaluation(featureKey, identity, attributes).Variation
}

// processFlagEvaluation evaluates a flag and then emits metrics
func (c *SimpleClient) processFlagEvaluation(
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

	ulf := c.evaluateFlag(featureKey, identity, attributes)

	if ulf.Variation != "" && ulf.Variation != "control" {

		// Record event
		c.eventsCountAggregator.Record(featureKey, ulf.Variation)

		// Record impression
		c.eventsRecorder.Record(&dtos.Event{
			CreatedTime: time.Now().UTC().UnixNano() / int64(time.Millisecond), // java time
			Key:         featureKey,
			Type:        "IMPRESSION",
			Properties:  nil,
			Sdk:         "Go",
			SdkVersion:  "0.0.1",
			Impression: dtos.Impression{
				FlagKey:          featureKey,
				UserID:           identity,
				VariationKey:     ulf.Variation,
				EvaluationReason: ulf.EvaluationReason,
				MachineName:      "UNKNOWN",
			},
		})
	} else {
		c.logger.Warn("skipping count and impression recorders both variation was not valid")
	}

	return ulf
}

func (c *SimpleClient) evaluateFlag(
	featureKey string,
	identity string,
	attributes map[string]interface{}) *dtos.UnlaunchFeature {
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

	if !c.FeatureStore.IsReady() {
		c.logger.Warn("the SDK is not ready. Returning the SDK default 'control' as variation which may not give the right result")
		return &dtos.UnlaunchFeature{
			Feature:                featureKey,
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "sdk was not ready",
		}
	}

	feature, err := c.FeatureStore.GetFeature(featureKey)

	if err != nil || feature == nil {
		c.logger.Error("error retrieving flag: ", err)
		return &dtos.UnlaunchFeature{
			Feature:                featureKey,
			Variation:              "control",
			VariationConfiguration: nil,
			EvaluationReason:       "feature flag was not found in memory",
		}
	}

	ulFeature, err := c.evaluator.Evaluate(feature, identity, &attributes)

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

// AwaitUntilReady blocks until the client initialization is done or timeout occurs, whichever comes first
func (c *SimpleClient) AwaitUntilReady(timeout time.Duration) error {
	if c.FeatureStore.IsReady() {
		return nil
	}

	if timeout <= 0 {
		return errors.New("the timeout must be a positive")
	}

	if c.shutdown {
		return errors.New("the client has been shutdown")
	}

	c.FeatureStore.Ready(timeout)
	return nil
}

// Shutdown shuts down the client, and all associated go routines
func (c *SimpleClient) Shutdown() {
	if !c.shutdown {
		c.FeatureStore.Shutdown()
		c.eventsRecorder.Shutdown()
		c.eventsCountAggregator.Shutdown()
		c.shutdown = true
	}
}
