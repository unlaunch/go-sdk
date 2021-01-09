package attributes

type TargetingRuleCondition struct {
}

func (tr *TargetingRuleCondition) Apply(attrType string, val interface{}, userVal interface{}, op string) bool {
	switch attrType {
	case "boolean":
		return boolApply(val, userVal, op)
	case "string":
		return stringApply(val, userVal, op)
	case "number":
		return numberApply(val, userVal, op)
	case "date":
		return dateOrDateTimeApply(val, userVal, op, true)
	case "datetime":
		return dateOrDateTimeApply(val, userVal, op, false)
	default:
		// todo log "invalid condition type"
		return false
	}
}

func NewTargetingRule() *TargetingRuleCondition {
	return &TargetingRuleCondition{}
}