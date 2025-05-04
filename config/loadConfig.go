package config

import (
	"encoding/json"
	"io"
	"os"
	"sunny_5_skiers/model"
)

func LoadConfig(path string) (model.Config, error) {
	config := model.Config{}

	file, err := os.Open(path)

	if err != nil {
		return config, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
