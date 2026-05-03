package logger

import (
	"go-cover-parroto/internal/configs"
	"sync"

	"go.uber.org/zap"
)

var (
	sugar *zap.SugaredLogger
	once  sync.Once
)

// S provides the global access point
func S() *zap.SugaredLogger {
	if sugar == nil {
		// This panic is acceptable because it catches developer
		// errors (forgetting to Init) rather than system errors.
		panic("logger accessed before Init()")
	}
	return sugar
}

func Init(cfg configs.LoggerConfig) error {
	var err error
	once.Do(func() {
		var z *zap.Logger

		if cfg.Mode == "development" {
			z, err = zap.NewDevelopment()
		} else {
			z, err = zap.NewProduction()
		}

		if err != nil {
			return
		}

		sugar = z.Sugar()
	})
	return err
}
