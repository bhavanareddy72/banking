package domain

type customerRepositorystub struct {
	customers []customer
}

func (s customerRepositorystub) FindAll() ([]customer, error) {
	return s.customers, nil
}
func NewcustomerRepositorystub() customerRepositorystub {
	customers []customer{
		{Id: "10", Name: "bhavana", city: "boston", zipcode: "02125", DateOfBirth: "oct 15 1997", status: "5"},
		{Id: "12", Name: "praneeth", city: "America", zipcode: "02127", DateOfBirth: "may 1 1992", status: "3"},
	}

	return customerRepositorystub{customers:customers}
}
