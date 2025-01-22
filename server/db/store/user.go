package store

import (
	"github.com/dxtym/yomu/server/db/models"
	"gorm.io/gorm"
)

func (s *Store) CreateUser(user *models.User) error {
	err := s.db.First(&models.User{}, "id = ?", user.UserId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(user).Error
		}
		return err
	}

	return nil
}
