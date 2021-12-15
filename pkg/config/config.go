package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Uploader struct {
		MaxFileSize int64 `json:"MaxFileSize"`
	} `json:"Upload"`
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		Host     string `json:"host"`
	} `json:"Database"`
}

func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
