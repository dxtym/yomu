package db

import "log"

func (s *Store) AddLibrary(record *Library) error {
	if err := s.db.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (s *Store) GetLibrary(userId uint) ([]Library, error) {
	var library []Library
	if err := s.db.Where("user_id = ?", userId).Find(&library).Error; err != nil {
		return nil, err
	}

	log.Printf("%v", library)
	return library, nil
}

func (s *Store) RemoveLibrary(userId uint, manga string) error {
	err := s.db.Where("user_id = ? AND manga = ?", userId, manga).Delete(&Library{}).Error
	if err != nil {
		return err
	}

	return nil
}
