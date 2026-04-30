package main

import (
	"errors"
	"fmt"
)

func main() {
	// Demonstrating the panic by accessing the element of index which is not available
	/* // Uncomment code to run

	arr := []int{1, 2, 3}

	for i := 0; i < 5; i++ {
	fmt.Println(arr[i])
	}
	*/
	// OUTPUT:
	/* panic: runtime error: index out of range [3] with length 3 */

	// Another example:

	// msg := "good-bye"
	// message(msg)
	// fmt.Println("This line will not print because the message func throw panic")

	// Example: this code shows how panic and defer statement function together
	test()

}

// func message(msg string) {
// 	if msg == "good-bye" {
// 		panic(errors.New("Something went wrong"))
// 	}
// }

func test() {
	n := func() {
		fmt.Println("Defer in test func")
	}
	defer n()
	msg := "hello"
	message(msg)

}

func message(msg string) {

	f := func() {
		fmt.Println("Defer in message func")
	}
	defer f()

	if msg == "hello" {
		panic(errors.New("Something went wrong"))
	}
}
