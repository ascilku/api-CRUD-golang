package campaign

type Service interface {
	FindAllSer() ([]Campaign, error)
	FindAllActiveImageAllSer() ([]Campaign, error)
	FindAllUserByIDSer(userID int) ([]Campaign, error)
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

func (s *service) FindAllActiveImageAllSer() ([]Campaign, error) {
	findAllRepo, err := s.repository.FindAllActiveImageRepo()
	if err != nil {
		return findAllRepo, err
	}
	return findAllRepo, nil
}

func (s *service) FindAllUserByIDSer(userID int) ([]Campaign, error) {
	if userID == 0 {
		findAllUserByID, err := s.repository.FindAllUserByID(userID)
		if err != nil {
			return findAllUserByID, err
		}
		return findAllUserByID, nil
	}
	findAllRepo, err := s.repository.FindAllRepo()
	if err != nil {
		return findAllRepo, err
	}
	return findAllRepo, nil
}
