package service

import "time"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.Apperror)
}
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.Apperror) {
	err := req.validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		customerId:  req.customerId,
		openingDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		status:      "1",
	}
	NewAccount, err := s.repo.save(a)
	if err != nil {
		return nil, err
	}
	response := NewAccount.ToNewAccountResponseDto()
	return &response, nil
}
func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
