package responses

type ScoringBreakdown interface {
	Year() int
}

type ScoringBreakdown2016 struct {
	Red  ScoringBreakdownAlliance2016 `json:"red"`
	Blue ScoringBreakdownAlliance2016 `json:"blue"`
}

func (sb ScoringBreakdown2016) Year() int {
	return 2016
}

type ScoringBreakdownAlliance2016 struct {
	AutoPoints             int    `json:"autoPoints"`
	TeleopPoints           int    `json:"teleopPoints"`
	BreachPoints           int    `json:"breachPoints"`
	FoulPoints             int    `json:"foulPoints"`
	CapturePoints          int    `json:"capturePoints"`
	AdjustPoints           int    `json:"adjustPoints"`
	TotalPoints            int    `json:"totalPoints"`
	Robot1Auto             string `json:"robot1Auto"`
	Robot2Auto             string `json:"robot2Auto"`
	Robot3Auto             string `json:"robot3Auto"`
	AutoReachPoints        int    `json:"autoReachPoints"`
	AutoCrossingPoints     int    `json:"autoCrossingPoints"`
	AutoBouldersLow        int    `json:"autoBouldersLow"`
	AutoBouldersHigh       int    `json:"autoBouldersHigh"`
	AutoBouldersPoints     int    `json:"autoBouldersPoints"`
	TeleopCrossingPoints   int    `json:"teleopCrossingPoints"`
	TeleopBouldersLow      int    `json:"teleopBouldersLow"`
	TeleopBouldersHigh     int    `json:"teleopBouldersHigh"`
	TeleopBouldersPoints   int    `json:"teleopBouldersPoints"`
	TeleopDefensesBreached bool   `json:"teleopDefensesBreached"`
	TeleopChallengePoints  int    `json:"teleopChallengePoints"`
	TeleopScalePoints      int    `json:"teleopScalePoints"`
	TeleopTowerCaptured    int    `json:"teleopTowerCaptured"`
	TowerFaceA             string `json:"towerFaceA"`
	TowerFaceB             string `json:"towerFaceB"`
	TowerFaceC             string `json:"towerFaceC"`
	TowerEndStrength       int    `json:"towerEndStrength"`
	TechnicalFoulCount     int    `json:"techFoulCount"`
	FoulCount              int    `json:"foulCount"`
	Position2              string `json:"position2"`
	Position3              string `json:"position3"`
	Position4              string `json:"position4"`
	Position5              string `json:"position5"`
	Position1Crossings     int    `json:"position1crossings"`
	Position2Crossings     int    `json:"position2crossings"`
	Position3Crossings     int    `json:"position3crossings"`
	Position4Crossings     int    `json:"position4crossings"`
	Position5Crossings     int    `json:"position5crossings"`
}

type ScoringBreakdown2017 struct {
	Red  ScoringBreakdownAlliance2017 `json:"red"`
	Blue ScoringBreakdownAlliance2017 `json:"blue"`
}

func (sb ScoringBreakdown2017) Year() int {
	return 2017
}

