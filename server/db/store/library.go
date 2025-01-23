package store

import (
	"github.com/dxtym/yomu/server/db/models"
)

func (s *Store) AddLibrary(record *models.Library) error {
	if err := s.db.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (s *Store) GetLibrary(userId int64) ([]models.Library, error) {
	var library []models.Library
	if err := s.db.Where("user_id = ?", userId).Find(&library).Error; err != nil {
		return nil, err
	}

	return library, nil
}

func (s *Store) RemoveLibrary(userId int64, manga string) error {
	return s.db.Where("user_id = ? AND manga = ?", userId, manga).Delete(&models.Library{}).Error
}
