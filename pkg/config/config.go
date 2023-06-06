package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Servers         map[string]string
	BotToken        string
	ChatID          int64
	PagerDutyToken  string
	PagerDutyUserID string
}

func LoadConfig(path string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
