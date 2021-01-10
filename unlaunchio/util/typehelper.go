package util

import (
	"errors"
	"fmt"
	"reflect"
)

// ConvertToFloat64 There has to be better way to do this
func ConvertToFloat64(attr interface{}) (float64, error) {
	uVal := 0.0

	// Todo: Move to util
	switch v := attr.(type) {
	case float64:
		uVal =attr.(float64)
	case float32:
		uVal = float64(attr.(float32))
	case int8:
		uVal = float64(attr.(int8))
	case int16:
		uVal = float64(attr.(int16))
	case int32:
		uVal = float64(attr.(int32))
	case int64:
		uVal = float64(attr.(int64))
	case int:
		uVal = float64(attr.(int))
	case uint8:
		uVal = float64(attr.(uint8))
	case uint16:
		uVal = float64(attr.(uint16))
	case uint32:
		uVal = float64(attr.(uint32))
	case uint64:
		uVal = float64(attr.(uint64))
	case uint:
		uVal = float64(attr.(uint))
	default:
		// TODO log error
		fmt.Sprintf("%v", v)
		return 0.0, errors.New("not a number")
	}

	return uVal, nil
}

// ConvertToBool
func ConvertToBool(attr interface{}) (bool, error) {
	switch v := attr.(type) {
	case bool:
		return attr.(bool), nil
	default:
		fmt.Sprintf("%v", v)
		return false, errors.New("not boolean")
	}
}

func ConvertToString(attr interface{}) (string, error) {
	switch v := attr.(type) {
	case string:
		return attr.(string), nil
	default:
		fmt.Sprintf("%v", v)
		return "", errors.New("not string")
	}
}

func ConvertToInt64(attr interface{}) (int64, error) {
	switch v := attr.(type) {
	case int64:
		return attr.(int64), nil
	case int:
		return int64(attr.(int)), nil
	case int32:
		return int64(attr.(int32)), nil
	default:
		fmt.Sprintf("%v", v)
		return 0, errors.New("not int64")
	}
}


func CovertToMap(attr interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(attr)
	if v.Kind() == reflect.Map {
		return attr.(map[string]interface{}), nil
	}

	return nil, errors.New("not map")


}
