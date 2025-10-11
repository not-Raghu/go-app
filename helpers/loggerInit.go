package helpers

import (
	// "github.com/rs/zerolog"

	"log"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	ReqLog zerolog.Logger
	ErrLog zerolog.Logger
}

var Log *Logger

func LoggerInit() {
	var reqWriter = os.Stderr
	var errWriter = os.Stderr

	f1, e1 := os.OpenFile("req.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e1 != nil {
		log.Println("failed to open error.log, fallback to os.stderr")
	} else {
		reqWriter = f1
	}
	reqLog := zerolog.New(reqWriter).With().Timestamp().Logger()

	f2, e2 := os.OpenFile(
		"error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if e2 != nil {
		log.Println("failed to open error.log, fallback to os.stderr.")
	} else {
		errWriter = f2
	}
	errLog := zerolog.New(errWriter).With().Timestamp().Logger()
	Log = &Logger{
		ReqLog: reqLog,
		ErrLog: errLog,
	}
}

func GetLogger() *Logger {
	if Log == nil {
		LoggerInit()
	}
	return Log
}
