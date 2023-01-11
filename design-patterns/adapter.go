package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type CardPayment struct{}

func (CardPayment) Pay(vendor string) {
	fmt.Printf("Card Payment with vendor %s\n", vendor)
}

type CardPaymentAdapter struct {
	CardPayment *CardPayment
	cardVendor  string
}

func (cpa CardPaymentAdapter) Pay() {
	cpa.CardPayment.Pay("VISA")
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	card := &CardPaymentAdapter{CardPayment: &CardPayment{}, cardVendor: "VISA"}
	ProcessPayment(card)
}
