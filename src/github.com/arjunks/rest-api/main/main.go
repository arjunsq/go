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

var p People

func (p *inter.People) initial() {
	p[1] = base.Person{
		Fname:    "arjun",
		Lname:    "ks",
		Password: "123",
		Image:    "arjun.jpg",
		Country:  "India",
		Age:      21,
		Id:       1,
	}
	p[2] = base.Person{
		Fname:    "shiahad",
		Lname:    "thayyil",
		Password: "123",
		Image:    "shihad.jpg",
		Country:  "India",
		Age:      22,
		Id:       2,
	}
}

func main() {
	p := make(map[int]base.Person)
	p.initial()
	r := mux.NewRouter()

	r.HandleFunc("/getpeople", GetPeopleHandler).Methods("GET")
	r.HandleFunc("/getperson", GetPersonHandler).Methods("GET")
	r.HandleFunc("/createperson", CreatePersonHandler).Methods("POST")
	r.HandleFunc("/updateperson", UpdatePersonHandler).Methods("PUT")
	r.HandleFunc("/deleteperson", DeletePersonHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Fprintln(w, string(bs))
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetPersonHandler invoked")
}

func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CreatePersonHandlinitial invoked")
}
func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdatePersonHandlinitial() invoked")

}
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeletePersonHandlinitial() invoked")

}
