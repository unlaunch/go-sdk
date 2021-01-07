package attributes

func BooleanApply(value bool, userValue bool, operator string) bool {
	if operator == "EQ" {
		return applyEquals(value, userValue)
	} else if operator == "NEQ" {
		return applyNotEquals(value, userValue)
	}

	// Todo log invalid op warning
	return false
}

func applyEquals(value bool, userValue bool) bool {
	return value == userValue
}

func applyNotEquals(value bool, userValue bool) bool {
	return value != userValue
}