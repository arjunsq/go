package mapstore

import (
	"errors"

	"github.com/arjunks/go-interfaces/domain"
)

type MapStore struct {
	MyCustomer map[string]domain.Customer
}

func (m *MapStore) Add(c domain.Customer) error {

	if _, ok := m.MyCustomer[c.Name]; ok {
		return errors.New("Already exists")
	}

	m.MyCustomer[c.Name] = c
	return nil
}

func (m *MapStore) Delete(str string) error {
	if _, ok := m.MyCustomer[str]; ok {
		delete(m.MyCustomer, str)
		return nil
	}

	return errors.New("Not Found")
}

func (m *MapStore) GetById(str string) (domain.Customer, error) {
	val, ok := m.MyCustomer[str]
	if ok {
		return m.MyCustomer[str], nil
	}
	return val, errors.New("Not Found")
}

func (m *MapStore) GetAll() ([]domain.Customer, error) {
	c := m.MyCustomer
	customer := []domain.Customer{}

	for _, val := range c {
		customer = append(customer, val)
	}
	if len(c) == 0 {
		return customer, errors.New("Nothing in slice")
	}
	return customer, nil
}
