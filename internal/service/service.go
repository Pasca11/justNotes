package service

type Service interface {
}

type ServiceImpl struct {
}

func New() (Service, error) {
	return &ServiceImpl{}, nil
}
