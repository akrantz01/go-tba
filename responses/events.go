package responses

type EventSimple struct {
	Key           string       `json:"key"`
	Name          string       `json:"name"`
	EventCode     string       `json:"event_code"`
	EventType     string       `json:"event_type"`
	DistrictList  DistrictList `json:"district_list"`
	City          string       `json:"city"`
	StateProvince string       `json:"state_prov"`
	Country       string       `json:"country"`
	StartDate     string       `json:"start_date"`
	EndDate       string       `json:"end_date"`
	Year          int          `json:"year"`
}

type Event struct {
	Key               string       `json:"key"`
	Name              string       `json:"name"`
	EventCode         string       `json:"event_code"`
	EventType         int          `json:"event_type"`
	DistrictList      DistrictList `json:"district_list"`
	City              string       `json:"city"`
	StateProvince     string       `json:"state_prov"`
	Country           string       `json:"country"`
	StartDate         string       `json:"start_date"`
	EndDate           string       `json:"end_date"`
	Year              int          `json:"year"`
	ShortName         string       `json:"short_name"`
	EventTypeString   string       `json:"event_type_string"`
	Week              int          `json:"week"`
	Address           string       `json:"address"`
	PostalCode        string       `json:"postal_code"`
	GoogleMapsPlaceId string       `json:"gmaps_place_id"`
	GoogleMapsUrl     string       `json:"gmaps_url"`
	Latitude          float64      `json:"lat"`
	Longitude         float64      `json:"long"`
	LocationName      string       `json:"location_name"`
	Timezone          string       `json:"timezone"`
	Website           string       `json:"website"`
	FirstEventId      string       `json:"first_event_id"`
	FirstEventCode    string       `json:"first_event_code"`
	DivisionKeys      []string     `json:"division_keys"`
	Webcasts          []Webcast    `json:"webcasts"`
	ParentEventKey    string       `json:"parent_event_key"`
	PlayoffType       int          `json:"playoff_type"`
	PlayoffTypeString string       `json:"playoff_type_string"`
}

type EventRanking struct {
	Rankings       []Ranking   `json:"rankings"`
	ExtraStatsInfo []SortOrder `json:"extra_stats_info"`
	SortOrderInfo  []SortOrder `json:"sort_order_info"`
}

type EventOPRs struct {
	OffensivePointRankings                 map[string]float64 `json:"oprs"`
	DefensivePointRankings                 map[string]float64 `json:"dprs"`
	CalculationContributionToWinningMargin map[string]float64 `json:"ccwms"`
}

type EventAlliance struct {
	Name   string `json:"name"`
	Backup struct {
		In  string `json:"in"`
		Out string `json:"out"`
	} `json:"backup"`
	Declines []string `json:"declines"`
	Picks []string `json:"picks"`
	Status   struct {
		PlayoffAverage     float64          `json:"playoff_average"`
		Level              string           `json:"level"`
		Record             WinLossTieRecord `json:"record"`
		CurrentLevelRecord WinLossTieRecord `json:"current_level_record"`
		Status             string           `json:"status"`
	} `json:"status"`
}

type EventPredictions map[string]interface{}

type EventInsights2019 struct {
	Qualifications map[string]interface{} `json:"qual"`
	Playoffs       map[string]interface{} `json:"playoff"`
}

type EventInsights2018 struct {
	AutoQuestAchieved                             []float64 `json:"auto_quest_achieved"`
	AverageBoostPlayed                            float64   `json:"average_boost_played"`
	AverageEndgamePoints                          float64   `json:"average_endgame_points"`
	AverageForcePlayed                            float64   `json:"average_force_played"`
	AverageFoulScore                              float64   `json:"average_foul_score"`
	AveragePointsAuto                             float64   `json:"average_points_auto"`
	AveragePointsTeleop                           float64   `json:"average_points_teleop"`
	AverageRunPointsAuto                          float64   `json:"average_run_points_auto"`
	AverageScaleOwnershipPoints                   float64   `json:"average_scale_ownership_points"`
	AverageScaleOwnershipPointsAuto               float64   `json:"average_scale_ownership_points_auto"`
	AverageScaleOwnershipPointsTeleop             float64   `json:"average_scale_ownership_points_teleop"`
	AverageScore                                  float64   `json:"average_score"`
	AverageSwitchOwnershipPoints                  float64   `json:"average_switch_ownership_points"`
	AverageSwitchOwnershipPointsAuto              float64   `json:"average_switch_ownership_points_auto"`
	AverageSwitchOwnershipPointsTeleop            float64   `json:"average_switch_ownership_points_teleop"`
	AverageVaultPoints                            float64   `json:"average_vault_points"`
	AverageWinMargin                              float64   `json:"average_win_margin"`
	AverageWinScore                               float64   `json:"average_win_score"`
	BoostPlayedCount                              []float64 `json:"boost_played_count"`
	ClimbCounts                                   []float64 `json:"climb_counts"`
	FaceTheBossAchieved                           []float64 `json:"face_the_boss_achieved"`
	ForcePlayedCounts                             []float64 `json:"force_played_counts"`
	HighScore                                     []string  `json:"high_score"`
	LevitatePlayedCounts                          []float64 `json:"levitate_played_counts"`
	RunCountsAuto                                 []float64 `json:"run_counts_auto"`
	ScaleNeutralPercentage                        float64   `json:"scale_neutral_percentage"`
	ScaleNeutralPercentageAuto                    float64   `json:"scale_neutral_percentage_auto"`
	ScaleNeutralPercentageTeleop                  float64   `json:"scale_neutral_percentage_teleop"`
	SwitchOwnedCountsAuto                         []float64 `json:"switch_owned_counts_auto"`
	UnicornMatches                                []float64 `json:"unicorn_matches"`
	WinningOpponentSwitchDenialPercentageTabletop float64   `json:"winning_opp_switch_denial_percentage_teleop"`
	WinningOwnSwitchOwnershipPercentage           float64   `json:"winning_own_switch_ownership_percentage"`
	WinningOwnSwitchOwnershipPercentageAuto       float64   `json:"winning_own_switch_ownership_percentage_auto"`
	WinningOwnSwitchOwnershipPercentageTeleop     float64   `json:"winning_own_switch_ownership_percentage_teleop"`
	WinningScaleOwnershipPercentage               float64   `json:"winning_scale_ownership_percentage"`
	WinningScaleOwnershipPercentageAuto           float64   `json:"winning_scale_ownership_percentage_auto"`
	WinningScaleOwnershipPercentageTeleop         float64   `json:"winning_scale_ownership_percentage_teleop"`
}

