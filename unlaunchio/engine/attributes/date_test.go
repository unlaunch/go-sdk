package attributes

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDateOperators(t *testing.T) {

	// Flag always contain datetime in Java milliseconds since epoch
	dateAug152009 := "1250298753000" // August 15, 2009 1:25:33 AM
	dateAug162009 := "1250423999000" // Aug 16, 2009
	dateAug172009 := "1250472333000" // August 17, 2009 1:25:33 AM

	var v int64

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug162009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "EQ", true) {
		t.Errorf("EQ dates should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug172009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "GT", true) {
		t.Errorf("GT than date should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug152009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "LT", true) {
		t.Errorf("LT date should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug162009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "GTE", true) {
		t.Errorf("GTE date should match")
	}

	if !dateOrDateTimeApply(dateAug162009, time.Now().UTC().Unix(), "GTE", true) {
		t.Errorf("GTE date should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug162009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "LTE", true) {
		t.Errorf("LTE equal date should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(dateAug152009, "000"), 10, 64)
	if !dateOrDateTimeApply(dateAug162009, v, "LTE", true) {
		t.Errorf("LTE date should match")
	}
}

func TestDateTimeOperators(t *testing.T) {
	// Some times
	time0 := "1250298753000" // August 15, 2009 1:12:33 AM
	time1 := "1250299113000" // August 15, 2009 1:18:33 AM
	time2 := "1250299533000" // August 15, 2009 1:25:33 AM

	var v int64
	v, _ = strconv.ParseInt(strings.TrimSuffix(time1, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "EQ", false) {
		t.Errorf("EQ datetime should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(time2, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "GT", false) {
		t.Errorf("GT datetime should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(time0, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "LT", false) {
		t.Errorf("LT datetime should match")
	}

	if !dateOrDateTimeApply(time1, time.Now().UTC().Unix(), "GTE", false) {
		t.Errorf("GTE datetime should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(time1, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "GTE", false) {
		t.Errorf("GTE equal datetime should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(time1, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "LTE", false) {
		t.Errorf("LTE equal datetime should match")
	}

	v, _ = strconv.ParseInt(strings.TrimSuffix(time0, "000"), 10, 64)
	if !dateOrDateTimeApply(time1, v, "LTE", false) {
		t.Errorf("LTE datetime should match")
	}
}
