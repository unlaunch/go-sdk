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

type ByVariationId []Rollout

func (ro ByVariationId) Len() int           { return len(ro) }
func (ro ByVariationId) Less(i, j int) bool { return ro[i].VariationId < ro[j].VariationId }
func (ro ByVariationId) Swap(i, j int)      { ro[i], ro[j] = ro[j], ro[i] }

type Rule struct {
	IsDefault 	bool
	Priority 	int
	Rollout 	[]Rollout `json:"splits"`
	Conditions  []Condition
}

type ByRulePriority []Rule

func (r ByRulePriority) Len() int           { return len(r) }
func (r ByRulePriority) Less(i, j int) bool { return r[i].Priority < r[j].Priority }
func (r ByRulePriority) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

type Condition struct {
	Id int
	AttributeId int
	Attribute string
	Type string
	Value string
	Op string
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