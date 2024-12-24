package db

type User struct {
	Id uint `gorm:"primaryKey"`
	UserId uint 
	FirstName string
	LastName string
	UserName string
}
