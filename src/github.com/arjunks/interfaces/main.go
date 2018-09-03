package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

type programmer struct {
	person
	lang []string
}

type agent struct {
	person
	licenseToKill bool
}

type men interface {
	describe()
}

func talk(m men) {
	m.describe()
}

func (p programmer) describe() {
	p.person.sayName()
	fmt.Println("Known languages:", p.lang)
}

func (a agent) describe() {
	a.person.sayName()
	if a.licenseToKill == true {
		fmt.Println("And has license to kill")
	} else {
		fmt.Println("But doesnt have license to kill")
	}
}

func (p person) sayName() {
	fmt.Println("I am", p.lname, ",", p.fname, p.lname)
}

func main() {

	james := agent{
		person{
			"James",
			"Bond",
			40,
		},
		true,
	}
	arjun := programmer{
		person{
			"Arjun",
			"K S",
			21,
		},
		[]string{"Java", "C", "Go", "Node.js"},
	}
	talk(james)
	talk(arjun)
}
