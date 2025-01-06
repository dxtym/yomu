package db

import "time"

type User struct {
	Id        uint   `gorm:"primary_key"`
	FirstName string `gorm:"not null"`
	LastName  string
	UserName  string
	CreatedAt time.Time `gorm:"default:now()"`
}

type Library struct {
	Id         uint      `gorm:"primary_key"`
	UserId     uint      `gorm:"index:idx_user_manga,unique;not null"`
	Manga      string    `gorm:"index:idx_user_manga,unique;not null"`
	CoverImage string    `gorm:"not null"`
	AddedAt    time.Time `gorm:"default:now()"`
}

type History struct {
	Id     uint      `gorm:"primary_key"`
	UserId uint      `gorm:"index:idx_user_manga,unique;not null"`
	Manga  string    `gorm:"index:idx_user_manga,unique;not null"`
	ReadAt time.Time `gorm:"default:now()"`
}

type Progress struct {
	Id       uint      `gorm:"primary_key"`
	UserId   uint      `gorm:"index:idx_user_manga;not null"`
	Manga    string    `gorm:"index:idx_user_manga;not null"`
	Chapter  uint      `gorm:"not null"`
	Page     uint      `gorm:"not null;default:1"`
	UpdateAt time.Time `gorm:"default:now()"`
}
