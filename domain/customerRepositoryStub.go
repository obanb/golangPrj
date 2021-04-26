package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "1001",
			Name:        "Ondra B",
			City:        "Kolin",
			Zipcode:     "29002",
			DateOfBirth: "1989-07-08",
			Status:      "1",
		},
		{
			Id:          "1002",
			Name:        "Ondra C",
			City:        "Prague",
			Zipcode:     "29001",
			DateOfBirth: "1990-07-08",
			Status:      "0",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}
