package attributes

func NumberApply(value float64, userValue float64, operator string) bool {
	if operator == "EQ" {
		return value == userValue
	} else if operator == "NEQ" {
		return value != userValue
	} else if operator == "GT" {
		return userValue > value
	} else if operator == "GTE" {
		return userValue >= value
	} else if operator == "LT" {
		return userValue < value
	} else if operator == "LTE" {
		return userValue <= value
	}

	// Todo log invalid op warning
	return false
}