package utils

import (
	"github.com/not-raghu/go-app/helpers"
	"github.com/rs/zerolog"
)

var (
	RLog = helpers.GetLogger().ReqLog
	ELog = helpers.GetLogger().ErrLog
)

func GetLog() (zerolog.Logger, zerolog.Logger) {
	return RLog, ELog
}