type EventInsights2017 struct {
	AverageFoulScore           float64   `json:"average_foul_score"`
	AverageFuelScore           float64   `json:"average_fuel_score"`
	AverageFuelPointsAuto      float64   `json:"average_fuel_points_auto"`
	AverageFuelPointsTeleop    float64   `json:"average_fuel_points_teleop"`
	AverageHighGoals           float64   `json:"average_high_goals"`
	AverageHighGoalsAuto       float64   `json:"average_high_goals_auto"`
	AverageHighGoalsTeleop     float64   `json:"average_high_goals_teleop"`
	AverageLowGoals            float64   `json:"average_low_goals"`
	AverageLowGoalsAuto        float64   `json:"average_low_goals_auto"`
	AverageLowGoalsTeleop      float64   `json:"average_low_goals_teleop"`
	AverageMobilityPointsAuto  float64   `json:"average_mobility_points_auto"`
	AveragePointsAuto          float64   `json:"average_points_auto"`
	AveragePointsTeleop        float64   `json:"average_points_teleop"`
	AverageRotorPoints         float64   `json:"average_rotor_points"`
	AverageRotorPointsAuto     float64   `json:"average_rotor_points_auto"`
	AverageRotorPointsTeleop   float64   `json:"average_rotor_points_teleop"`
	AverageScore               float64   `json:"average_score"`
	AverageTakeoffPointsTeleop float64   `json:"average_takeoff_points_teleop"`
	AverageWinMargin           float64   `json:"average_win_margin"`
	AverageWinScore            float64   `json:"average_win_score"`
	HighKpa                    []string  `json:"high_kpa"`
	HighScore                  []string  `json:"high_score"`
	KpaAchieved                []float64 `json:"kpa_achieved"`
	MobilityCounts             []float64 `json:"mobility_counts"`
	Rotor1Engaged              []float64 `json:"rotor_1_engaged"`
	Rotor1EngagedAuto          []float64 `json:"rotor_1_engaged_auto"`
	Rotor2Engaged              []float64 `json:"rotor_2_engaged"`
	Rotor2EngagedAuto          []float64 `json:"rotor_2_engaged_auto"`
	Rotor3Engaged              []float64 `json:"rotor_3_engaged"`
	Rotor4Engaged              []float64 `json:"rotor_4_engaged"`
	TakeoffCounts              []float64 `json:"takeoff_counts"`
	UnicornMatches             []float64 `json:"unicorn_matches"`
}

type EventInsights2016 struct {
	LowBar               []float64 `json:"LowBar"`
	AChevalDeFrise       []float64 `json:"A_ChevalDeFrise"`
	APortcullis          []float64 `json:"A_Portcullis"`
	BRamparts            []float64 `json:"B_Ramparts"`
	BMoat                []float64 `json:"B_Moat"`
	CSallyPort           []float64 `json:"C_Sally_Port"`
	CDrawbridge          []float64 `json:"C_Drawbridge"`
	DRoughTerrain        []float64 `json:"D_Rough_Terrain"`
	DRockWall            []float64 `json:"D_Rock_Wall"`
	AverageHighGoals     float64   `json:"average_high_goals"`
	AverageLowGoals      float64   `json:"average_low_goals"`
	Breaches             []float64 `json:"breaches"`
	Scales               []float64 `json:"scales"`
	Challenges           []float64 `json:"challenges"`
	Captures             []float64 `json:"captures"`
	AverageWinScore      float64   `json:"average_win_score"`
	AverageWinMargin     float64   `json:"average_win_margin"`
	AverageScore         float64   `json:"average_score"`
	AverageAutoScore     float64   `json:"average_auto_score"`
	AverageCrossingScore float64   `json:"average_crossing_score"`
	AverageBoulderScore  float64   `json:"average_boulder_score"`
	AverageTowerScore    float64   `json:"average_tower_score"`
	AverageFoulScore     float64   `json:"average_foul_score"`
	HighScore            []string  `json:"high_score"`
}
