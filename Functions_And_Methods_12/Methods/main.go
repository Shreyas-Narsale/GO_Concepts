package main

import "fmt"

type User struct {
	Name    string
	Surname string
	Status  string
	Email   string
	Age     int
}

func main() {
	var shreyas = User{"Shreyas", "Narsale", "Active", "shree@gmail.com", 25}
	shreyas.GetStatus()

	err := shreyas.UpdateAge(22)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated age is:", shreyas.Age)
	}
}

func (user *User) GetStatus() {
	fmt.Println("Status is:", user.Status)
}

func (user *User) UpdateAge(age int) error {
	user.Age = age
	return nil
}
