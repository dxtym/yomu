package db

import (
	"gorm.io/gorm"
)

func (s *Store) CreateUser(user *User) error {
	u := s.db.First(user)
	if u.Error != nil {
		if u.Error == gorm.ErrRecordNotFound {
			return s.db.Create(user).Error
		}
		return u.Error
	}
	return nil
}
