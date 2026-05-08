package main

import "fmt"

/* Singleton Pattern:( Creational Design Pattern)
only one object(instance) is created
and provides a global access point to it.
Thread-Safe Singleton
Use: sync.Once

Advantages:
1. Single Shared Resource
	One config/logger/db manager.
2. Saves Memory

Disadvantage:
1. Tight Coupling

Real-Time Uses Cases:
	DB Connection
*/

type DBConn struct {
	connection string
}

var instance *DBConn

func GetConnection() *DBConn {
	if instance == nil {
		instance = &DBConn{connection: "DB Connection!!!"}
	}
	return instance
}

func main() {
	conn1 := GetConnection()
	conn2 := GetConnection()

	if conn1 == conn2 {
		fmt.Println("Same Conn reference")
	}

}


