package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strings"
)

func stringApply(val interface{}, userVal interface{}, op string) bool {
	v := val.(string)
	uv, err := util.ConvertToString(userVal)

	if err != nil {
		// TODO log warning that name matches but type is not bool
		return false
	}

	if op == "EQ" {
		return v == userVal
	} else if op == "NEQ" {
		return v != userVal
	} else if op == "SW" {
		return strings.HasPrefix(uv, v)
	} else if op == "NSW" {
		return !strings.HasPrefix(uv, v)
	} else if op == "EW" {
		return strings.HasSuffix(uv, v)
	} else if op == "NEW" {
		return !strings.HasSuffix(uv, v)
	} else if op == "CON" {
		return strings.Contains(uv, v)
	} else if op == "NCON" {
		return !strings.Contains(uv, v)
	}

	// Todo log invalid op warning
	return false
}

