package store

import (
	"log"
	"os"
	"time"

	"github.com/dxtym/yomu/server/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Store struct {
	db *gorm.DB
}

func NewStore(dsn string) (*Store, error) {
	logging := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logging,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Library{}, &models.History{}, &models.Progress{})
	return &Store{db: db}, nil
}
