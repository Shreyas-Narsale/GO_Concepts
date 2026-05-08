package main

import (
	"fmt"
)

/* factory Pattern:
used to create objects without exposing the exact creation logic to the client.
you create objects through a factory function/method.
*/

type Payment interface {
	Pay(Amount int)
}

type CreditCard struct{}

func (card *CreditCard) Pay(amount int) {
	fmt.Println("CreditCard: Payment Done: ", amount)
}

type UPI struct{}

func (card *UPI) Pay(amount int) {
	fmt.Println("UPI: Payment Done: ", amount)
}

type Paypal struct{}

func (card *Paypal) Pay(amount int) {
	fmt.Println("Paypal: Payment Done: ", amount)
}

// Factory Method
func GetPaymentMethod(PaymentType string) Payment {
	switch PaymentType {
	case "CreditCard":
		return &CreditCard{}
	case "UPI":
		return &UPI{}
	case "Paypal":
		return &Paypal{}
	default:
		return nil
	}
}

func main() {

	obj := GetPaymentMethod("CreditCard")
	obj.Pay(100)

	obj = GetPaymentMethod("UPI")
	obj.Pay(100)

}

/*Advantages:
Loose Coupling,
Easy to Extend: Add new payment method,
Centralized Object Creation: All creation logic in one place.
*/

/*
Real-Time Uses Cases:
Database Drivers
	sql.Open("mysql")
	sql.Open("postgres")

	Driver selected internally.

Logger Creation
	GetLogger("file")
	GetLogger("console")
Cloud Providers
	GetStorage("aws")
	GetStorage("gcp")

*/
