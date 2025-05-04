package main

import (
	"fmt"
	"log"
	"sunny_5_skiers/config"
	"sunny_5_skiers/parser"
)

func main() {
	conf, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("ошибка загрузки конфигурации: %v", err)
	}
	fmt.Println(conf)

	events, err := parser.ParseEvents("events")
	if err != nil {
		log.Fatalf("ошибка парсинга событий: %v", err)
	}
	fmt.Println(events)

	competitors := parser.BuildCompetitors(events, conf)
	for _, competitor := range competitors {
		fmt.Println(competitor)
	}
}
