package interface_example

import "fmt"

// PaymentProcessor defines the interface for payment methods
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

// CreditCard implements PaymentProcessor interface
type CreditCard struct {
	cardNumber string
}

func (c *CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f using card %s\n", amount, c.cardNumber)
	// Actual credit card payment logic...
	return nil
}

// PayPal implements PaymentProcessor interface
type PayPal struct {
	email string
}

func (p *PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f using email %s\n", amount, p.email)
	// Actual PayPal payment logic...
	return nil
}

func Interface2() {
	var processor PaymentProcessor

	processor = &CreditCard{cardNumber: "4111111111111111"}
	processor.ProcessPayment(100.50)

	processor = &PayPal{email: "user@example.com"}
	processor.ProcessPayment(200.75)
}
