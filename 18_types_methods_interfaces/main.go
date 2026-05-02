package main

import (
	"fmt"
	"math"
)

// interfaces In GO: An interface in GO defines behavior.
// * It's main purpose is to decouple the code from concrete types
// * If type implements the required methods, it automatically satisfies the Go interface
// * NOTE: The idiomatic way to name interfaces with suffix "er" like: Incrementer for Increment, Speaker for Speak, etc..

// Example : 1 starts here
type Shaper interface {
	Area() float32
}

type Rectangle struct {
	width, height float32
}

func (r Rectangle) Area() float32 {
	return r.width * r.height
}

type Square struct {
	length float32
}

func (s Square) Area() float32 {
	return s.length * s.length
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// Here is the main usage of interface
func calculateArea(s Shaper) float32 {
	return s.Area()
}

// Example : 1 ends here

// Example : 2 starts here
// Payment system
// behavior contract
type PaymentProcessor interface {
	Pay(amount int) error
}

// concrete implementation
type PayPal struct{}

func (p PayPal) Pay(amount int) error {
	fmt.Println("Paid with paypal: ", amount)
	return nil
}

type Stripe struct{}

func (s Stripe) Pay(amount int) error {
	fmt.Println("Paid with Stripe: ", amount)
	return nil
}

type RazorPay struct{}

func (s RazorPay) Pay(amount int) error {
	fmt.Println("Paid with RazorPay: ", amount)
	return nil
}

// business logic depends on interface
type Checkout struct {
	processor PaymentProcessor
}

func (c Checkout) Process(amount int) {
	c.processor.Pay(amount)
}

// Example : 2 ends here
func main() {

	// Example : 1 starts here
	areaOfRect := Rectangle{width: 10, height: 12}
	areaOfSquare := Square{length: 12}
	areaOfCircle := Circle{radius: 13}

	fmt.Println("Rectangle Area: ", calculateArea(areaOfRect))
	fmt.Println("Square Area: ", calculateArea(areaOfSquare))
	fmt.Println("Circle Area: ", calculateArea(areaOfCircle))
	// Example : 1 Ends here

	// Example : 2 starts here
	c1 := Checkout{processor: PayPal{}}
	c1.Process(100)
	c2 := Checkout{processor: Stripe{}}
	c2.Process(200)
	c3 := Checkout{processor: RazorPay{}}
	c3.Process(900)
	// Example : 2 ends here
}
