package db

type User struct {
	Id        int64 `gorm:"primary_key"`
	FirstName string
}

type Manga struct {
	Id       int64  `gorm:"primary_key"`
	Title    string `gorm:"unique"`
	MangaUrl string `gorm:"unique"`
}

type Library struct {
	Id        int64 `gorm:"primary_key"`
	UserId    int64
	MangaId   int64
	ChapterId int64
	PageId    int64
}
