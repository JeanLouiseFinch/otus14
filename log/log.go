package log

import (
	"go.uber.org/zap"
)

func GetLogger(env string) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)
	switch env {
	case "prod":
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	case "dev":
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	default:
		logger = zap.NewExample()
	}

	return logger, nil
}
