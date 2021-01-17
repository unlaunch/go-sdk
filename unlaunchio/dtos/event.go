package dtos

type Event struct {
	CreatedTime  int64                  `json:"createdTime"`
	Type         string                 `json:"type"`
	Key          string                 `json:"key"`
	Properties   map[string]interface{} `json:"properties"`
	Sdk          string                 `json:"sdk"`
	SdkVersion   string                 `json:"sdkVersion"`
	SecondaryKey string                 `json:"secondaryKey"`
	Impression
}
