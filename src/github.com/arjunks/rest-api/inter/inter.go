package inter

import (
	"errors"

	"github.com/arjunks/rest-api/base"
)

type People map[int]base.Person

func (P *People) GetPeople() People {
	return *P
}

func (P *People) GetPerson(Id int, password string) (base.Person, error) {
	if (*P)[Id].Password == password {
		return (*P)[Id], nil
	}
	var novalue = base.Person{}
	return novalue, errors.New("Person not present")
}

func (P *People) CreatePerson(p base.Person) error {
	if _, ok := (*P)[p.Id]; ok {
		return errors.New("User already present")
	}
	(*P)[p.Id] = p
	return nil
}

func (P *People) UpdatePerson(p base.Person) error {
	if _, ok := (*P)[p.Id]; !ok {
		return errors.New("User not present")
	}
	(*P)[p.Id] = p
	return nil
}

func (P *People) DeletePerson(Id int, password string) error {
	if _, ok := (*P)[Id]; !ok {
		return errors.New("User not present")
	} else if (*P)[Id].Password != password {
		return errors.New("Invalid password entered")
	}
	delete(*P, Id)
	return nil
}
