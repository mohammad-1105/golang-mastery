package main

import "fmt"

/*
	iota is Go's ways to generate sequential, typed constants inside a `const` block. It's not an Enum keyword so don't confuse with that. But it gives enum like behavior clean, safe and compile time behavior

	Just simple: without `iota` we manually assigns value and with `iota` GO automatically assigns value

*/

// Examples: Without iota
func httpStatusCodeWithoutIota() {
	type Status int

	// creating enums of HTTP status codes
	const (
		StatusOK         Status = 200
		StatusCreated    Status = 201
		StatusBadRequest Status = 400
		StatusNotFound   Status = 404
	)

	fmt.Printf("HTTP OK Status Code %d\n", StatusOK) // 200

}

// Example: with iota (Idiomatic Go enum pattern)
func httpStatusCodeWithIota() {
	type Status int

	// const (
	// 	StatusOK         Status = iota // 0
	// 	StatusCreated                  // 1
	// 	StatusBadRequest               // 2
	// 	StatusNotFound                 // 3
	// )

	// fmt.Print(StatusNotFound) // 3

	// The above code have problems that we don't have these 0,1,2,3 status codes so we customize it.
	const (
		StatusOK         Status = 200 + iota // 200
		StatusCreated                        // 201
		StatusBadRequest                     // 202 // breaks here
		StatusNotFound                       // 203 // breaks here
	)

	fmt.Print(StatusNotFound) // 203

	// The above code breaks after StatusCreated because we don't have 202 , 203 status code for bad request and not found. So, use `iota` when values are sequential or pattern based
}

func main() {
	httpStatusCodeWithoutIota()
	httpStatusCodeWithIota()

	// Better iota examples
	type OrderStatus int

	const (
		Pending OrderStatus = iota
		Confirmed
		Shipped
		Delivered
		Cancelled
	)
}
