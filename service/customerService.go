package service

type customerService interface {
	GetAllcustomer(string) ([]domain.customer, *errs.AppError)
	Getcustomer(string) (*dto.customerResponse, *errs.AppError)
}
type DefaultcustomerService struct {
	repo domain.customerRepository
}

func (s DefaultcustomerService) GetAllcustomer(status string) ([]domain.customer, *errs.AppError) {
	if status == "active" {
		status = "1"

	} else if status == "inactive" {
		status = "0"

	} else {
		status = ""
	}
	return s.repo.FindAll(status)

}
func (s DefaultcustomerService) Getcustomer(id string) (*dto.customerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	c.ToDto()
	response := c.ToDto()
	return &response, nil
}
func NewcustomerService(repository domain.customerRepository) DefaultcustomerService {
	return DefaultcustomerService{repo: repository}
}
