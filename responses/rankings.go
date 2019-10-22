package responses

type WinLossTieRecord struct {
	Losses int `json:"losses"`
	Ties   int `json:"ties"`
	Wins   int `json:"wins"`
}

type SortOrder struct {
	Precision int    `json:"precision"`
	Name      string `json:"name"`
}

type Ranking struct {
	MatchesPlayed        int              `json:"matches_played"`
	QualificationAverage float64          `json:"qual_average"`
	Record               WinLossTieRecord `json:"record"`
	Rank                 int              `json:"rank"`
	Disqualifications    int              `json:"dq"`
	TeamKey              string           `json:"team_key"`
	SortOrders           []float64        `json:"sort_orders"`
	ExtraStats           []float64        `json:"extra_stats"`
}
