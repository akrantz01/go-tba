package responses

type ScoringBreakdown2016 struct {
	Red  ScoringBreakdownAlliance2016 `mapstructure:"red"`
	Blue ScoringBreakdownAlliance2016 `mapstructure:"blue"`
}

type ScoringBreakdownAlliance2016 struct {
	AutoPoints             int    `mapstructure:"autoPoints"`
	TeleopPoints           int    `mapstructure:"teleopPoints"`
	BreachPoints           int    `mapstructure:"breachPoints"`
	FoulPoints             int    `mapstructure:"foulPoints"`
	CapturePoints          int    `mapstructure:"capturePoints"`
	AdjustPoints           int    `mapstructure:"adjustPoints"`
	TotalPoints            int    `mapstructure:"totalPoints"`
	Robot1Auto             string `mapstructure:"robot1Auto"`
	Robot2Auto             string `mapstructure:"robot2Auto"`
	Robot3Auto             string `mapstructure:"robot3Auto"`
	AutoReachPoints        int    `mapstructure:"autoReachPoints"`
	AutoCrossingPoints     int    `mapstructure:"autoCrossingPoints"`
	AutoBouldersLow        int    `mapstructure:"autoBouldersLow"`
	AutoBouldersHigh       int    `mapstructure:"autoBouldersHigh"`
	AutoBouldersPoints     int    `mapstructure:"autoBouldersPoints"`
	TeleopCrossingPoints   int    `mapstructure:"teleopCrossingPoints"`
	TeleopBouldersLow      int    `mapstructure:"teleopBouldersLow"`
	TeleopBouldersHigh     int    `mapstructure:"teleopBouldersHigh"`
	TeleopBouldersPoints   int    `mapstructure:"teleopBouldersPoints"`
	TeleopDefensesBreached bool   `mapstructure:"teleopDefensesBreached"`
	TeleopChallengePoints  int    `mapstructure:"teleopChallengePoints"`
	TeleopScalePoints      int    `mapstructure:"teleopScalePoints"`
	TeleopTowerCaptured    int    `mapstructure:"teleopTowerCaptured"`
	TowerFaceA             string `mapstructure:"towerFaceA"`
	TowerFaceB             string `mapstructure:"towerFaceB"`
	TowerFaceC             string `mapstructure:"towerFaceC"`
	TowerEndStrength       int    `mapstructure:"towerEndStrength"`
	TechnicalFoulCount     int    `mapstructure:"techFoulCount"`
	FoulCount              int    `mapstructure:"foulCount"`
	Position2              string `mapstructure:"position2"`
	Position3              string `mapstructure:"position3"`
	Position4              string `mapstructure:"position4"`
	Position5              string `mapstructure:"position5"`
	Position1Crossings     int    `mapstructure:"position1crossings"`
	Position2Crossings     int    `mapstructure:"position2crossings"`
	Position3Crossings     int    `mapstructure:"position3crossings"`
	Position4Crossings     int    `mapstructure:"position4crossings"`
	Position5Crossings     int    `mapstructure:"position5crossings"`
}

type ScoringBreakdown2017 struct {
	Red  ScoringBreakdownAlliance2017 `mapstructure:"red"`
	Blue ScoringBreakdownAlliance2017 `mapstructure:"blue"`
}

