package db

func (s *Store) GetLibrary(userId uint) ([]string, error) {
	var library []string
	res := s.db.Model(&Library{}).Where("user_id = ?", userId).Pluck("manga_url", &library)
	if res.Error != nil {
		return nil, res.Error
	}

	return library, nil
}
