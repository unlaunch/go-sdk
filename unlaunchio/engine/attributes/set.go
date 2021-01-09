package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strings"
)

func setApply(val interface{}, userVal interface{}, op string) bool {

	jsonValue, _ := util.GetString(val)
	jsonValues := strings.Split(jsonValue, ",")
	jsonMap := make(map[string]interface{})

	for _, v := range jsonValues {
		jsonMap[v] = nil
	}


	userValuesMap, _ := util.GetSet(userVal)


	switch op {
	case "HA":
		for _, val := range jsonValues {
			if _, ok := userValuesMap[val]; ok {
				return true
			}
		}
	case "AO":
		for key, _ := range jsonMap {
			if _, ok := userValuesMap[key]; !ok {
				return false
			}
		}
		return true
	case "NHA":
		for _, val := range jsonValues {
			if _, ok := userValuesMap[val]; ok {
				return false
			}
		}
		return true
	default:
		// Todo log invalid op warning
		return false
	}

	// Todo log invalid op warning
	return false
}