type ScoringBreakdownAlliance2017 struct {
	AutoPoints                int    `mapstructure:"autoPoints"`
	TeleopPoints              int    `mapstructure:"teleopPoints"`
	FoulPoints                int    `mapstructure:"foulPoints"`
	AdjustPoints              int    `mapstructure:"adjustPoints"`
	TotalPoints               int    `mapstructure:"totalPoints"`
	Robot1Auto                string `mapstructure:"robot1Auto"`
	Robot2Auto                string `mapstructure:"robot2Auto"`
	Robot3Auto                string `mapstructure:"robot3Auto"`
	Rotor1Auto                bool   `mapstructure:"rotor1Auto"`
	Rotor2Auto                bool   `mapstructure:"rotor2Auto"`
	AutoFuelLow               int    `mapstructure:"autoFuelLow"`
	AutoFuelHigh              int    `mapstructure:"autoFuelHigh"`
	AutoMobilityPoints        int    `mapstructure:"autoMobilityPoints"`
	AutoRotorPoints           int    `mapstructure:"autoRotorPoints"`
	AutoFuelPoints            int    `mapstructure:"autoFuelPoints"`
	TeleopFuelPoints          int    `mapstructure:"teleopFuelPoints"`
	TeleopFuelLow             int    `mapstructure:"teleopFuelLow"`
	TeleopFuelHigh            int    `mapstructure:"teleopFuelHigh"`
	TeleopRotorPoints         int    `mapstructure:"teleopRotorPoints"`
	KpaRankingPointAchieved   bool   `mapstructure:"kPaRankingPointAchieved"`
	TeleopTakeoffPoints       int    `mapstructure:"teleopTakeoffPoints"`
	KpaBonusPoints            int    `mapstructure:"kPaBonusPoints"`
	RotorBonusPoints          int    `mapstructure:"rotorBonusPoints"`
	Rotor1Engaged             bool   `mapstructure:"rotor1Engaged"`
	Rotor2Engaged             bool   `mapstructure:"rotor2Engaged"`
	Rotor3Engaged             bool   `mapstructure:"rotor3Engaged"`
	Rotor4Engaged             bool   `mapstructure:"rotor4Engaged"`
	RotorRankingPointAchieved bool   `mapstructure:"rotorRankingPointAchieved"`
	TechnicalFoulCount        int    `mapstructure:"techFoulCount"`
	FoulCount                 int    `mapstructure:"foulCount"`
	TouchpadNear              string `mapstructure:"touchpadNear"`
	TouchpadMiddle            string `mapstructure:"touchpadMiddle"`
	TouchpadFar               string `mapstructure:"touchpadFar"`
}

type ScoringBreakdown2018 struct {
	Red  ScoringBreakdownAlliance2018 `mapstructure:"red"`
	Blue ScoringBreakdownAlliance2018 `mapstructure:"blue"`
}

type ScoringBreakdownAlliance2018 struct {
	AdjustPoints                 int     `mapstructure:"adjustPoints"`
	AutoOwnershipPoints          int     `mapstructure:"autoOwnershipPoints"`
	AutoPoints                   int     `mapstructure:"autoPoints"`
	AutoQuestRankingPoint        bool    `mapstructure:"autoQuestRankingPoint"`
	AutoRobot1                   string  `mapstructure:"autoRobot1"`
	AutoRobot2                   string  `mapstructure:"autoRobot2"`
	AutoRobot3                   string  `mapstructure:"autoRobot3"`
	AutoRunPoints                int     `mapstructure:"autoRunPoints"`
	AutoScaleOwnershipSeconds    float64 `mapstructure:"autoScaleOwnershipSec"`
	AutoSwitchAtZero             bool    `mapstructure:"autoSwitchAtZero"`
	AutoSwitchOwnershipSeconds   float64 `mapstructure:"autoSwitchOwnershipSec"`
	EndgamePoints                int     `mapstructure:"endgamePoints"`
	EndgameRobot1                string  `mapstructure:"endgameRobot1"`
	EndgameRobot2                string  `mapstructure:"endgameRobot2"`
	EndgameRobot3                string  `mapstructure:"endgameRobot3"`
	FaceTheBossRankingPoint      bool    `mapstructure:"faceTheBossRankingPoint"`
	FoulCount                    int     `mapstructure:"foulCount"`
	FoulPoints                   int     `mapstructure:"foulPoints"`
	RankingPoints                int     `mapstructure:"rp"`
	TechnicalFoulCount           int     `mapstructure:"techFoulCount"`
	TeleopOwnershipPoints        int     `mapstructure:"teleopOwnershipPoints"`
	TeleopPoints                 int     `mapstructure:"teleopPoints"`
	TeleopScaleBoostSeconds      float64 `mapstructure:"teleopScaleBoostSec"`
	TeleopScaleForceSeconds      float64 `mapstructure:"teleopScaleForceSec"`
	TeleopScaleOwnershipSeconds  float64 `mapstructure:"teleopOwnershipSec"`
	TeleopSwitchBoostSeconds     float64 `mapstructure:"teleopSwitchBoostSec"`
	TeleopSwitchForceSeconds     float64 `mapstructure:"teleopSwitchForceSec"`
	TeleopSwitchOwnershipSeconds float64 `mapstructure:"teleopSwitchOwnershipSec"`
	TotalPoints                  int     `mapstructure:"totalPoints"`
	VaultBoostPlayed             int     `mapstructure:"vaultBoostPlayed"`
	VaultBoostTotal              int     `mapstructure:"vaultBoostTotal"`
	VaultForcePlayed             int     `mapstructure:"vaultForcePlayed"`
	VaultForceTotal              int     `mapstructure:"vaultForceTotal"`
	VaultLevitatePlayed          int     `mapstructure:"vaultLevitatePlayed"`
	VaultLevitateTotal           int     `mapstructure:"vaultLevitateTotal"`
	VaultPoints                  int     `mapstructure:"vaultPoints"`
	TbaGameData                  string  `mapstructure:"tba_gameData"`
}

