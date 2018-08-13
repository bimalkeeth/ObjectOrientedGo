package payment

import "fmt"

type Cash struct{}

func CreateCashAccount() *Cash {
	return &Cash{}
}
func (c Cash) ProcessPayment(amount float32) bool {

	fmt.Println("Processing cash transaction...")
	return true
}
