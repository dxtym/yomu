package store

import (
	"github.com/dxtym/yomu/server/db/models"
	"gorm.io/gorm"
)

func (s *Store) UpdateProgress(progress *models.Progress) error {
	var p models.Progress
	err := s.db.Where(
		"user_id = ? AND manga = ? AND chapter = ?",
		progress.UserId, progress.Manga, progress.Chapter,
	).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Create(progress).Error
		}

		return err
	}

	p.Page = progress.Page
	p.UpdatedAt = progress.UpdatedAt
	return s.db.Save(&p).Error
}

func (s *Store) GetProgress(userId int64, manga string, chapter string) (int64, error) {
	var page int64
	err := s.db.Model(&models.Progress{}).Select("page").Where(
		"user_id = ? AND manga = ? AND chapter = ?",
		userId, manga, chapter,
	).Find(&page).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return page, nil
		}

		return page, err
	}

	return page, nil
}
