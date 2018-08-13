package main

import "ObjectOrientedGo/polymorphism/paypol"

func main() {

	var option paypol.PaymentOption
	option = &paypol.CreditCard{}
	option.ProcessPayment(500)
	option = &paypol.CheckingAccount{}
	option.ProcessPayment(200)

}
