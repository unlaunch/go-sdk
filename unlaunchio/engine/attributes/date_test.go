package attributes

import (
	"testing"
	"time"
)

func TestDateAndDateTimeOperators(t *testing.T) {

	// Flag always contain datetime in Java milliseconds since epoch
	flagDateTime := "1232071462000" // Aug 16, 2009

	if !dateOrDateTimeApply(flagDateTime, 1232071462, "EQ", true) {
		t.Errorf("equal dates should match")
	}

	if !dateOrDateTimeApply(flagDateTime, 1232071462, "EQ", false) {
		t.Errorf("equal datetime should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix(), "GT", true) {
		t.Errorf("greater than date should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix(), "GT", false) {
		t.Errorf("greater than datetime should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix(), "GTE", true) {
		t.Errorf("greater than date should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix(), "GTE", false) {
		t.Errorf("greater than or equals datetime should match")
	}

	// Subtract ~15 years
	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix() - 473412761 , "LT", true) {
		t.Errorf("less than date should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix() - 473412761, "LT", false) {
		t.Errorf("less than datetime should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix() - 473412761, "LTE", true) {
		t.Errorf("less than or equals date should match")
	}

	if !dateOrDateTimeApply(flagDateTime, time.Now().UTC().Unix() - 473412761, "LTE", false) {
		t.Errorf("less than or equals datetime should match")
	}



}

