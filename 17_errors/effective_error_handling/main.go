package main

import (
	"errors"
	"fmt"
)

/*
	! NOTES:
	* Rule no 1. Always remember Errors are values in GO.

	* Handling error in Go is little bit different than other programming languages like JavaScript, Python, Java. Go's built-in "errors" package don't contain the stack trace or nor they support the try-catch blocks like javaScript
*/

// * The Error Type
type error interface {
	Error() string
}

// The above interface type means: anything that implements `Error()` method and returns error as string is an error in GO

// * Constructing errors
// This is the static errors. And it can be constructed using GO's built-in fmt and errors package.
// func doSomething() error {
// 	return errors.New("Something went wrong")
// }

// using fmt package
// func Divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("Can't divide %d by zero", a)
// 	}

// 	return (a / b), nil
// }

// defining sentinel errors:
var ErrNegativeNumMultiply = errors.New("Negative number passed to multiply")

func multiply(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrNegativeNumMultiply
	}
	return a * b, nil
}

// Defining Custom error
type DivisionError struct {
	IntA int
	IntB int
	Msg  string
}

// receiver pointer for Divide Error
func (d *DivisionError) Error() string {
	return d.Msg
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Msg:  fmt.Sprintf("Cannot divide %d by zero", a),
			IntA: a,
			IntB: b,
		}
	}
	return a / b, nil
}

func handlingDivide() {
	a, b := 10, 2
	result, err := Divide(a, b)

	if err != nil {
		var divErr *DivisionError
		switch {
		case errors.As(err, &divErr):
			fmt.Printf("%d / %d is mathematically Invalid: %s\n",
				divErr.IntA, divErr.IntB, divErr.Error(),
			)
		default:
			fmt.Printf("Unexpected division error: %s\n", divErr)
		}
		return
	}
	fmt.Printf("%d / %d = %d\n", a, b, result)

}

func main() {

	// use case of multiple func
	a, b := 10, 3
	result, err := multiply(a, b)

	if err != nil {
		switch {
		case errors.Is(err, ErrNegativeNumMultiply):
			fmt.Println("Please don't pass negative number to multiply")
		default:
			fmt.Printf("Unexpected error: %v\n", err)
		}
		return
	}
	fmt.Printf("%d * %d = %d\n", a, b, result)

	handlingDivide()
}
