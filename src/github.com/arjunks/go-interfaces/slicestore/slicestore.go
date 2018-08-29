package slicestore

import (
	"errors"
	"fmt"

	"github.com/arjunks/go-interfaces/domain"
)

type Slicestore struct {
	slice []domain.Customer
}

func (s *Slicestore) Add(c domain.Customer) error {
	fmt.Println(s.slice)
	flag := 0
	for k, v := range s.slice {
		if v.Name == c.Name {
			flag = 1
			break
		}
	}
	if flag == 0 {
		s.slice = append(s.slice, c)
		return nil
	}
	fmt.Println(s.slice)
	return errors.New("Customer Present")
}

func (s *Slicestore) Delete(str string) error {
	if _, ok := m.MyCustomer[str]; ok {
		delete(m.MyCustomer, str)
		return nil
	}

	return errors.New("Not Found")
}

func (s *Slicestore) GetById(str string) (domain.Customer, error) {
	val, ok := m.MyCustomer[str]
	if ok {
		return m.MyCustomer[str], nil
	}
	return val, errors.New("Not Found")
}

func (s *Slicestore) GetAll() ([]domain.Customer, error) {
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
