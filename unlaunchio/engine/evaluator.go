package engine

import (
	"errors"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	attributes2 "github.com/unlaunch/go-sdk/unlaunchio/engine/attributes"
	"github.com/unlaunch/go-sdk/unlaunchio/util/logger"
	"math"
	"strings"
)

var tr = attributes2.NewTargetingRule()


type Evaluator interface {
	Evaluate(feature *dtos.Feature, identity string, attributes *map[string]interface{})(*dtos.UnlaunchFeature, error)
}

type SimpleEvaluator struct{
	logger	logger.LoggerInterface
}

func NewEvaluator(logger logger.LoggerInterface) Evaluator {
	return &SimpleEvaluator{
		logger: logger,
	}
}

func (e *SimpleEvaluator) Evaluate(
	feature *dtos.Feature,
	identity string,
	attributes *map[string]interface{})(*dtos.UnlaunchFeature, error) {
	result := new(dtos.UnlaunchFeature)
	result.Feature = feature.Key
	result.EvaluationReason = "this SDK is not yet complete"

	if feature.Enabled() == false {
		offVariation, err := e.getOffVariation(feature)
		if err != nil {
			e.logger.Error("unexpected error.", e)
			return nil, err
		}
		result.EvaluationReason = "Flag disabled. Default Variation served"
		result.Variation = offVariation.Key
		result.VariationConfiguration = offVariation.Properties
		return result, nil
	} else if v := e.variationIfUserInAllowList(feature, identity); v != nil {
		result.Variation = v.Key
		result.VariationConfiguration = v.Properties
		result.EvaluationReason = "User is in Target Users List"
		return result, nil
	} else if v := e.matchTargetingRules(feature, identity, attributes); v != nil {
		result.Variation = v.Key
		result.VariationConfiguration = v.Properties
		result.EvaluationReason = "Targeting Rule Match"
		return result, nil
	} else if v := e.defaultRule(feature, identity, attributes); v != nil {
		result.Variation = v.Key
		result.VariationConfiguration = v.Properties
		result.EvaluationReason = "Default Rule Match"
		return result, nil
	} else {
		return nil, errors.New("something went wrong")
	}
}

func (e *SimpleEvaluator)getOffVariation(f *dtos.Feature) (*dtos.Variation, error) {
	offVarId := f.OffVariation

	for _, variation := range f.Variations {
		if offVarId == variation.Id {
			return &variation, nil
		}
	}

	return &dtos.Variation{}, errors.New(
		fmt.Sprintf("error - offVariation %d was not found", offVarId),
	)
}

func (e *SimpleEvaluator)variationIfUserInAllowList(f *dtos.Feature, identity string) *dtos.Variation {
	for _, variation := range f.Variations {
		if variation.AllowList != "" {

			l := strings.Split(strings.ReplaceAll(variation.AllowList, " ", ""), ",")

			for _, id := range l {
				if identity == id {
					return &variation
				}
			}
		}
	}

	return nil
}

func (e *SimpleEvaluator)matchTargetingRules(feature *dtos.Feature, identity string, attributes *map[string]interface{}) *dtos.Variation {

	if attributes == nil {
		return nil
	}

	for _, rule := range feature.Rules {
		matched := false

		// Iterate through all conditions
		for _, condition := range rule.Conditions {
			if attr, ok := (*attributes)[condition.Attribute]; ok {
				if !tr.Apply(condition.Type, condition.Value, attr, condition.Op) {
					matched = false
					break
				} else {
					matched = true
				}
			}
		}
		if matched {
			return e.getRuleVariation(&rule, feature, identity)
		}
	}

	return nil
}

func (e *SimpleEvaluator)defaultRule(feature *dtos.Feature, identity string, attributes *map[string]interface{}) *dtos.Variation {
	defaultRule := feature.DefaultRule()
	return e.getRuleVariation(defaultRule, feature, identity)
}

func (e *SimpleEvaluator)getRuleVariation(rule *dtos.Rule, feature *dtos.Feature, identity string) *dtos.Variation {
	calculatedBucket := e.bucket(feature.Key + identity)

	// Return Default Rule if targeting rules don't match
	if len(rule.Rollout) == 1 && rule.Rollout[0].RolloutPercentage == 100 {
		return e.variationById(rule.Rollout[0].VariationId, feature)
	} else {
		vId, _ := e.foo(rule, calculatedBucket)
		return feature.VariationById(vId)
	}

	return nil
}

func (e *SimpleEvaluator)variationById(id int, feature *dtos.Feature) *dtos.Variation {
	for _, variation := range feature.Variations {
		if id == variation.Id {
			return &variation
		}
	}
	return nil
}

func (e *SimpleEvaluator)bucket(key string) int {
	var hashKey uint32
	hashKey = Murmur32Hash([]byte(key), 0)
	return int(math.Abs(float64(hashKey%100)) + 1)
}

func (e *SimpleEvaluator)foo(rule *dtos.Rule, bucket int) (int, error) {
	var sum = 0
	for _, rollout := range rule.Rollout {
		sum += rollout.RolloutPercentage

		if bucket <= sum {
			return rollout.VariationId, nil
		}
	}

	return -1, errors.New("unable to find variation. Internal error")

}
