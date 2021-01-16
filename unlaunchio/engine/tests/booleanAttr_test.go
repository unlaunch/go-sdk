package tests

import (
	"encoding/json"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func initializeBoolAttr() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/boolean.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func TestWhen_BooleanDefRule(t *testing.T) {
	r, u := initializeBoolAttr()

	expectedVariation := "defrule"

	ulf, _ := evaluator.Evaluate(&r, u, nil)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_BooleanEqualsTrueRuleMatch(t *testing.T) {
	r, u := initializeBoolAttr()

	expectedVariation := "off"

	attributes := make(map[string]interface{})
	attributes["boolAttr"] = true

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_BooleanEqualsFalseRuleMatch(t *testing.T) {
	r, u := initializeBoolAttr()

	expectedVariation := "on"

	attributes := make(map[string]interface{})
	attributes["boolAttr"] = false

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func initializeBoolAttrNotEquals() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/boolean_notequals.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func TestWhen_BooleanDefRuleNotEqualsFile(t *testing.T) {
	r, u := initializeBoolAttrNotEquals()

	expectedVariation := "defrule"

	ulf, _ := evaluator.Evaluate(&r, u, nil)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_BooleanNotEqualsTrue(t *testing.T) {
	r, u := initializeBoolAttrNotEquals()

	expectedVariation := "offf"

	attributes := make(map[string]interface{})
	attributes["boolAttr"] = false

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}
