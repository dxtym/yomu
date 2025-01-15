package db

import "time"

type User struct {
	Id        uint64 `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string
	UserName  string
	CreatedAt time.Time `gorm:"default:now()"`
}

type Library struct {
	Id         uint64    `gorm:"primary_key"`
	UserId     uint64    `gorm:"index:idx_lib_user_manga,unique;not null"`
	Manga      string    `gorm:"index:idx_lib_user_manga,unique;not null"`
	CoverImage string    `gorm:"not null"`
	AddedAt    time.Time `gorm:"default:now()"`
}

type History struct {
	Id     uint64    `gorm:"primary_key"`
	UserId uint64    `gorm:"index:idx_hs_user_manga,unique;not null"`
	Manga  string    `gorm:"index:idx_hs_user_manga,unique;not null"`
	ReadAt time.Time `gorm:"default:now()"`
}

type Progress struct {
	Id       uint64    `gorm:"primary_key"`
	UserId   uint64    `gorm:"index:idx_prog_user_manga;not null"`
	Manga    string    `gorm:"index:idx_prog_user_manga;not null"`
	Chapter  string    `gorm:"index:idx_prog_user_manga;not null"`
	Page     uint64    `gorm:"not null;default:1"`
	UpdateAt time.Time `gorm:"default:now()"`
}
