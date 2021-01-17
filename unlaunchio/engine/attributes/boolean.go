package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strconv"
)

func boolApply(val interface{}, userVal interface{}, op string) bool {
	v, _ := strconv.ParseBool(val.(string))
	uv, err := util.ConvertToBool(userVal)
	if err != nil {
		// TODO log warning that name matches but type is not bool
		return false
	}

	if op == "EQ" {
		return uv == v
	} else if op == "NEQ" {
		return uv != v
	}

	// Todo log invalid op warning; flag json broken
	return false
}
