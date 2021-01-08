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

func TestWhen_FlagIsDisabled_Then_DefaultVariationIsReturned(t *testing.T) {

	var mockedDisabledFlag, _ = ioutil.ReadFile("../../testdata/disabledFlag.json")
	var responseDto dtos.TopLevelEnvelope
	err := json.Unmarshal(mockedDisabledFlag, &responseDto)

	if err != nil {
		t.Error("Error parsing mock flag JSON ", err)
	}

	ulf, err := engine.Evaluate(&responseDto.Data.Features[0], "d", nil)

	if err != nil {
		t.Error("evaluation threw error ", err)
	}

	if ulf.Variation.Key != "OFF_DEFAULT" {
		t.Error(fmt.Sprintf("variation should be 'OFF_DEFAULT'. It was '%s'", ulf.Variation.Key))
	}

	if ulf.Variation.Id != 17 {
		t.Error("Variation Id was not 18")
	}
}

func TestWhen_FlagIsEnabled_Then_DefaultRuleIsReturned(t *testing.T) {
	var mockedDisabledFlag, _ = ioutil.ReadFile("../../testdata/enabledFlagWithAllowList.json")
	var responseDto dtos.TopLevelEnvelope
	err := json.Unmarshal(mockedDisabledFlag, &responseDto)

	if err != nil {
		t.Error("Error parsing mock flag JSON ", err)
	}

	ulf, err := engine.Evaluate(&responseDto.Data.Features[0], "d", nil)

	if err != nil {
		t.Error("evaluation threw error ", err)
	}

	if ulf.Variation.Key != "ON_DEFAULT_RULE" {
		t.Error(fmt.Sprintf("variation should be 'ON_DEFAULT_RULE'. It was '%s'", ulf.Variation.Key))
	}

	if ulf.Variation.Id != 18 {
		t.Error("Variation Id was not 18")
	}
}

func TestWhen_FlagIsEnabledAndUserIsInAllowList_Then_AllowListVariationIsReturned(t *testing.T) {
	var mockedDisabledFlag, _ = ioutil.ReadFile("../../testdata/enabledFlagWithAllowList.json")
	var responseDto dtos.TopLevelEnvelope
	err := json.Unmarshal(mockedDisabledFlag, &responseDto)

	if err != nil {
		t.Error("Error parsing mock flag JSON ", err)
	}

	ulf, err := engine.Evaluate(&responseDto.Data.Features[0], "user123", nil)

	if err != nil {
		t.Error("evaluation threw error ", err)
	}
	if ulf.Variation.Key != "off" {
		t.Error(fmt.Sprintf("variation should be 'off'. It was '%s'", ulf.Variation.Key))
	}

	if ulf.Variation.Id != 17 {
		t.Error("Variation Id was not 17")
	}
}

func TestWhen_RollOutIsEnabledForAUser_Then_VariationSameAssignedConsistently(t *testing.T) {
	var mockedDisabledFlag, _ = ioutil.ReadFile("../../testdata/flag1WithDefaultRuleRollout.json")
	var responseDto dtos.Feature
	json.Unmarshal(mockedDisabledFlag, &responseDto)

		sort.Sort(dtos.ByRulePriority(responseDto.Rules))
		for _, rule := range responseDto.Rules {
			sort.Sort(dtos.ByVariationId(rule.Rollout))
		}


	// This is the same flag as above but splits are unordered
	var mockedDisabledFlagReversed, _ = ioutil.ReadFile("../../testdata/flag1WithDefaultRuleRolloutReversedOrder.json")
	var responseDtoReversed dtos.Feature
	json.Unmarshal(mockedDisabledFlagReversed, &responseDtoReversed)

	for i := 0; i<10; i++ {
		userId := "user-"+ strconv.Itoa(rand.Intn(1000))
		ulf, _ := engine.Evaluate(&responseDto, userId, nil)
		variation := ulf.Variation.Key

		// Evaluate using both ordered and (reversed) order data
		// Result should be the same
		ulf, _ = engine.Evaluate(&responseDtoReversed, userId, nil)

		if variation != ulf.Variation.Key {
			t.Error(fmt.Sprintf("expected variation %s actual variation %s on the iteration #%d",
				variation, ulf.Variation.Key, i+1))
			return
		}
	}
}

func TestWhen_RollOutIsEnabled_Then_VariationIsAllocatedByBucketing(t *testing.T) {
	var mockedDisabledFlag, _ = ioutil.ReadFile("../../testdata/flag1WithDefaultRuleRollout.json")
	var responseDto dtos.Feature
	err := json.Unmarshal(mockedDisabledFlag, &responseDto)

	if err != nil {
		t.Error("Error parsing mock flag JSON ", err)
	}

	countOn, countOff := 0, 0
	for i:= 0; i<50; i++ {
		ulf, err := engine.Evaluate(&responseDto, "user-" + strconv.Itoa(i), nil)

		if err != nil {
			t.Error("evaluation threw error ", err)
		}

		if ulf.Variation.Key == "on" {
			countOn++
		} else if ulf.Variation.Key == "off" {
			countOff++
		} else {
			t.Error("Only on and off variation were expected.")
		}
	}

	if countOn < 15 || countOff < 15 {
		t.Error(fmt.Sprintf("Variation bucketing distribution was not even. on: %d, off: %d", countOn, countOff))
	}
}

