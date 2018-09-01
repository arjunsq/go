package main

import (
	"fmt"
	"os"

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
	m := mapstore.MapStore{make(map[string]domain.Customer)}

	c := customerController{&m}
	ch := 0
	for {
		fmt.Printf("1. %s\t2. %s\t3. %s\t4. %s\t5. %s\n", "Add Customer", "Delete Customer", "Get A Customer", "Get All Customers", "Exit Application")
		fmt.Scanf("%d", &ch)
		switch ch {
		case 1:
			fmt.Println("Enter Customer Name,Email and age")
			name := ""
			email := ""
			age := 0
			fmt.Scanf("%s %s %d", &name, &email, &age)

			cus := domain.Customer{
				Name:  name,
				Email: email,
				Age:   age,
			}
			fmt.Println("here", cus)
			err := c.Add(cus)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			fmt.Println("Enter Customer Name")
			name := ""
			fmt.Scanf("%s", &name)
			err := c.Delete(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(name, "is deleted")
			}
		case 3:
			fmt.Println("Enter Customer Name")
			name := ""
			fmt.Scanf("%s", &name)
			cu, err := c.GetById(name)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Details of ", name, ": ", cu)
			}
		case 4:
			clist, err := c.GetAll()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				for _, v := range clist {
					fmt.Println(v)
				}
			}

		case 5:
			fmt.Println("PROGRAM EXITS")
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice , Enter right option")
		}
	}
}
