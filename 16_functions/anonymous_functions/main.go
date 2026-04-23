package main

import "fmt"

// * Anonymous function is a function that do not have names. This functions use most of the time in our code.

/* Anonymous functions are used for :
* Closure implementations
* defer statements
* Defining a code block to be used with a goroutine
* Defining a function for a one-time use
* passing a function to another function
 */

func main() {
	// example : anonymous function
	func() {
		fmt.Println("Greeting")
	}()

	// passing argument to the anonymous function
	message := "Hello World"
	func(m string) {
		fmt.Println(m)
	}(message)

	// save anonymous function inside a variable
	f := func() {
		fmt.Println("Executing Anonymous function using a variable")
	}
	fmt.Println("Line after anonymous function declaration")
	f()

	// creating an anonymous function to calculate the square root of a number
	x := 9
	squareRoot := func(num int) int {
		return num * num
	}
	result := squareRoot(x)
	fmt.Printf("The square root of %d is %d\n", x, result)
}
