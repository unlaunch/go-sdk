package attributes

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAllOf_WhenEqualSets_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := getSetOfNumbersUpTo(3)

	runApplyAndCheck(flagSet, userSet, "AO", true, t)
}

func TestAllOf_WhenUserSetIsSuperset_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := getSetOfNumbersUpTo(4)

	runApplyAndCheck(flagSet, userSet, "AO", true, t)
}

func TestAllOf_WhenUserSetIsSubset_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := getSetOfNumbersUpTo(2)

	runApplyAndCheck(flagSet, userSet, "AO", false, t)
}

func TestAllOf_WhenUserSetIsEmpty_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := make(map[string]interface{})

	runApplyAndCheck(flagSet, userSet, "AO", false, t)
}

func TestAllOf_WhenUserSetIsNil_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	runApplyAndCheck(flagSet, nil, "AO", false, t)

}


func runApplyAndCheck(set1 string, set2 map[string]interface{}, op string, expected bool, t *testing.T) {
	b := setApply(set1, set2, op)

	if b != expected {
		t.Error(fmt.Sprintf("failed for operator %s. Expected %v Got %v", op, expected, b))
	}
}

func getSetOfNumbersUpTo(n int) map[string]interface{} {
	setnums := make(map[string]interface{})

	var i int64
	for i = 1; i<= int64(n); i++ {
		setnums[strconv.FormatInt(i, 32)] = nil
	}

	return setnums
}