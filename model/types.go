package model

type Config struct {
	Laps        int     `json:"laps"`
	LapLen      float64 `json:"lapLen"`
	PenaltyLen  float64 `json:"penaltyLen"`
	FiringLines int     `json:"firingLines"`
	Start       string  `json:"start"`
	StartDelta  string  `json:"startDelta"`
}

type Event struct {
	Time        string
	Id          int
	Competitor  int
	ExtraParams string
}

type Lap struct {
	Time      float64          `json:"time"`
	Speed     float64          `json:"average_speed"`
	Length    float64          `json:"-"`
	IsPenalty bool             `json:"-"`
	Shots     map[int]struct{} `json:"-"`
	Hits      int              `json:"-"`
	Start     string           `json:"-"`
	Finish    string           `json:"-"`
}
type Competitor struct {
	TotalTime     string  `json:"total_time"`
	PlanedStart   string  `json:"-"`
	StartLine     string  `json:"-"`
	ActualStart   string  `json:"-"`
	Started       bool    `json:"-"`
	Laps          []Lap   `json:"laps"`
	CurrentLapId  int     `json:"-"`
	PenaltyTime   string  `json:"penalty_laps_time"`
	PenaltySpeed  float64 `json:"penalty_laps_speed"`
	ShootingStats string  `json:"shooting_stats"`
}

// [NotFinished] 1 [{00:29:03.872, 2.093}, {,}] {00:01:44.296, 0.481} 4/5
//Final report
//The final report should contain the list of all registered competitors sorted by ascending time.
//
//Total time includes the difference between scheduled and actual start time or NotStarted/NotFinished marks
//Time taken to complete each lap
//Average speed for each lap [m/s]
//Time taken to complete penalty laps
//Average speed over penalty laps [m/s]
//Number of hits/number of shots