type ScoringBreakdown2019 struct {
	Red  ScoringBreakdownAlliance2019 `mapstructure:"red"`
	Blue ScoringBreakdownAlliance2019 `mapstructure:"blue"`
}

type ScoringBreakdownAlliance2019 struct {
	AdjustPoints               int    `mapstructure:"adjustPoints"`
	AutoPoints                 int    `mapstructure:"autoPoints"`
	Bay1                       string `mapstructure:"bay1"`
	Bay2                       string `mapstructure:"bay2"`
	Bay3                       string `mapstructure:"bay3"`
	Bay4                       string `mapstructure:"bay4"`
	Bay5                       string `mapstructure:"bay5"`
	Bay6                       string `mapstructure:"bay6"`
	Bay7                       string `mapstructure:"bay7"`
	Bay8                       string `mapstructure:"bay8"`
	CargoPoints                int    `mapstructure:"cargoPoints"`
	CompleteRocketRankingPoint bool   `mapstructure:"completeRocketRankingPoint"`
	CompletedRocketFar         bool   `mapstructure:"completedRocketFar"`
	CompletedRocketNear        bool   `mapstructure:"completedRocketNear"`
	EndgameRobot1              string `mapstructure:"endgameRobot1"`
	EndgameRobot2              string `mapstructure:"endgameRobot2"`
	EndgameRobot3              string `mapstructure:"endgameRobot3"`
	FoulCount                  int    `mapstructure:"foulCount"`
	FoulPoints                 int    `mapstructure:"foulPoints"`
	HabClimbPoints             int    `mapstructure:"habClimbPoints"`
	HabDockingRankingPoint     bool   `mapstructure:"habDockingRankingPoint"`
	HabLineRobot1              string `mapstructure:"habLineRobot1"`
	HabLineRobot2              string `mapstructure:"habLineRobot2"`
	HabLineRobot3              string `mapstructure:"habLineRobot3"`
	HatchPanelPoints           int    `mapstructure:"hatchPanelPoints"`
	LowLeftRocketFar           string `mapstructure:"lowLeftRocketFar"`
	LowLeftRocketNear          string `mapstructure:"lowLeftRocketNear"`
	LowRightRocketFar          string `mapstructure:"lowRightRocketFar"`
	LowRightRocketNear         string `mapstructure:"lowRightRocketNear"`
	MidLeftRocketFar           string `mapstructure:"midLeftRocketFar"`
	MidLeftRocketNear          string `mapstructure:"midLeftRocketNear"`
	MidRightRockerFar          string `mapstructure:"midRightRocketFar"`
	MidRightRocketNear         string `mapstructure:"midRightRocketNear"`
	PreMatchBay1               string `mapstructure:"preMatchBay1"`
	PreMatchBay2               string `mapstructure:"preMatchBay2"`
	PreMatchBay3               string `mapstructure:"preMatchBay3"`
	PreMatchBay6               string `mapstructure:"preMatchBay6"`
	PreMatchBay7               string `mapstructure:"preMatchBay7"`
	PreMatchBay8               string `mapstructure:"preMatchBay8"`
	PreMatchRobot1             string `mapstructure:"preMatchLevelRobot1"`
	PreMatchRobot2             string `mapstructure:"preMatchLevelRobot2"`
	PreMatchRobot3             string `mapstructure:"preMatchLevelRobot3"`
	RankingPoints              int    `mapstructure:"rp"`
	SandStormBonusPoints       int    `mapstructure:"sandStormBonusPoints"`
	TechnicalFoulCount         int    `mapstructure:"techFoulCount"`
	TeleopPoints               int    `mapstructure:"teleopPoints"`
	TopLeftRocketFar           string `mapstructure:"topLeftRocketFar"`
	TopLeftRocketNear          string `mapstructure:"topLeftRocketNear"`
	TopRightRocketFar          string `mapstructure:"topRightRocketFar"`
	TopRightRocketNear         string `mapstructure:"topRightRocketNear"`
	TotalPoints                int    `mapstructure:"totalPoints"`
}

