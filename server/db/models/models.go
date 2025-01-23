package models

import (
	"gorm.io/gorm"
)

type Library struct {
	gorm.Model
	UserId     int64  `gorm:"index:idx_lib_user_manga,unique;not null"`
	Manga      string `gorm:"index:idx_lib_user_manga,unique;not null"`
	CoverImage string `gorm:"not null"`
}

type History struct {
	gorm.Model
	UserId int64  `gorm:"index:idx_hs_user_manga,unique;not null"`
	Manga  string `gorm:"index:idx_hs_user_manga,unique;not null"`
}

type Progress struct {
	gorm.Model
	UserId  int64  `gorm:"index:idx_prog_user_manga;not null"`
	Manga   string `gorm:"index:idx_prog_user_manga;not null"`
	Chapter string `gorm:"index:idx_prog_user_manga;not null"`
	Page    int64  `gorm:"not null;default:1"`
}
