package attributes

import "testing"

func TestNumberOperators(t *testing.T) {
	flagNumber := "1"

	if !numberApply(flagNumber, 1.0, "EQ") {
		t.Errorf("equal should match")
	}

	if !numberApply(flagNumber, 100, "NEQ") {
		t.Errorf("not equal should match")
	}

	if !numberApply(flagNumber, 2, "GT") {
		t.Errorf("greater than should match")
	}

	if !numberApply(flagNumber, 1, "GTE") {
		t.Errorf("greater than or equals should match")
	}

	if !numberApply(flagNumber, -1.1, "LT") {
		t.Errorf("less than should match")
	}

	if !numberApply(flagNumber, -1.0, "LTE") {
		t.Errorf("less than or equals should match")
	}

	if numberApply(flagNumber, nil, "EQ") {
		t.Errorf("nil should return false")
	}
}
