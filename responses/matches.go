package responses

type MatchSimple struct {
	Key string `json:"key"`
	CompetitionLevel string `json:"comp_level"`
	SetNumber int `json:"set_number"`
	MatchNumber int `json:"match_number"`
	Alliances struct{
		Red MatchAlliance `json:"red"`
		Blue MatchAlliance `json:"blue"`
	} `json:"alliances"`
	WinningAlliance string `json:"winning_alliance"`
	EventKey string `json:"event_key"`
	Time int64 `json:"time"`
	PredictedTime int64 `json:"predicted_time"`
	ActualTime int64 `json:"actual_time"`
}

type Match struct {
	Key string `json:"key"`
	CompetitionLevel string `json:"comp_level"`
	SetNumber int `json:"set_number"`
	MatchNumber int `json:"match_number"`
	Alliances struct{
		Red MatchAlliance `json:"red"`
		Blue MatchAlliance `json:"blue"`
	} `json:"alliances"`
	WinningAlliance string `json:"winning_alliance"`
	EventKey string `json:"event_key"`
	Time int64 `json:"time"`
	PredictedTime int64 `json:"predicted_time"`
	ActualTime int64 `json:"actual_time"`
	PostResultTime int64 `json:"post_result_time"`
	ScoreBreakdown ScoringBreakdown
	Videos []struct{
		Type string `json:"type"`
		Key string `json:"key"`
	} `json:"videos"`
}

type MatchAlliance struct {
	Score int `json:"score"`
	TeamKeys []string `json:"team_keys"`
	SurrogateTeamKeys []string `json:"surrogate_team_keys"`
	DisqualifiedTeamKeys []string `json:"dq_team_keys"`
}
