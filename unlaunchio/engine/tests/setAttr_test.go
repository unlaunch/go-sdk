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

func initializeSetAttr1() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/setHasAnyHasAll.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func initializeSetAttr2() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/setDoesntHasAnyHasAll.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func TestWhen_SetHasAnyMatch(t *testing.T) {
	r, u := initializeSetAttr1()

	expectedVariation := "on"

	characters := make(map[string]interface{})
	characters["batman"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

	//-- no attribute

	expectedVariation = "defrule"

	characters2 := make(map[string]interface{})
	characters2["liono"] = nil

	attributes2 := make(map[string]interface{})
	attributes2["setAttr"] = characters2

	ulf, _ = engine.Evaluate(&r, u, &attributes2)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

}

func TestWhen_SetHasAnyDontMatch(t *testing.T) {
	r, u := initializeSetAttr1()

	expectedVariation := "defrule"

	characters := make(map[string]interface{})
	characters["liono"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

}

func TestWhen_SetHasAllOfMatch(t *testing.T) {
	r, u := initializeSetAttr1()

	expectedVariation := "off"

	characters := make(map[string]interface{})
	characters["superman"] = nil
	characters["wonderwoman"] = nil
	characters["greenlantern"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}

}

func TestWhen_SetHasAllOfDontMatch(t *testing.T) {
	r, u := initializeSetAttr1()

	expectedVariation := "defrule"

	characters := make(map[string]interface{})
	characters["superman"] = nil
	characters["wonderwoman"] = nil
	characters["catwoman"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_SetDoesNotHaveAnyOfMatch(t *testing.T) {
	r, u := initializeSetAttr2()

	expectedVariation := "on"

	characters := make(map[string]interface{})
	characters["superman"] = nil
	characters["wonderwoman"] = nil
	characters["catwoman"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}

func TestWhen_SetDoesNotHaveAnyOfDontMatch(t *testing.T) {
	r, u := initializeSetAttr2()

	expectedVariation := "defrule"

	characters := make(map[string]interface{})
	characters["superman"] = nil
	characters["tigera"] = nil
	characters["catwoman"] = nil

	attributes := make(map[string]interface{})
	attributes["setAttr"] = characters

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation.Key != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation.Key))
	}
}