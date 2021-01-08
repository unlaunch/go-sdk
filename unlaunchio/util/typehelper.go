package util

import (
	"errors"
	"fmt"
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

// IsBool
func IsBool(attr interface{}) bool {
	switch v := attr.(type) {
	case bool:
		return true
	default:
		fmt.Sprintf("%v", v)
		return false
	}
}

func IsString(attr interface{}) bool {
	switch v := attr.(type) {
	case string:
		return true
	default:
		fmt.Sprintf("%v", v)
		return false
	}
}

func IsNumber(attr interface{}) bool {
	switch v := attr.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		fmt.Sprintf("%v", v)
		return false
	}
}