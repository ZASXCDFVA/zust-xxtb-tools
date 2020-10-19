package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram struct {
		Token  string `yaml:"token"`
		ChatId string `yaml:"chat-id"`
	} `yaml:"telegram"`

	Users []struct {
		Id       string `yaml:"id"`
		Province string `yaml:"province"`
		City     string `yaml:"city"`
		Country  string `yaml:"country"`
	} `yaml:"users"`
}

func ParseConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
