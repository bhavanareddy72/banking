package domain

type Account struct {
	AccountId   string
	customerId  string
	openingdate string
	Accounttype string
	Amount      float64
	status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return NewAccountResponse{a.AccountId}

}

type AccountRepository interface {
	save(Account) (*Account, *errs.Apperror)
}
