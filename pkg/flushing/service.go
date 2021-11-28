package flushing

type Service interface {
	DeleteAll() error
}

type Repository interface {
	DeleteAll() error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) DeleteAll() error {
	return s.r.DeleteAll()
}
