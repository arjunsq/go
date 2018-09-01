package source

import "errors"

type Man struct {
	Name     string
	Age      int
	Secret   string
	Password string
}

type Manstore []Man

type Manoperations interface {
	Add(s Man) error
	RetrieveSecret(name string, password string) (Man, error)
}

func (m *Manstore) Add(s Man) error {
	for _, val := range *m {
		if s.Name == val.Name {
			return errors.New("Man already present")
		}
	}
	*m = append(*m, s)
	return nil

}
func (m *Manstore) RetrieveSecret(name string, password string) (Man, error) {
	d := Man{}
	for _, val := range *m {
		if name == val.Name {
			if password == val.Password {
				return val, nil
			}
		}
	}
	return d, errors.New("Unregistered man or invalid login credentials")
}