type ScoringBreakdown2020 struct {
	Red ScoringBreakdown2020Alliance `mapstructure:"red"`
	Blue ScoringBreakdown2020Alliance `mapstructure:"blue"`
}

type ScoringBreakdown2020Alliance struct {
	InitLineRobot1 string `mapstructure:"initLineRobot1"`
	EndgameRobot1 string `mapstructure:"endgameRobot1"`
	InitLineRobot2 string `mapstructure:"initLineRobot2"`
	EndgameRobot2 string `mapstructure:"endgameRobot2"`
	InitLineRobot3 string `mapstructure:"initLineRobot3"`
	EndgameRobot3 string `mapstructure:"endgameRobot3"`
	AutoCellsBottom int `mapstructure:"autoCellsBottom"`
	AutoCellsOuter int `mapstructure:"autoCellsOuter"`
	AutoCellsInner int `mapstructure:"autoCellsInner"`
	TeleopCellsBottom int `mapstructure:"teleopCellsBottom"`
	TeleopCellsOuter int `mapstructure:"teleopCellsOuter"`
	TeleopCellsInner int `mapstructure:"teleopCellsInner"`
	Stage1Activated bool `mapstructure:"stage1Activated"`
	Stage2Activated bool `mapstructure:"stage2Activated"`
	Stage3Activated bool `mapstructure:"stage3Activated"`
	Stage3TargetColor string `mapstructure:"stage3TargetColor"`
	EndgameRungIsLevel string `mapstructure:"endgameRungIsLevel"`
	AutoInitLinePoints int `mapstructure:"autoInitLinePoints"`
	AutoCellPoints int `mapstructure:"autoCellPoints"`
	AutoPoints int `mapstructure:"autoPoints"`
	TeleopCellPoints int `mapstructure:"teleopCellPoints"`
	ControlPanelPoints int `mapstructure:"controlPanelPoints"`
	EndgamePoints int `mapstructure:"endgamePoints"`
	TeleopPoints int `mapstructure:"teleopPoints"`
	ShieldOperationalRankingPoint bool `mapstructure:"shieldOperationalRankingPoint"`
	ShieldEnergizedRankingPoint bool `mapstructure:"shieldEnergizedRankingPoint"`
	FoulCount int `mapstructure:"foulCount"`
	TechFoulCount int `mapstructure:"techFoulCount"`
	AdjustPoints int `mapstructure:"adjustPoints"`
	FoulPoints int `mapstructure:"foulPoints"`
	RankingPoints int `mapstructure:"rp"`
	TotalPoints int `mapstructure:"totlaPoints"`
}
