package parser

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sunny_5_skiers/model"
)

func ParseEvents(path string) ([]model.Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var events []model.Event
	stringEvents := strings.Split(string(data), "\n")

	for _, stringEvent := range stringEvents {
		if stringEvent == "" {
			continue
		}

		stringEvent = strings.TrimSpace(stringEvent)
		parts := strings.Split(stringEvent, " ")
		if len(parts) < 3 {
			return nil, fmt.Errorf("некорректный формат события: %s", stringEvent)
		}

		eventId, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		competitorId, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		event := model.Event{
			Time:       parts[0],
			Id:         eventId,
			Competitor: competitorId,
		}
		if len(parts) > 3 {
			event.ExtraParams = parts[3]
		}
		events = append(events, event)
	}
	return events, nil
}
