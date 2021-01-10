package dtos

type Impression struct {
	FlagKey          string `json:"flagKey"`
	UserId           string `json:"userId"`
	VariationKey     string `json:"variationKey"`
	FlagStatus       string `json:"flagStatus"`
	EvaluationReason string `json:"evaluationReason"`
	MachineName      string `json:"machineName"`
}
