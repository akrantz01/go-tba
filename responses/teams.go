package responses

type TeamSimple struct {
	Key           string `json:"key"`
	TeamNumber    int    `json:"team_number"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
	City          string `json:"city"`
	StateProvince string `json:"state_prov"`
	Country       string `json:"country"`
}

type Team struct {
	Key               string            `json:"key"`
	TeamNumber        int               `json:"team_number"`
	Nickname          string            `json:"nickname"`
	Name              string            `json:"name"`
	City              string            `json:"city"`
	StateProvince     string            `json:"state_prov"`
	Country           string            `json:"country"`
	Address           string            `json:"address"`
	PostalCode        string            `json:"postal_code"`
	GoogleMapsPlaceId string            `json:"gmaps_place_id"`
	GoogleMapsUrl     string            `json:"gmaps_url"`
	Latitude          float64           `json:"lat"`
	Longitude         float64           `json:"long"`
	LocationName      string            `json:"location_name"`
	Website           string            `json:"website"`
	RookieYear        int               `json:"rookie_year"`
	HomeChampionship  map[string]string `json:"home_championship"`
}

type TeamRobot struct {
	Year      int    `json:"year"`
	RobotName string `json:"robot_name"`
	Key       string `json:"key"`
	TeamKey   string `json:"team_key"`
}

type TeamEventStatus struct {
	Qualification  TeamEventStatusRank     `json:"qual"`
	Alliance       TeamEventStatusAlliance `json:"alliance"`
	Playoff        TeamEventStatusPlayoff  `json:"playoff"`
	AllianceStatus string                  `json:"alliance_status_str"`
	PlayoffStatus  string                  `json:"playoff_status_str"`
	OverallStatus  string                  `json:"overall_status_str"`
	NextMatchKey   string                  `json:"next_match_key"`
	LastMatchKey   string                  `json:"last_match_key"`
}

type TeamEventStatusRank struct {
	NumberOfTeams int         `json:"num_teams"`
	Ranking       Ranking     `json:"ranking"`
	SortOrderInfo []SortOrder `json:"sort_order_info"`
	Status        string      `json:"status"`
}

type TeamEventStatusAlliance struct {
	Name   string                        `json:"name"`
	Number int                           `json:"number"`
	Pick   int                           `json:"pick"`
	Backup TeamEventStatusAllianceBackup `json:"backup"`
}

type TeamEventStatusPlayoff struct {
	Level              string           `json:"level"`
	CurrentLevelRecord WinLossTieRecord `json:"current_level_record"`
	Record             WinLossTieRecord `json:"record"`
	Status             string           `json:"status"`
	PlayoffAverage     int              `json:"playoff_average"`
}

type TeamEventStatusAllianceBackup struct {
	Out string `json:"out"`
	In  string `json:"in"`
}
