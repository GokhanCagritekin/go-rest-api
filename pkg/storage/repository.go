package storage

type Storage struct {
	DB map[string]string
}

func NewStorage() *Storage {
	db := new(Storage)
	return db
}
