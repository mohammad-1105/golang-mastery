package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ! Always remember: In GO, an error is just a value
// That means:
/**
* Not Special
* Not exceptions
* Just like int, string, etc
 */

// Error interface.
type error interface {
	Error() string
}

// Any type that has Error() string is an error.

// Let's build our own error

// custom type
type MyError struct {
	msg string
}

// Implement error interface
func (e MyError) Error() string {
	return e.msg
}

func main() {
	var err error = MyError{msg: "Something went wrong"}
	fmt.Println(err)         // Something went wrong
	fmt.Println(err.Error()) // Something went wrong

	// * `MyError` is now a valid error because it implements `Error() string`
	// * This is called the interface based designed not like Java inheritance

	// * Case study : Why Go does this ?
	/*
		* Example in Code like Java we do:
					throw new Exception("fail")
					* Control jumps
					* Flow is hidden
		* But in GO,
					return value, err
					* flow is explicit
					* No surprise
	*/

	// Examples from standard package

	s, err := strconv.Atoi("1234")

	fmt.Println(s)   // 1234
	fmt.Println(err) // <nil>

	// proper handling

	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(s)

	// Don't try to ignore errors: example

	str, _ := strconv.Atoi("abc")
	fmt.Println("str: ", str) // 0 (wrong assumptions)

	// usage of ageValidation func
	er := ageValidation(29)
	if er != nil {
		fmt.Println(er)
	}
}

// Let's see How Go creates error internally

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// Understand carefully:
// - `errorString` is an struct
// - It becomes an error because it implements `Error()`

// * create a New function like we have in GO errors

func New(text string) error {
	return &errorString{text}
}

/*
	SO,
	- We pass string
	- Go wraps in it a struct
	- Return it as error
*/

// Real world usage:

// validation
func ageValidation(age int) error {
	if age > 20 {
		return errors.New("You are not a child so please Go and study")
	}
	return nil
}
