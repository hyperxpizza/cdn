package logger

import "github.com/hyperxpizza/cdn/pkg/config"

type Logger struct {
	outputFilePath string
}

func NewLogger(c *config.Config) (*Logger, error) {
	logger := Logger{}

	return &logger, nil
}
