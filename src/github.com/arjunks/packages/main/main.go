package main

import (
	"fmt"
	"os"

	"github.com/arjunks/packages/source"
)

func main() {

	m := source.Manstore{}
	ch := 0

	for {
		fmt.Printf("1. %s\t2. %s\t3. %s\t4. %s\n", "Add man", "Retrive A man", "Get All man", "Exit Application")
		fmt.Scanf("%d", &ch)
		switch ch {
		case 1:
			fmt.Println("Enter Customer Name,Age,Secret and Password")
			tname := ""
			tage := 0
			tsecret := ""
			tpassword := ""

			fmt.Scanf("%s %d %s %s", &tname, &tage, &tsecret, &tpassword)

			tempman := source.Man{
				Name:     tname,
				Age:      tage,
				Secret:   tsecret,
				Password: tpassword,
			}
			err := m.Add(tempman)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			fmt.Println("Enter Customer Name and password")
			name := ""
			password := ""
			fmt.Scanf("%s %s", &name, &password)
			tempman, err := m.RetrieveSecret(name, password)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Man and his secret :", tempman)
			}

		case 3:
			fmt.Println("Details of Men", "\n", m)

		case 4:
			fmt.Println("PROGRAM EXITS")
			os.Exit(0)
		default:
			fmt.Println("Wrong Choice , Enter right option")
		}
	}
}
