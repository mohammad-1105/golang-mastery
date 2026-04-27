package main

import (
	"fmt"
	"strconv"
)

// * There are three types of Errors we might encounter:
// 1. Syntax Errors
// - Caused by invalid Go syntax (missing brackets, wrong structure, etc.)
// - Detected at compile time
// - Program will not run until fixed

// 2. Runtime Errors (Panics)
// - Occur while the program is running
// - Caused by invalid operations (nil pointer, index out of range, divide by zero)
// - Crash the program unless handled with recover()

// 3. Semantic Errors (Logical Errors)
// - Code compiles and runs, but produces incorrect results
// - Caused by wrong logic or flawed assumptions
// - Hardest to detect; require testing and debugging

// * Other programming languages like Java, Python, JavaScript they use try-catch or exception but Go use different approach:
/*
* Errors are explicit in GO
* We can see exactly where failure is handled
* No hidden control flow like try-catch or exception
 */

// ! Most Important: Errors are values in GO: which means error is just a value not special like exceptions.
/* * In GO we can:
* store error
* compare error
* return error
* pass error
 */

// Error is go must satisfy the error interface
type error interface {
	Error() string
}

// ! NOTE: any type is an error if it has this method: Error() string

func main() {
	// example

	v := "10"

	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T %v\n", s, s)
	} else {
		fmt.Println(err)
	}

	v = "s2"
	s, err := strconv.Atoi(v)
	if err != nil {
		fmt.Println(s, err)
	}

}
