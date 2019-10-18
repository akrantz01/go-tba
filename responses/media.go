package responses

type Webcast struct {
	Type string `json:"type"`
	Channel string `json:"channel"`
	File string `json:"file"`
}

type Media struct {
	Key string `json:"key"`
	Type string `json:"type"`
	ForeignKey string `json:"foreign_key"`
	Details map[string]interface{} `json:"details"`
	Preferred bool `json:"preferred"`
	DirectUrl string `json:"direct_url"`
	ViewUrl string `json:"view_url"`
}
