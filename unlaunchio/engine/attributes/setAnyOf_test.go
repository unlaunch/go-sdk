package attributes

import "testing"

func TestAnyOf_WhenEqualSets_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(4)

	runApplyAndCheck(flagSet, userSet, "HA", true, t)
}

func TestAnyOf_WhenUserSetIsSuperSet_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(5)

	runApplyAndCheck(flagSet, userSet, "HA", true, t)
}

func TestAnyOf_WhenUserSetIsSubSet_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(2)

	runApplyAndCheck(flagSet, userSet, "HA", true, t)
}

func TestAnyOf_WhenUserSetIsSubSetUnordered_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := make(map[string]interface{})
	userSet["3"] = nil
	userSet["2"] = nil
	userSet["1"] = nil
	userSet["10"] = nil

	runApplyAndCheck(flagSet, userSet, "HA", true, t)
}

func TestAnyOf_WhenUserSetIsDisjoint_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := make(map[string]interface{})
	userSet["4"] = nil
	userSet["5"] = nil

	runApplyAndCheck(flagSet, userSet, "HA", false, t)
}

func TestAnyOf_WhenUserSetIsEmptyOrNil_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := make(map[string]interface{})

	runApplyAndCheck(flagSet, userSet, "HA", false, t)
	runApplyAndCheck(flagSet, nil, "HA", false, t)
}
