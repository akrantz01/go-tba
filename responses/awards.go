package responses

type Award struct {
	Name          string           `json:"name"`
	AwardType     int              `json:"award_type"`
	EventKey      string           `json:"event_key"`
	RecipientList []AwardRecipient `json:"recipient_list"`
	Year          int              `json:"year"`
}

type AwardRecipient struct {
	TeamKey string `json:"team_key"`
	Awardee string `json:"awardee"`
}
