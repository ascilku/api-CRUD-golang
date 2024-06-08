package campaign

type Service interface {
	FindAllSer() ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllSer() ([]Campaign, error) {
	findAllRepo, err := s.repository.FindAllRepo()
	if err != nil {
		return findAllRepo, err
	}
	return findAllRepo, nil
}
