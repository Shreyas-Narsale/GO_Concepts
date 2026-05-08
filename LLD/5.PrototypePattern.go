package main

import "fmt"

/* Prototype Pattern:( Creational Design Pattern)
Instead of building from scratch, you clone an existing object (prototype) and then modify it if needed.

Real-Time Uses Cases:
API Request Templates (Very common in Go services):
	When calling external APIs:
	Instead of building request every time:
	baseRequest → clone → modify headers/body

*/

type Prototype interface {
	Clone() Prototype
}

type User struct {
	Name  string
	Age   int
	Email string
}

func (user *User) Clone() Prototype {
	return &User{
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}
}

func main() {
	user1 := User{
		Name:  "Shreyas",
		Age:   25,
		Email: "xyz@gmail.com",
	}

	user2 := user1.Clone().(*User)

	user1.Age = 26

	user3 := user1.Clone().(*User)

	fmt.Println("User1 Name: ", user1.Name)
	fmt.Println("User1 Age: ", user1.Age)
	fmt.Println("User1 Email: ", user1.Email)

	fmt.Println("User2 Name: ", user2.Name)
	fmt.Println("User2 Age: ", user2.Age)
	fmt.Println("User2 Email: ", user2.Email)

	fmt.Println("User3 Name: ", user3.Name)
	fmt.Println("User3 Age: ", user3.Age)
	fmt.Println("User3 Email: ", user3.Email)

}


