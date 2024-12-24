package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Store struct {
	db *gorm.DB
}

func NewStore(dsn string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return &Store{db: db}, nil
}
