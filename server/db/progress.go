package db

import (
	"gorm.io/gorm"
)

func (s *Store) UpdateProgress(progress *Progress) error {
	var p Progress
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
	p.UpdateAt = progress.UpdateAt
	return s.db.Save(&p).Error
}

func (s *Store) GetProgress(userId uint64, manga string, chapter string) (uint64, error) {
	var page uint64
	err := s.db.Model(&Progress{}).Select("page").Where(
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
