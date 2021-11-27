package listing

type Service interface {
	Get(key string) (string, error)
}

type Repository interface {
	Get(key string) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Get(key string) (string, error) {
	return s.r.Get(key)
}
