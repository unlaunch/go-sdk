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
	"time"
)

func initializeDateTimeAttrFile1() (dtos.Feature, string) {
	var mockedFlagNum, _ = ioutil.ReadFile("../../testdata/attributes/dateTimeEqGtLt.json")

	var responseDto dtos.Feature
	json.Unmarshal(mockedFlagNum, &responseDto)
	sort.Sort(dtos.ByRulePriority(responseDto.Rules))
	for _, rule := range responseDto.Rules {
		sort.Sort(dtos.ByVariationID(rule.Rollout))
	}

	userID := "user-" + strconv.Itoa(rand.Intn(1000))

	return responseDto, userID
}

// TODO: Uncomment this when Mahrukh fixes on backend
//func TestWhen_DateTimeGreaterThanMatch(t *testing.T) {
//	r, u := initializeDateTimeAttrFile1()
//
//	expectedVariation := "gt"
//
//	attributes := make(map[string]interface{})
//	attributes["dateTimeAttr"] = time.Date(
//		2021, 1, 2, 22, 33, 0, 0, time.Local).Unix()
//
//	ulf, _ := evaluator.Evaluate(&r, u, &attributes)
//
//	if ulf.Variation != expectedVariation {
//
//		t.Log(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
//	}
//}

func TestWhen_DateTimeLessThanMatch(t *testing.T) {
	r, u := initializeDateTimeAttrFile1()

	expectedVariation := "eq"

	attributes := make(map[string]interface{})
	attributes["dateTimeAttr"] = time.Date(
		2020, 5, 7, 13, 35, 34, 899, time.UTC).Unix()

	ulf, _ := evaluator.Evaluate(&r, u, &attributes)

	if ulf.Variation != expectedVariation {
		t.Error(fmt.Sprintf("Expected '%s'. Got '%s'", expectedVariation, ulf.Variation))
	}
}
