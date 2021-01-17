package attributes

import "testing"

func TestPartOf_WhenEqualSets_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(4)

	runApplyAndCheck(flagSet, userSet, "PO", true, t)
}

func TestPartOf_WhenUserSetIsSuperset_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(5)

	runApplyAndCheck(flagSet, userSet, "PO", false, t)
}

func TestPartOf_WhenUserSetIsSubset_ShouldMatch(t *testing.T) {
	flagSet := "1, 2, 3, 4"

	userSet := getSetOfNumbersUpTo(3)

	runApplyAndCheck(flagSet, userSet, "PO", true, t)
}

func TestPartOf_WhenUserSetIsEmptyOrNil_ShouldNotMatch(t *testing.T) {
	flagSet := "1, 2, 3"

	userSet := make(map[string]interface{})

	runApplyAndCheck(flagSet, userSet, "PO", false, t)
	runApplyAndCheck(flagSet, nil, "PO", false, t)
}
