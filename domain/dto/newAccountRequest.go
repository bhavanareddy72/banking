package dto

import "strings"

type NewAccountRequest struct {
	customerId  string  `json:"customer_id"`
	AccountType string  `json:"Account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) validate() *errs.Apperror {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil

}
