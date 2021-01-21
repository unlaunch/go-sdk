package client

import (
	"runtime/debug"
	"time"

	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"github.com/unlaunch/go-sdk/unlaunchio/service"
	"github.com/unlaunch/go-sdk/unlaunchio/service/api"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
)

// OfflineClient is the main interface for interacting with Unlaunch
type OfflineClient struct {
	FeatureStore          service.FeatureStore
	eventsRecorder        api.EventsRecorder
	eventsCountAggregator api.EventsCountAggregator
	logger                logger.Interface
	shutdown              bool
	evaluator             engine.Evaluator
}

// Feature ...
func (c *OfflineClient) Feature(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) *dtos.UnlaunchFeature {
	return c.processFlagEvaluation(featureKey, identity, attributes)
}

func (c *OfflineClient) IsShutdown() bool {
	return c.shutdown
}

// Variation evaluates and returns the variation (variation key) for this feature. Variations are defined using the
// Unlaunch console at https://app.unlaunch.io
func (c *OfflineClient) Variation(
	featureKey string,
	identity string,
	attributes map[string]interface{},
) string {
	return c.processFlagEvaluation(featureKey, identity, attributes).Variation
}

// processFlagEvaluation evaluates a flag and then emits metrics
func (c *OfflineClient) processFlagEvaluation(
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

	return ulf
}

func (c *OfflineClient) evaluateFlag(
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

	controlFeature := &dtos.UnlaunchFeature{
		Feature:                featureKey,
		Variation:              "control",
		VariationConfiguration: nil,
		EvaluationReason:       "Client is initialized in Offline Mode. Returning 'control' variation for all flags.",
	}

	c.logger.Debug("flag evaluation reason: ", controlFeature.EvaluationReason)

	return controlFeature
}

// AwaitUntilReady blocks until the client initialization is done or timeout occurs, whichever comes first
func (c *OfflineClient) AwaitUntilReady(timeout time.Duration) error {
	// Do nothing same as Java Sdk
	return nil
}

// Shutdown shuts down the client, and all associated go routines
func (c *OfflineClient) Shutdown() {
	// Do nothing same as Java Sdk
}
