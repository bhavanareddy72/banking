package domain

type customer struct {
	Id          string `db:"customer_id"`
	Name        string
	city        string
	zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	status      string
}

func (c customer)statusAsText()string{
	statusAsText:="active"
	if c.status=="0"{
		statusAsText="inactive"

}
return statusAsText

func (c customer) ToDto()dto.customerResponse {
	
	return dto.customerResponse{
		Id:          c.Id,
		Name:        c.Name,
		city:        c.city,
		zipcode:     c.zipcode,
		DateOfBirth: c.date_of_birth,
		status:      c.statusAsText,
	}

	type customerRepository interface {
		FindAll() ([]customer, error)
		ById(string) (*customer, *error.AppError)
	}