type ScoringBreakdownAlliance2017 struct {
	AutoPoints                int    `json:"autoPoints"`
	TeleopPoints              int    `json:"teleopPoints"`
	FoulPoints                int    `json:"foulPoints"`
	AdjustPoints              int    `json:"adjustPoints"`
	TotalPoints               int    `json:"totalPoints"`
	Robot1Auto                string `json:"robot1Auto"`
	Robot2Auto                string `json:"robot2Auto"`
	Robot3Auto                string `json:"robot3Auto"`
	Rotor1Auto                bool   `json:"rotor1Auto"`
	Rotor2Auto                bool   `json:"rotor2Auto"`
	AutoFuelLow               int    `json:"autoFuelLow"`
	AutoFuelHigh              int    `json:"autoFuelHigh"`
	AutoMobilityPoints        int    `json:"autoMobilityPoints"`
	AutoRotorPoints           int    `json:"autoRotorPoints"`
	AutoFuelPoints            int    `json:"autoFuelPoints"`
	TeleopFuelPoints          int    `json:"teleopFuelPoints"`
	TeleopFuelLow             int    `json:"teleopFuelLow"`
	TeleopFuelHigh            int    `json:"teleopFuelHigh"`
	TeleopRotorPoints         int    `json:"teleopRotorPoints"`
	KpaRankingPointAchieved   bool   `json:"kPaRankingPointAchieved"`
	TeleopTakeoffPoints       int    `json:"teleopTakeoffPoints"`
	KpaBonusPoints            int    `json:"kPaBonusPoints"`
	RotorBonusPoints          int    `json:"rotorBonusPoints"`
	Rotor1Engaged             bool   `json:"rotor1Engaged"`
	Rotor2Engaged             bool   `json:"rotor2Engaged"`
	Rotor3Engaged             bool   `json:"rotor3Engaged"`
	Rotor4Engaged             bool   `json:"rotor4Engaged"`
	RotorRankingPointAchieved bool   `json:"rotorRankingPointAchieved"`
	TechnicalFoulCount        int    `json:"techFoulCount"`
	FoulCount                 int    `json:"foulCount"`
	TouchpadNear              string `json:"touchpadNear"`
	TouchpadMiddle            string `json:"touchpadMiddle"`
	TouchpadFar               string `json:"touchpadFar"`
}

type ScoringBreakdown2018 struct {
	Red  ScoringBreakdownAlliance2018 `json:"red"`
	Blue ScoringBreakdownAlliance2018 `json:"blue"`
}

func (sb ScoringBreakdown2018) Year() int {
	return 2018
}

type ScoringBreakdownAlliance2018 struct {
	AdjustPoints                 int     `json:"adjustPoints"`
	AutoOwnershipPoints          int     `json:"autoOwnershipPoints"`
	AutoPoints                   int     `json:"autoPoints"`
	AutoQuestRankingPoint        bool    `json:"autoQuestRankingPoint"`
	AutoRobot1                   string  `json:"autoRobot1"`
	AutoRobot2                   string  `json:"autoRobot2"`
	AutoRobot3                   string  `json:"autoRobot3"`
	AutoRunPoints                int     `json:"autoRunPoints"`
	AutoScaleOwnershipSeconds    float64 `json:"autoScaleOwnershipSec"`
	AutoSwitchAtZero             bool    `json:"autoSwitchAtZero"`
	AutoSwitchOwnershipSeconds   float64 `json:"autoSwitchOwnershipSec"`
	EndgamePoints                int     `json:"endgamePoints"`
	EndgameRobot1                string  `json:"endgameRobot1"`
	EndgameRobot2                string  `json:"endgameRobot2"`
	EndgameRobot3                string  `json:"endgameRobot3"`
	FaceTheBossRankingPoint      bool    `json:"faceTheBossRankingPoint"`
	FoulCount                    int     `json:"foulCount"`
	FoulPoints                   int     `json:"foulPoints"`
	RankingPoints                int     `json:"rp"`
	TechnicalFoulCount           int     `json:"techFoulCount"`
	TeleopOwnershipPoints        int     `json:"teleopOwnershipPoints"`
	TeleopPoints                 int     `json:"teleopPoints"`
	TeleopScaleBoostSeconds      float64 `json:"teleopScaleBoostSec"`
	TeleopScaleForceSeconds      float64 `json:"teleopScaleForceSec"`
	TeleopScaleOwnershipSeconds  float64 `json:"teleopOwnershipSec"`
	TeleopSwitchBoostSeconds     float64 `json:"teleopSwitchBoostSec"`
	TeleopSwitchForceSeconds     float64 `json:"teleopSwitchForceSec"`
	TeleopSwitchOwnershipSeconds float64 `json:"teleopSwitchOwnershipSec"`
	TotalPoints                  int     `json:"totalPoints"`
	VaultBoostPlayed             int     `json:"vaultBoostPlayed"`
	VaultBoostTotal              int     `json:"vaultBoostTotal"`
	VaultForcePlayed             int     `json:"vaultForcePlayed"`
	VaultForceTotal              int     `json:"vaultForceTotal"`
	VaultLevitatePlayed          int     `json:"vaultLevitatePlayed"`
	VaultLevitateTotal           int     `json:"vaultLevitateTotal"`
	VaultPoints                  int     `json:"vaultPoints"`
	TbaGameData                  string  `json:"tba_gameData"`
}

