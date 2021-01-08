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

func initialize() (dtos.Feature, string) {
	var mockedFlag, _ = ioutil.ReadFile("../../testdata/attributes/string.json")
	var responseDto dtos.Feature
	var userId = "user-"+ strconv.Itoa(rand.Intn(1000))

	json.Unmarshal(mockedFlag, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	return responseDto, userId
}

func TestWhen_NoAttributesArePassed_Then_DefaultRuleIsServed(t *testing.T) {
	r, u := initialize()

	expectedVariation := "def"

	ulf, _ := engine.Evaluate(&r, u, nil)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringEqualsMatch_Then_RightVariationIsReturned(t *testing.T) {
	r, u := initialize()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "equals"
	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringNotEqualsMatch_Then_RightVariationIsReturned(t *testing.T) {
	r, u := initialize()

	expectedVariation := "neq"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "random"

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}
func TestWhen_StringStartsWithMatch_Then_RightVariationIsReturned(t *testing.T) {
	r, u := initialize()

	expectedVariation := "con"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "starts with this sentence and more"

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringDoesNotStartsWith_Then_RightVariationIsReturned(t *testing.T) {
	var mockedFlagSpecial, _ = ioutil.ReadFile("../../testdata/attributes/string_doesnotStartsWith.json")
	var responseDto dtos.Feature
	var userId = "user-"+ strconv.Itoa(rand.Intn(1000))
	json.Unmarshal(mockedFlagSpecial, &responseDto)

	expectedVariation := "gte"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "doesnot start with joker"

	ulf, _ := engine.Evaluate(&responseDto, userId, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

	// Start with joker, get default
	expectedVariation = "def"
	attributes["strAttr"] = "joker is at the start"
	ulf, _ = engine.Evaluate(&responseDto, userId, &attributes)
	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringEndsWithMatch_Then_RightVariationIsReturned(t *testing.T) {
	r, u := initialize()

	expectedVariation := "gt"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "this should match with ger that is liger"

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringDoesNotEndsWith_Then_RightVariationIsReturned(t *testing.T) {
	var mockedFlagSpecial, _ = ioutil.ReadFile("../../testdata/attributes/string_doesNotEndsWith.json")
	var responseDto dtos.Feature
	var userId = "user-"+ strconv.Itoa(rand.Intn(1000))
	json.Unmarshal(mockedFlagSpecial, &responseDto)

	expectedVariation := "lte"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "new york"

	ulf, _ := engine.Evaluate(&responseDto, userId, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

	// Start with joker, get default
	expectedVariation = "def"
	attributes["strAttr"] = "San Francisco"
	ulf, _ = engine.Evaluate(&responseDto, userId, &attributes)
	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringContainsMatch_Then_RightVariationIsReturned(t *testing.T) {
	r, u := initialize()

	expectedVariation := "lt"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "this contains dog"

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_StringNotContainsMatch_Then_RightVariationIsReturned(t *testing.T) {
	var mockedFlagSpecial, _ = ioutil.ReadFile("../../testdata/attributes/string_notcontains.json")
	var responseDto dtos.Feature
	var userId = "user-"+ strconv.Itoa(rand.Intn(1000))
	json.Unmarshal(mockedFlagSpecial, &responseDto)

	expectedVariation := "lte"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = "this sentence does not contain the word c-a-t"

	ulf, _ := engine.Evaluate(&responseDto, userId, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

	// If it contains cat, the default variation is returned
	expectedVariation = "def"
	attributes["strAttr"] = "his sentence contains the word cat."
	ulf, _ = engine.Evaluate(&responseDto, userId, &attributes)
	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_AttributeNameMatchButTypeIsNotString_Then_NoPanic(t *testing.T) {
	r, u := initialize()

	expectedVariation := "def"

	attributes := make(map[string]interface{})
	attributes["strAttr"] = 1.2 // use number instead of string

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}



