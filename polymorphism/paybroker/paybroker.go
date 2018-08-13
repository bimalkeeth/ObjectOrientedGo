package paybroker

import "fmt"

type PaymentBrokerAccount struct{}

func (p *PaymentBrokerAccount) ProcessPayment(amount float32) bool {

	fmt.Println("Processing Payment through payment broker ...")
	return true
}
