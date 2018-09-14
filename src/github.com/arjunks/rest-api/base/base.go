package base

type Person struct {
	Fname    string `json: "fname"`
	Lname    string `json: "lname"`
	Password string `json: "password"`
	Image    string `json: "image"`
	Country  string `json: "country"`
	Age      int    `json: "age"`
	Id       int    `json: "id"`
}

type Operations interface {
	GetPeople() ([]Person, error)
	GetPerson(Id int, password string) (Person, error)
	CreatePerson(p Person) error
	UpdatePerson(p Person) error
	DeletePerson(Id int, password string) error
}
