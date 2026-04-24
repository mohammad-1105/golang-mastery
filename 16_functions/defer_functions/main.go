package main

import (
	"fmt"
	"os"
)

// * defer keyword in Go function:
/* * defer schedules a function call to run when surroundings functions returns.
* It always executes at the end of the current function
* Multiple defer calls run in LIFO order (stack behavior)
* Arguments to deferred calls are evaluated immediately, not at execution time
 */

// examples
func done() {
	fmt.Println("DONE")
}
func main() {

	// ! NOTE:  single defer used here
	defer done()
	fmt.Println("Main start")
	fmt.Println("Main end")
	/*
		Expected OUTPUT :
			Main start
			Main end
			DONE
	*/

	// * Multiple defer calls run in LIFO order (stack behavior)
	// Example: using multiple defer calls here

	fmt.Println("Start Again")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	defer fmt.Println("defer 4")
	fmt.Println("End Again")
	/*
		Expected OUTPUT:
			Start Again
			End Again
			defer 4
			defer 3
			defer 2
			defer 1
	*/

	//
	a()
	fmt.Println("Main after a")
	/*
		Expected OUTPUT:
			from a
			Main after a
	*/
	// here defer runs when func a returns, not after the end of main.

	// example: file handling
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error in file opening: ", err)
	}
	fmt.Println(f)
	defer f.Close()
}

// defer is function scoped not program scoped
func a() {
	defer fmt.Println("from a")
}
