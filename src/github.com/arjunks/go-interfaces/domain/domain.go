package domain

type Customer struct {
	Name  string
	Email string
	Age   int
}

type CustomerStore interface {
	Add(Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
