package db

type User struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint
	FirstName string
	LastName  string
	UserName  string
}

type Library struct {
	Id      uint `gorm:"primaryKey"`
	UserId  uint
	MangaId uint
}
