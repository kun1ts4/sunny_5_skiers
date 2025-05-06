package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sunny_5_skiers/model"
	"time"
)

func BuildCompetitors(events []model.Event, config model.Config) map[int]*model.Competitor {
	competitors := make(map[int]*model.Competitor)
	for _, event := range events {
		switch event.Id {
		case 1: // The competitor registered
			laps := make([]model.Lap, config.Laps)
			for i := range laps {
				laps[i] = model.Lap{
					Shots:  make(map[int]struct{}),
					Length: config.LapLen,
				}
			}
			competitors[event.Competitor] = &model.Competitor{
				TotalTime: "[NotStarted]",
				Laps:      laps,
			}
			log.Printf("%s The competitor(%d) registered", event.Time, event.Competitor)

		case 2: // The start time was set by a draw
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.PlanedStart = "[" + event.ExtraParams + "]"
				log.Printf("%s The start time for the competitor(%d) was set by a draw to %s", event.Time, event.Competitor, competitor.PlanedStart)
			}

		case 3: // The competitor is on the start line
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.StartLine = event.Time
				log.Printf("%s The competitor(%d) is on the start line", event.Time, event.Competitor)
			}

		case 4: // The competitor has started
			if competitor, exists := competitors[event.Competitor]; exists {
				competitor.ActualStart = event.Time
				competitor.Started = true
				competitor.TotalTime = TimeDiffStr(event.Time, competitor.PlanedStart)
				lapId := competitor.CurrentLapId
				if lapId < len(competitor.Laps) {
					if isValidTime(competitor.PlanedStart) {
						competitor.TotalTime = TimeDiffStr(competitor.PlanedStart, event.Time)
					}
					competitor.Laps[lapId].Start = event.Time
				}
				log.Printf("%s The competitor(%d) has started", event.Time, event.Competitor)
			}

		case 5: // The competitor is on the firing range
			log.Printf("%s The competitor(%d) is on the firing range(%s)", event.Time, event.Competitor, event.ExtraParams)

		case 6: // The target has been hit
			if competitor, exists := competitors[event.Competitor]; exists {
				lapId := competitor.CurrentLapId
				target, err := strconv.Atoi(event.ExtraParams)
				if err != nil {
					log.Fatal(err)
				}
				competitor.Laps[lapId].Shots[target] = struct{}{}
				log.Printf("%s The target(%d) has been hit by competitor(%d)", event.Time, target, event.Competitor)
			}

		case 7: // The competitor left the firing range
			if competitor, exists := competitors[event.Competitor]; exists {
				lapId := competitor.CurrentLapId
				competitor.Laps[lapId].Hits = len(competitor.Laps[lapId].Shots)
				log.Printf("%s The competitor(%d) left the firing range", event.Time, event.Competitor)
			}

		case 8: // The competitor entered the penalty laps
			log.Printf("%s The competitor(%d) entered the penalty laps", event.Time, event.Competitor)

		case 9: // The competitor left the penalty laps
			log.Printf("%s The competitor(%d) left the penalty laps", event.Time, event.Competitor)

		case 10: // The competitor ended the main lap
			if competitor, exists := competitors[event.Competitor]; exists {
				lapId := competitor.CurrentLapId
				competitor.Laps[lapId].Finish = event.Time

				if competitor.TotalTime != "[NotStarted]" && competitor.Laps[lapId].Start != "" {
					competitor.Laps[lapId].Time = TimeToSeconds(TimeDiffStr(competitor.Laps[lapId].Start, event.Time))
					competitor.Laps[lapId].Speed = competitor.Laps[lapId].Length / competitor.Laps[lapId].Time
				}
				competitor.CurrentLapId++
				log.Printf("%s The competitor(%d) ended the main lap", event.Time, event.Competitor)
			}

		case 11: // The competitor can`t continue
			if c, ok := competitors[event.Competitor]; ok {
				c.TotalTime = "[NotFinished]"
				log.Printf("%s The competitor(%d) can't continue: %s", event.Time, event.Competitor, event.ExtraParams)
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

func TimeToSeconds(tStr string) float64 {
	parts := strings.Split(tStr, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	secondsAndMilliseconds := parts[2]
	secondsParts := strings.Split(secondsAndMilliseconds, ".")
	seconds, _ := strconv.Atoi(secondsParts[0])
	milliseconds := 0
	if len(secondsParts) > 1 {
		milliseconds, _ = strconv.Atoi(secondsParts[1])
	}

	return float64(hours*3600+minutes*60+seconds) + float64(milliseconds)/1000
}

func isValidTime(str string) bool {
	return len(str) > 2 && str[0] == '[' && str[len(str)-1] == ']'
}
