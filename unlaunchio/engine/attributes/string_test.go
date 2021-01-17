package attributes

import "testing"

func TestStringOperators(t *testing.T) {
	flagString := "ahiman"

	if !stringApply(flagString, "ahiman", "EQ") {
		t.Errorf("string equal should match")
	}

	if !stringApply(flagString, "abc", "NEQ") {
		t.Errorf("string not equal should match")
	}

	if stringApply(flagString, "abc", "CON") {
		t.Errorf("contains shouldn't match")
	}

	if !stringApply(flagString, "fishisahiman", "CON") {
		t.Errorf("contains should match")
	}

	if !stringApply(flagString, "this also contains ahiman so it should match", "CON") {
		t.Errorf("contains should match")
	}

	if !stringApply(flagString, "doesntcontain", "NCON") {
		t.Errorf("not contains should match")
	}

	if !stringApply(flagString, "ahimanatstart", "SW") {
		t.Errorf("starts with should match")
	}

	if !stringApply(flagString, "noahimanatstart", "NSW") {
		t.Errorf("not starts with should match")
	}

	if !stringApply(flagString, "endswithahiman", "EW") {
		t.Errorf("ends with should match")
	}

	if !stringApply(flagString, "doesntendswithahiman.", "NEW") {
		t.Errorf("does not ends with should match")
	}

	if stringApply(flagString, nil, "SW") {
		t.Errorf("nil string shouldn't match")
	}

	if stringApply(flagString, "", "EW") {
		t.Errorf("empty string shouldn't match")
	}

}
