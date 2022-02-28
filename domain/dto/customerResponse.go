package dto

type customerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	city        string `json:"city"`
	zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	status      string `json:"status"`
}
