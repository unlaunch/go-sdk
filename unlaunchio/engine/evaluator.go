package engine

import (
	"errors"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"strings"
)

func Evaluate(feature *dtos.Feature, identity string, attributes *map[string]interface{}) (*dtos.UnlaunchFeature, error) {
	result := new(dtos.UnlaunchFeature)
	result.Feature = feature.Key
	result.EvaluationReason = "this SDK is not yet complete"

	if feature.Enabled() == false {
		result.EvaluationReason = "Flag disabled. Default Variation served"
		offVariation, err := getOffVariation(feature)

		if err != nil {
			return nil, err
		}

		result.Variation = offVariation
		return result, nil
	} else if v := variationIfUserInAllowList(feature, identity); v != nil {
		result.Variation = v
		result.EvaluationReason = "User is in Target Users List"
		return result, nil
	} else if v := matchTargetingRules(feature, identity, attributes); v != nil {
		result.Variation = v
		result.EvaluationReason = "Targeting Rule Match"
		return result, nil

	} else if v := defaultRule(feature, identity, attributes); v != nil {
		result.Variation = v
		result.EvaluationReason = "Default Rule Match"
		return result, nil
	} else {
		return nil, errors.New("not yet implemented")
	}
}

func getOffVariation(f *dtos.Feature) (*dtos.Variation, error) {
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

func variationIfUserInAllowList(f *dtos.Feature, identity string) *dtos.Variation {
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



func matchTargetingRules(feature *dtos.Feature, identity string, attributes *map[string]interface{}) *dtos.Variation {
	return nil
}

func defaultRule(feature *dtos.Feature, identity string, attributes *map[string]interface{}) *dtos.Variation {
	var defaultRule dtos.Rule
	for _, rule := range feature.Rules {
		if rule.IsDefault {
			defaultRule = rule
		}
	}

	// Return Default Rule if targeting rules don't match
	if len(defaultRule.Rollout) == 1 && defaultRule.Rollout[0].RolloutPercentage == 100 {
		return variationById(defaultRule.Rollout[0].VariationId, feature)
	}

	return nil
}



func variationById(id int, feature *dtos.Feature) *dtos.Variation {
	for _, variation := range feature.Variations {
		if id == variation.Id {
			return &variation
		}
	}

	return nil
}
