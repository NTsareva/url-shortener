package storage

type Storage struct {
	Storage map[string]string
}

func (s *Storage) New() {
	s.Storage = make(map[string]string)
}
