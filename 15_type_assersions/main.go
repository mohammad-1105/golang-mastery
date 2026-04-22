package main

import "fmt"

// * Type assertions and interfaces
/*
* We have used fmt.Print and its siblings a lot of time. It's a great deal for writing our code. But how does a function such as fmt.Print take any type of value if Go is a strongly typed language ? Let's take a look into actual Go standard code for fmt.Print

 */
/*
	Print formats using the default formats for its operands and writes
	to standard output.
	Spaces are added between operands when neither is a string.
	It returns the number of bytes written and any write error
	encountered.
	func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

*/

/*
* In Go interfaces{} (or any) can hold any value, but we lose access to its concrete type.
* Type assertion is how we say: "I know what's inside-give it back to me as that type"
* It's a runtime check , not compile time -> can panic if wrong
 */

func main() {

	// what interface{} actually stores

	// when we do this
	var v interface{}
	v = "hello" // v is not "a string". Internally it hold:
	// {type: string, value: "hello"}
	// But the compiler only sees interface{}, so it blocks us from doing
	// fmt.Println(v + " World") // compile error

	/* Reason:
	* compiler says: i don't know if this is a string. Could be int, struct, anything.
	 */

	fmt.Println(v)

	// Type Assertions
	var greet interface{} = "hello"

	s := greet.(string) // assert it's a string
	fmt.Println(s)
	fmt.Printf("Type of s: %T\n", s) // string

	var numValue interface{} = 10
	stringNum, ok := numValue.(string)

	if !ok {
		fmt.Println("Not a string")
	}
	fmt.Println(stringNum)
}

// func Print(a ...interface{})

/* * This means:
* Accept any number of values
* Each value can be any type
 */
