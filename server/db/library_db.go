package db

func (s *Store) GetLibrary(userId uint) ([]uint, error) {
	var library []uint
	res := s.db.Model(&Library{}).Where("user_id = ?", userId).Pluck("manga_id", &library)
	if res.Error != nil {
		return nil, res.Error
	}

	return library, nil
}
