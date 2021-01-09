package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strconv"
)

func numberApply(val interface{}, userVal interface{}, op string) bool {

	v, _ := strconv.ParseFloat(val.(string), 64)
	uv, err := util.GetFloat64(userVal)

	if err != nil {
		// TODO log warning that name matches but type is not bool
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
		// Todo log invalid op warning
		return false

	}
}