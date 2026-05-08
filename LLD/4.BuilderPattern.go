package main

import "fmt"

/* Builder Pattern:( Creational Design Pattern)
used to create complex objects step-by-step.
*/

type User struct {
	Name  string
	Age   int
	Email string
}

type UserBuilder struct {
	user User // as user is private so we can's use outside package directly without Build() method , and Build method return user
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (u *UserBuilder) SetName(name string) *UserBuilder {
	u.user.Name = name
	return u
}

func (u *UserBuilder) SetAge(age int) *UserBuilder {
	u.user.Age = age
	return u
}

func (u *UserBuilder) SetEmail(email string) *UserBuilder {
	u.user.Email = email
	return u
}

func (u *UserBuilder) Build() User {
	return u.user
}

func main() {

	user := NewUserBuilder().SetName("Shreyas").SetAge(25).SetEmail("xyz@gmail.com").Build()
	fmt.Println(user)
}

/*
Real-Time Uses Cases:
Database Query Builder
	query := Select("*").From("users").Where("age > 18").OrderBy("name").Limit(10).Build()
Kubernaties Config Builder:
	pod := NewPodBuilder().SetName("nginx").SetImage("nginx:latest").SetCPU("500m").SetMemory("1Gi").AddPort(80).Build()
*/