type ScoringBreakdown2019 struct {
	Red  ScoringBreakdownAlliance2019 `json:"red"`
	Blue ScoringBreakdownAlliance2019 `json:"blue"`
}

func (sb ScoringBreakdown2019) Year() int {
	return 2019
}

type ScoringBreakdownAlliance2019 struct {
	AdjustPoints               int    `json:"adjustPoints"`
	AutoPoints                 int    `json:"autoPoints"`
	Bay1                       string `json:"bay1"`
	Bay2                       string `json:"bay2"`
	Bay3                       string `json:"bay3"`
	Bay4                       string `json:"bay4"`
	Bay5                       string `json:"bay5"`
	Bay6                       string `json:"bay6"`
	Bay7                       string `json:"bay7"`
	Bay8                       string `json:"bay8"`
	CargoPoints                int    `json:"cargoPoints"`
	CompleteRocketRankingPoint bool   `json:"completeRocketRankingPoint"`
	CompletedRocketFar         bool   `json:"completedRocketFar"`
	CompletedRocketNear        bool   `json:"completedRocketNear"`
	EndgameRobot1              string `json:"endgameRobot1"`
	EndgameRobot2              string `json:"endgameRobot2"`
	EndgameRobot3              string `json:"endgameRobot3"`
	FoulCount                  int    `json:"foulCount"`
	FoulPoints                 int    `json:"foulPoints"`
	HabClimbPoints             int    `json:"habClimbPoints"`
	HabDockingRankingPoint     bool   `json:"habDockingRankingPoint"`
	HabLineRobot1              string `json:"habLineRobot1"`
	HabLineRobot2              string `json:"habLineRobot2"`
	HabLineRobot3              string `json:"habLineRobot3"`
	HatchPanelPoints           int    `json:"hatchPanelPoints"`
	LowLeftRocketFar           string `json:"lowLeftRocketFar"`
	LowLeftRocketNear          string `json:"lowLeftRocketNear"`
	LowRightRocketFar          string `json:"lowRightRocketFar"`
	LowRightRocketNear         string `json:"lowRightRocketNear"`
	MidLeftRocketFar           string `json:"midLeftRocketFar"`
	MidLeftRocketNear          string `json:"midLeftRocketNear"`
	MidRightRockerFar          string `json:"midRightRocketFar"`
	MidRightRocketNear         string `json:"midRightRocketNear"`
	PreMatchBay1               string `json:"preMatchBay1"`
	PreMatchBay2               string `json:"preMatchBay2"`
	PreMatchBay3               string `json:"preMatchBay3"`
	PreMatchBay6               string `json:"preMatchBay6"`
	PreMatchBay7               string `json:"preMatchBay7"`
	PreMatchBay8               string `json:"preMatchBay8"`
	PreMatchRobot1             string `json:"preMatchLevelRobot1"`
	PreMatchRobot2             string `json:"preMatchLevelRobot2"`
	PreMatchRobot3             string `json:"preMatchLevelRobot3"`
	RankingPoints              int    `json:"rp"`
	SandStormBonusPoints       int    `json:"sandStormBonusPoints"`
	TechnicalFoulCount         int    `json:"techFoulCount"`
	TeleopPoints               int    `json:"teleopPoints"`
	TopLeftRocketFar           string `json:"topLeftRocketFar"`
	TopLeftRocketNear          string `json:"topLeftRocketNear"`
	TopRightRocketFar          string `json:"topRightRocketFar"`
	TopRightRocketNear         string `json:"topRightRocketNear"`
	TotalPoints                int    `json:"totalPoints"`
}
