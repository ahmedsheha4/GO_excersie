package businesslogic

type Repository interface {
	GetPhonesFromRepository(country string, state string) ([]PhoneNumber, error)
}

type Service interface {
	GetPhones(country string, state string) ([]PhoneNumber, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetPhones(country string, state string) ([]PhoneNumber, error) {
	phones, err := s.r.GetPhonesFromRepository(country, state)
	if err != nil {
		return nil, err
	}
	return phones, nil
}
