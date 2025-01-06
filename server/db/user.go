package db

import (
	"gorm.io/gorm"
)

func (s *Store) CreateUser(user *User) error {
	err := s.db.First(&User{}, "id = ?", user.Id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(user).Error
		}
		return err
	}
	
	return nil
}
