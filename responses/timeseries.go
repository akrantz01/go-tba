package responses

type Timeseries2018 struct {
	EventKey                 string `json:"event_key"`
	MatchId                  string `json:"match_id"`
	Mode                     string `json:"mode"`
	Play                     int    `json:"play"`
	TimeRemaining            int    `json:"time_remaining"`
	BlueAutoQuest            int    `json:"blue_auto_quest"`
	BlueBoostCount           int    `json:"blue_boost_count"`
	BlueBoostPlayed          int    `json:"blue_boost_played"`
	BlueCurrentPowerup       string `json:"blue_current_powerup"`
	BlueFaceTheBoss          int    `json:"blue_face_the_boss"`
	BlueForceCount           int    `json:"blue_force_count"`
	BlueForcePlayed          int    `json:"blue_force_played"`
	BlueLevitateCount        int    `json:"blue_levitate_count"`
	BlueLevitatePlayed       int    `json:"blue_levitate_played"`
	BluePowerupTimeRemaining int    `json:"blue_powerup_time_remaining"`
	BlueScaleOwned           int    `json:"blue_scale_owned"`
	BlueScore                int    `json:"blue_score"`
	BlueSwitchOwned          int    `json:"blue_switch_owned"`
	RedAutoQuest             int    `json:"red_auto_quest"`
	RedBoostCount            int    `json:"red_boost_count"`
	RedBoostPlayed           int    `json:"red_boost_played"`
	RedCurrentPowerup        string `json:"red_current_powerup"`
	RedFaceTheBoss           int    `json:"red_face_the_boss"`
	RedForceCount            int    `json:"red_force_count"`
	RedForcePlayed           int    `json:"red_force_played"`
	RedLevitateCount         int    `json:"red_levitate_count"`
	RedLevitatePlayed        int    `json:"red_levitate_played"`
	RedPowerupTimeRemaining  int    `json:"red_powerup_time_remaining"`
	RedScaleOwned            int    `json:"red_scale_owned"`
	RedScore                 int    `json:"red_score"`
	RedSwitchOwned           int    `json:"red_switch_owned"`
}
