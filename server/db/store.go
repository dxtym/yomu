package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(dsn string) (*Store, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{}, &Library{})

	return &Store{db: db}, nil
}
