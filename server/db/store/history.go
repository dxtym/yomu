package store

import (
	"time"

	"github.com/dxtym/yomu/server/db/models"
	"gorm.io/gorm"
)

func (s *Store) AddHistory(record *models.History) error {
	var h models.History
	err := s.db.Where("user_id = ? AND manga = ?", record.UserId, record.Manga).First(&h).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(record).Error
		}

		return err
	}

	h.UpdatedAt = time.Now()
	return s.db.Save(&h).Error
}

func (s *Store) GetHistory(userId int64) ([]models.History, error) {
	var h []models.History
	err := s.db.Order("updated_at desc").Where("user_id = ?", userId).Find(&h).Error
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (s *Store) RemoveHistory(id, userId int64) error {
	return s.db.Where("user_id = ? AND id = ?", userId, id).Delete(&models.History{}).Error
}
