package tests

import (
	"encoding/json"
	"fmt"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
	"github.com/unlaunch/go-sdk/unlaunchio/engine"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/number.json")


func initializeNum() (dtos.Feature, string) {
	var responseDtoNum dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDtoNum)
	sort.Sort(dtos.ByRulePriority(responseDtoNum.Rules))
	for _, rule := range responseDtoNum.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDtoNum, userId
}

func TestWhen_NoNumberAttributesArePassed_Then_DefaultRuleIsServed(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "defrule"

	ulf, _ := engine.Evaluate(&r, u, nil)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForFloat32(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n float32 = 1.0
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForFloat64(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n float64 = 1.0
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForInt8(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n int8 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForInt16(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n int16 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForInt32(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n int32 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForInt64(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n int64 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForInt(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n int = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForUInt(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n uint = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForUInt8(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n uint8 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForUInt16(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n uint16 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForUInt32(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n uint32 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_NumberEqualsRuleMatchForUInt64(t *testing.T) {
	r, u := initializeNum()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	var n uint64 = 1
	attributes["numberAttr"] = n
	ulf, _ := engine.Evaluate(&r, u, &attributes)
	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}