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
	"time"
)

func initializeDateAttrFile1() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/dateEqGtLt.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func initializeDateAttrFile2() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/dateGteLte.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationId(rule.Rollout))
	}

	userId := "user-"+ strconv.Itoa(rand.Intn(1000))

	return responseDto, userId
}

func TestWhen_DateEqualsMatch(t *testing.T) {
	r, u := initializeDateAttrFile1()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	attributes["dateAttr"] = time.Date(
		2021, 1, 2, 10, 12, 45, 1000, time.UTC).Unix()

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_DateGreaterThanMatch(t *testing.T) {
	r, u := initializeDateAttrFile1()

	expectedVariation := "gt"

	attributes := make(map[string]interface{})
	attributes["dateAttr"] = time.Now().UTC().Unix() // original time is Jan 3, 2021 UTC

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_DateLessThanMatch(t *testing.T) {
	r, u := initializeDateAttrFile1()

	expectedVariation := "lt"

	attributes := make(map[string]interface{})
	attributes["dateAttr"] = time.Date(
		2020, 12, 31, 9, 12, 45, 1000, time.UTC).Unix()

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}

func TestWhen_DateLessThanEqualsMatch(t *testing.T) {
	r, u := initializeDateAttrFile2()

	expectedVariation := "eq" // I choose wrong variation in JSON

	attributes := make(map[string]interface{})
	attributes["dateAttr"] = time.Date(
		2021, 1, 14, 10, 12, 45, 1000, time.UTC).Unix()

	ulf, _ := engine.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}