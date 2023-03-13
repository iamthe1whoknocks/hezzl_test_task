package logger

import (
	"log"

	"go.uber.org/zap"
)

// logger struct
type Logger struct {
	L *zap.Logger
}

// set logger
func Set(mode string) *Logger {
	var (
		l   *zap.Logger
		err error
	)

	if mode == "debug" {
		l, err = zap.NewDevelopment(zap.AddStacktrace(zap.DPanicLevel))
		if err != nil {
			log.Fatal("error creating logger : ", err.Error())
		}
		l.Debug("Logger started", zap.String("mode", "debug"))
	} else {
		l, err = zap.NewProduction(zap.AddStacktrace(zap.DPanicLevel))
		if err != nil {
			log.Fatal("error creating logger : ", err.Error())
		}
		l.Info("Logger started", zap.String("mode", "production"))
	}
	return &Logger{
		L: l,
	}
}
