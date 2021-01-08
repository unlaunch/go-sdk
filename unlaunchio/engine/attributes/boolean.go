package attributes

func BooleanApply(value bool, userValue bool, operator string) bool {
	if operator == "EQ" {
		return userValue == value
	} else if operator == "NEQ" {
		return userValue != value
	}

	// Todo log invalid op warning
	return false
}