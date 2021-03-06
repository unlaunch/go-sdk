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

func initializeNumAttr() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/number.json")

	var responseDtoNum dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDtoNum)
	sort.Sort(dtos.ByRulePriority(responseDtoNum.Rules))
	for _, rule := range responseDtoNum.Rules {
		sort.Sort(dtos.ByVariationID(rule.Rollout))
	}

	userID := "user-" + strconv.Itoa(rand.Intn(1000))

	return responseDtoNum, userID
}

func TestWhen_NumberGreaterThan(t *testing.T) {
	r, u := initializeNumAttr()

	expectedVariation := "gt"

	attributes := make(map[string]interface{})
	n := 101
	attributes["numberAttr"] = n

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberGreaterThanOrEquals(t *testing.T) {
	r, u := initializeNumAttr()

	expectedVariation := "gte"

	attributes := make(map[string]interface{})
	n := 50
	attributes["numberAttr"] = n

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}
func TestWhen_NumberLessThan(t *testing.T) {
	r, u := initializeNumAttr()

	expectedVariation := "lt"

	attributes := make(map[string]interface{})
	n := -101
	attributes["numberAttr"] = n

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberLessThanOrEquals(t *testing.T) {
	r, u := initializeNumAttr()

	expectedVariation := "lte"

	attributes := make(map[string]interface{})
	n := -51
	attributes["numberAttr"] = n

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}
