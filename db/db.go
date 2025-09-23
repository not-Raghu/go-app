package db

import (
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//what is life

var Db *gorm.DB

func ConnectDb() {

	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("NO DATABASE URL")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: Logger(),
	})

	if err != nil {
		log.Fatal("couldn't connect to the database")
	}

	Db = db
}

// logging options for gorm
func Logger() logger.Interface {
	logLevel := logger.Error

	switch os.Getenv("GORM_LOG_LEVEL") {
	case "silent":
		logLevel = logger.Silent
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	}

	slowThreshold := 500 * time.Millisecond

	if os.Getenv("GORM_SLOWTHRESHOLD") != "" {
		threshold, _ := strconv.Atoi(os.Getenv("GORM_SLOWTHRESHOLD"))
		slowThreshold = time.Duration(threshold) * time.Millisecond
	}

	colorful := true

	if os.Getenv("GORM_LOGGER_COLORFUL") != "" {
		colorful, _ = strconv.ParseBool(os.Getenv("GORM_LOGGER_COLORFUL"))
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowThreshold,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  colorful,
		},
	)

	return newLogger
}
