package main

import "fmt"

// A contract
type paymentProcessor interface {
	pay(amount int) error
}

// structs of various payment providers
type paypal struct{}
type stripe struct{}
type razorPay struct{}

// methods for each struct
func (p paypal) pay(amount int) error {
	fmt.Println("Paid with paypal: ", amount)
	return nil
}
func (p stripe) pay(amount int) error {
	fmt.Println("Paid with stripe: ", amount)
	return nil
}
func (p razorPay) pay(amount int) error {
	fmt.Println("Paid with razorPay: ", amount)
	return nil
}

// business logic depend on interface
type checkout struct {
	processor paymentProcessor
}

func (c checkout) Process(amount int) {
	_ = c.processor.pay(amount)
}
func main() {
	c1 := checkout{
		processor: stripe{},
	}
	c1.processor.pay(100)

	c2 := checkout{
		processor: paypal{},
	}
	c2.processor.pay(200)

	c3 := checkout{
		processor: razorPay{},
	}
	c3.processor.pay(300)

}
