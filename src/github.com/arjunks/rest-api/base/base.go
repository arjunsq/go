package base

type Person struct {
	Fname, Lname, Password, Image, Country string
	Age, Id                                int
}

type Operations interface {
	GetPeople() ([]Person, error)
	GetPerson(Id int, password string) (Person, error)
	CreatePerson(p Person) error
	UpdatePerson(p Person) error
	DeletePerson(Id int, password string) error
}
