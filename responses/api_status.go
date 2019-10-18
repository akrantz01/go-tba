package responses

type ApiStatus struct {
	CurrentSeason  int                 `json:"current_season"`
	MaxSeason      int                 `json:"max_season"`
	IsDatafeedDown bool                `json:"is_datafeed_down"`
	DownEvents     []string            `json:"down_events"`
	IosVersion     ApiStatusAppVersion `json:"ios"`
	AndroidVersion ApiStatusAppVersion `json:"android"`
}

type ApiStatusAppVersion struct {
	MinAppVersion    int `json:"min_app_version"`
	LatestAppVersion int `json:"latest_app_version"`
}
