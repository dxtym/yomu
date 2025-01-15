package db

import "gorm.io/gorm"

func (s *Store) AddHistory(record *History) error {
	var h History
	err := s.db.Where("user_id = ? AND manga = ?", record.UserId, record.Manga).First(&h).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(record).Error
		}

		return err
	}

	h.ReadAt = record.ReadAt
	return s.db.Save(&h).Error
}

func (s *Store) GetHistory(userId uint64) ([]History, error) {
	var h []History
	if err := s.db.Order("read_at desc").Where("user_id = ?", userId).Find(&h).Error; err != nil {
		return nil, err
	}

	return h, nil
}

func (s *Store) RemoveHistory(userId uint64, id uint64) error {
	return s.db.Where("user_id = ? AND id = ?", userId, id).Delete(&History{}).Error
}