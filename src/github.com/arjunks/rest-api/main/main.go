package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arjunks/rest-api/base"
	"github.com/arjunks/rest-api/inter"
	"github.com/gorilla/mux"
)

type p inter.People

func initial(p *inter.People) {
	(*p)[1] = base.Person{
		Fname:    "arjun",
		Lname:    "ks",
		Password: "123",
		Image:    "arjun.jpg",
		Country:  "India",
		Age:      21,
		Id:       1,
	}
	(*p)[2] = base.Person{
		Fname:    "shihad",
		Lname:    "thayyil",
		Password: "123",
		Image:    "shihad.jpg",
		Country:  "India",
		Age:      22,
		Id:       2,
	}
}

var pstore inter.People

func main() {
	pstore = make(map[int]base.Person)

	initial(&pstore)
	r := mux.NewRouter()

	r.HandleFunc("/getpeople", GetPeopleHandler).Methods("GET")
	r.HandleFunc("/getperson", GetPersonHandler).Methods("POST")
	r.HandleFunc("/createperson", CreatePersonHandler).Methods("POST")
	r.HandleFunc("/updateperson", UpdatePersonHandler).Methods("PUT")
	r.HandleFunc("/deleteperson", DeletePersonHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPeopleHandler invoked")
	bs, err := json.MarshalIndent(pstore.GetPeople(), "", "  ")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Fprintln(w, string(bs))
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPersonHandler invoked")
	var t struct {
		Id       int    `json:"id"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err, "hee")
	}
	fmt.Println("Value got is", t)
	person, err := pstore.GetPerson(t.Id, t.Password)
	if err != nil {
		fmt.Println(err, "hee")
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, person)
}

func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePersonHandlinitial invoked")
	var t base.Person
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err, "hee")
	}
	fmt.Println("Value got is", t)
	err = pstore.CreatePerson(t)
	if err != nil {
		fmt.Println(err, "hee")
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "Created Person")
}
func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdatePersonHandlinitial() invoked")

}
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeletePersonHandlinitial() invoked")

}
