package attributes

import "testing"

func TestEquals_WhenEqualSets_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(4)

	runApplyAndCheck(flagSet, userSet, "EQ", true, t)
}

func TestEquals_WhenUserSetIsSuperset_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(8)

	runApplyAndCheck(flagSet, userSet, "EQ", false, t)
}


func TestEquals_WhenUserSetIsSubset_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(3)

	runApplyAndCheck(flagSet, userSet, "EQ", false, t)
}

func TestEquals_WhenUserSetIsEmptyOrNil_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := make(map[string]interface{})

	runApplyAndCheck(flagSet, userSet, "EQ", false, t)
	runApplyAndCheck(flagSet, nil, "EQ", false, t)
}