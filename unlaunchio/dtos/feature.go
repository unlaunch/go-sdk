package dtos

type Feature struct {
	Key 			string `json:"key"`
	Name 			string `json:"name"`
	State 			string `json:"state"`
	Variations 		[]Variation
	OffVariation 	int `json:"offVariation"`
	Rules 			[]Rule
}

func (f *Feature) Enabled() bool {
	if f.State == "ACTIVE" {
		return true
	} else {
		return false
	}
}

func (f *Feature) DefaultRule() *Rule {
	for _, rule := range f.Rules {
		if rule.IsDefault {
			return &rule
		}
	}

	// TODO log error; flag format is wrong
	return nil
}

func (f *Feature) VariationById(id int) *Variation {
	for _, variation := range f.Variations {
		if id == variation.Id {
			return &variation
		}
	}
	return nil

}

type Variation struct {
	Id 			int
	Key 		string
	AllowList 	string
}

type Rollout struct {
	Id 					int
	VariationId 		int
	RolloutPercentage 	int
}

type Rule struct {
	IsDefault 	bool
	Priority 	uint64
	Rollout 	[]Rollout `json:"splits"`
}

type Data struct {
	ProjectName string
	EnvName 	string
	Features 	[]Feature `json:"flags"`
}

type TopLevelEnvelope struct {
	Data Data
}

type UnlaunchFeature struct {
	Feature 		string
	Variation 		*Variation
	EvaluationReason string
}