package parser

import (
	"fmt"
	"log"
	"sunny_5_skiers/model"
	"time"
)

func BuildCompetitors(events []model.Event, config model.Config) map[int]*model.Competitor {
	competitors := make(map[int]*model.Competitor)

	for _, event := range events {
		switch event.Id {
		case 1:
			competitors[event.Competitor] = &model.Competitor{
				TotalTime: "[NotStarted]",
			}
		case 2:
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.PlanedStart = "[" + event.ExtraParams + "]"
			}
		case 3:
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.StartLine = event.Time
			}
		case 4:
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.ActualStart = event.Time
				competitor.Started = true
				competitor.TotalTime = TimeDiffStr(event.Time, competitor.PlanedStart)
				lapId := competitor.CurrentLapId
				if lapId < len(competitor.Laps) {
					competitor.Laps[lapId].Start = TimeDiffStr(event.Time, competitor.TotalTime)
				}
			}
		case 5:
			if competitor, exists := competitors[event.Competitor]; exists {
				lapId := competitor.CurrentLapId
				if lapId < len(competitor.Laps) {
					competitor.Laps[lapId].Finish = event.Time
				}
			}

		}
	}

	return competitors
}

func ParseTime(str string) time.Time {
	str = str[1 : len(str)-1]

	t, err := time.Parse("15:04:05.000", str)
	if err != nil {
		log.Fatalf("ошибка парсинга времени: %v", err)
	}
	return t
}

func TimeDiffStr(t1Str, t2Str string) string {
	t1 := ParseTime(t1Str)
	t2 := ParseTime(t2Str)
	d := t2.Sub(t1)
	if d < 0 {
		d = -d
	}

	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	millis := int(d.Milliseconds()) % 1000

	return fmt.Sprintf("[%02d:%02d:%02d.%03d]", hours, minutes, seconds, millis)
}
