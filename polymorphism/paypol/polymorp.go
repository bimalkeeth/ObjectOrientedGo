package paypol

import "fmt"

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type CreditCard struct{}

func (c CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing credit card account ...")
	return true
}

type CheckingAccount struct{}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	fmt.Println("Processing Checking account ...")
	return true
}
