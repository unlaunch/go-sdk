package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strconv"
	"time"
)

func dateOrDateTimeApply(val interface{}, userVal interface{}, op string, discardTime bool) bool {
	// This is the value in Java
	v, _ := strconv.ParseInt(val.(string), 10, 64)
	v = javaTimeToEpoc(v)

	uv, err := util.GetInt64(userVal)

	if discardTime {
		v = tsWithZeroTime(v)
		uv = tsWithZeroTime(uv)
	}

	if err != nil {
		// TODO log warning that name matches but type is not right
		return false
	}

	switch op {
	case "EQ":
		return uv == v
	case "GT":
		return uv > v
	case "GTE":
		return uv >= v
	case "LT":
		return uv < v
	case "LTE":
		return uv <= v
	case "NEQ":
		return uv != v
	default:
		return false
	}
}

func javaTimeToEpoc(ts int64) int64 {
	return ts / 1000
}

func tsWithZeroTime(ts int64) int64 {
	t := time.Unix(ts, 0).UTC() // seconds to time
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return rounded.Unix()
}

