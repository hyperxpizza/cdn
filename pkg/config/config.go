package config

import (
	"encoding/json"
	"fmt"
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
	FileBrowser struct {
		Rootpath string `json:"rootPath"`
	} `json:"FileBrowser"`
	Grpc struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"Grpc"`
	Rest struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"Rest"`
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

func (c *Config) PrettyPrint() {
	data, err := json.MarshalIndent(c, " ", "")
	if err != nil {
		return
	} else {
		fmt.Println(string(data))
	}
}
