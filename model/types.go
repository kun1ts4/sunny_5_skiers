package model

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`
	StartDelta  string `json:"startDelta"`
}

type Event struct {
	Time        string
	Id          int
	Competitor  int
	ExtraParams string
}

type Lap struct {
	Time      float64
	Speed     float64
	IsPenalty bool
	Hits      int
	Shots     int
	Start     string
	Finish    string
}

type Competitor struct {
	TotalTime    string
	PlanedStart  string
	StartLine    string
	ActualStart  string
	Started      bool
	Laps         []Lap
	CurrentLapId int
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
