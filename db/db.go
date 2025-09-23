package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//what is life

type Repository struct {
	Db *gorm.DB
}

var Dbase *Repository

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

	Dbase = &Repository{Db: db}
}

func Logger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	return newLogger
}
