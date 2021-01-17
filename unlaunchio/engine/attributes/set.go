package attributes

import (
	"github.com/unlaunch/go-sdk/unlaunchio/engine/datatypes/set"
	"github.com/unlaunch/go-sdk/unlaunchio/util"
	"strings"
)

func setApply(val interface{}, userVal interface{}, op string) bool {

	vals, _ := util.ConvertToString(val)
	flagSet := set.NewSet()
	for _, item := range strings.Split(vals, ",") {
		flagSet.Add(strings.TrimSpace(item))
	}

	uValues, err := util.CovertToMap(userVal)
	if err != nil {
		// TODO log that user must pass map[string]interface{}
		return false
	}

	userValuesSet := set.NewSet()
	for k := range uValues {
		userValuesSet.Add(strings.TrimSpace(k))
	}

	switch op {
	case "AO": // All of
		return userValuesSet.IsSuperset(flagSet)
	case "NAO":
		return !userValuesSet.IsSuperset(flagSet)
	case "HA": // Has any of
		i := userValuesSet.Intersect(flagSet)
		return i.Cardinality() > 0
	case "NHA": // Not Has any of
		i := userValuesSet.Intersect(flagSet)
		return i.Cardinality() == 0
	case "EQ": // Equals
		return flagSet.Equal(userValuesSet)
	case "NEQ": // Equals
		return !flagSet.Equal(userValuesSet)
	case "PO": // Part of
		if userValuesSet.Cardinality() < 1 {
			return false
		}
		return userValuesSet.IsSubset(flagSet)
	case "NPO": // Not Part of
		if userValuesSet.Cardinality() < 1 {
			return false
		}
		return !userValuesSet.IsSubset(flagSet)
	default:
		// Todo log invalid op warning
		return false
	}

	// Todo log invalid op warning
	return false
}
