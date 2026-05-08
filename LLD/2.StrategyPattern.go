package main

import (
	"fmt"
)

/* Strategy Pattern:( Behavioral Design Pattern)
define multiple algorithms/behaviors,
switch obj dynamically at runtime.
means: at runtime: like suppose at runtime in phonepay : switch payment pay option from icici cred card to Hdfc debit card and then pay.
So there we set startrgey(context) at runtime and also chnages that startegy

Advantages:
At runtime:
	object changes, behavior changes automatically.
	Open/Closed Principle:
		Add new strategy without modifying old code. ( no need to changes anything now even Startegry context, just easily add new one)

Real-Time Uses Cases:
Payment Gateway
*/

type PaymentStrategy interface {
	Pay(Amount int)
}

type CreditCard struct{}

func (card *CreditCard) Pay(Amount int) {
	fmt.Println("CreditCard: Payment Done: ", Amount)
}

type UPI struct{}

func (card *UPI) Pay(Amount int) {
	fmt.Println("UPI: Payment Done: ", Amount)
}

type Paypal struct{}

func (card *Paypal) Pay(Amount int) {
	fmt.Println("Paypal: Payment Done: ", Amount)
}

// Startgery Ctx
type ShoppingCart struct {
	payment PaymentStrategy
}

func (s *ShoppingCart) SetPaymentMethod(p PaymentStrategy) {
	s.payment = p
}

func (s *ShoppingCart) Checkout(amount int) {
	s.payment.Pay(amount)
}

func main() {
	shopCart := ShoppingCart{}

	shopCart.SetPaymentMethod(&CreditCard{})
	shopCart.Checkout(100)

	shopCart.SetPaymentMethod(&Paypal{})
	shopCart.Checkout(1009)

}


