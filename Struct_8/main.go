package main

import "fmt"

type Details struct {
	Name    string
	Surname string
	email   string // private only accesible within this package
	Age     int
}

func main() {
	// member by list (seq matter)
	var shreyas = Details{"shreyas", "narsale", "shree@gmail.com", 25}

	fmt.Printf("Struct data is : %v \n", shreyas)
	fmt.Printf("Struct data is : %+v \n", shreyas)
	fmt.Println(shreyas.Age)
	fmt.Println(shreyas.email)

	var tushar Details
	tushar.Name = "tushar"
	tushar.Surname = "dj"
	tushar.email = "tu@gmail.com"
	tushar.Age = 55

	fmt.Printf("Struct data is : %v \n", tushar)
	fmt.Printf("Struct data is : %+v \n", tushar)
}
