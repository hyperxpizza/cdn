package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Uploader struct {
		MaxFileSize int64
	}
	Database struct {
		User     string
		Password string
		Port     int
		Name     string
		Host     string
	}
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
