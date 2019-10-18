package responses

type DistrictList struct {
	Abbreviation string `json:"abbreviation"`
	DisplayName  string `json:"display_name"`
	Key          string `json:"key"`
	Year         int    `json:"year"`
}

type DistrictRanking struct {
	TeamKey     string                       `json:"team_key"`
	Rank        int                          `json:"rank"`
	RookieBonus int                          `json:"rookie_bonus"`
	PointTotal  int                          `json:"point_total"`
	EventPoints []DistrictRankingEventPoints `json:"event_points"`
}

type DistrictRankingEventPoints struct {
	DistrictCompetition bool   `json:"district_cmp"`
	Total               int    `json:"total"`
	AlliancePoints      int    `json:"alliance_points"`
	EliminationPoints   int    `json:"elimination_points"`
	AwardPoints         int    `json:"award_points"`
	EventKey            string `json:"event_key"`
	QualificationPoints int    `json:"qualification_points"`
}

type EventDistrictPoints struct {
	Points map[string]struct {
		Total               int `json:"total"`
		AlliancePoints      int `json:"alliance_points"`
		EliminationPoints   int `json:"elim_points"`
		AwardPoints         int `json:"award_points"`
		QualificationPoints int `json:"qual_points"`
	} `json:"points"`
	TieBreakers map[string]struct {
		HighestQualificationScore []int `json:"highest_qual_score"`
		QualificationWins         int   `json:"qual_wins"`
	} `json:"tie_breakers"`
}
