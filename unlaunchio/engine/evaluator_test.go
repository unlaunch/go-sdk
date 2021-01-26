package engine

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"github.com/unlaunch/go-sdk/unlaunchio/dtos"
)

var e = SimpleEvaluator{}

func TestWhen_TargetRule_MultipleConditionsMatch(t *testing.T) {

	var mockedFlag, _ = ioutil.ReadFile("../testdata/flagWithMultipleConditionRule.json")
	var responseDto *dtos.Feature

	attributes := make(map[string]interface{})
	attributes["device"] = "ABCS"
	attributes["registered"] = true

	json.Unmarshal(mockedFlag, &responseDto)

	v := e.matchTargetingRules(responseDto, "user123", &attributes)

	if v == nil {
		t.Errorf("Both conditions are not matched")
	}

	if v != nil && v.Key != "on" {
		t.Errorf("Expected '%s'. Got '%s'", "on", v.Key)
	}
}
