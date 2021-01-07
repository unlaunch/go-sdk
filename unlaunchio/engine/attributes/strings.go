package attributes

import "strings"

func StringApply(value string, userValue string, operator string) bool {
	if operator == "EQ" {
		return value == userValue
	} else if operator == "NEQ" {
		return value != userValue
	} else if operator == "SW" {
		return strings.HasPrefix(userValue, value)
	} else if operator == "NSW" {
		return !strings.HasPrefix(userValue, value)
	} else if operator == "EW" {
		return strings.HasSuffix(userValue, value)
	} else if operator == "NEW" {
		return !strings.HasSuffix(userValue, value)
	} else if operator == "CON" {
		return strings.Contains(userValue, value)
	} else if operator == "NCON" {
		return !strings.Contains(userValue, value)
	}

	// Todo log invalid op warning
	return false
}

