package main

import (
	"fmt"

	"github.com/arjunks/go-interfaces/domain"
	"github.com/arjunks/go-interfaces/mapstore"
)

type customerController struct {
	store domain.CustomerStore
}

func (c *customerController) Add(customer domain.Customer) error {
	return c.store.Add(customer)
}

func (c *customerController) Delete(str string) error {
	return c.store.Delete(str)
}

func (c *customerController) GetById(str string) (domain.Customer, error) {
	return c.store.GetById(str)
}

func (c *customerController) GetAll() ([]domain.Customer, error) {
	return c.store.GetAll()
}
func main() {
	m := &mapstore.MapStore{make(map[string]domain.Customer)}
	c := customerController{m}
	c.Add(domain.Customer{"shihad", "adsa@aed", 23})
	c.Add(domain.Customer{"tejus", "adsa@aed", 23})
	res, _ := c.GetAll()
	fmt.Println(res)

	c.Delete("thejus")
	res, _ = c.GetAll()
	fmt.Println(res)
	d, _ := c.GetById("shihad")
	fmt.Println(d)
}
