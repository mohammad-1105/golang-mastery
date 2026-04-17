package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Go’s if is straightforward but does not require parentheses, and it supports a short statement before the condition—commonly used for error handling.
*/
func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x == 5 {
		fmt.Println("x is equal to 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// if with short statement
	if y := x * 2; y > 15 {
		fmt.Println("y is greater than 15:", y)
	}

	if err := doSomething(true); err != nil {
		fmt.Println("Error ", err)
		return
	}
	fmt.Println("success")

	/* Don't do like this:

		err := doSomething()
		if err != nil {
			fmt.Println(err)
		}
	This leaks the err into a wider scope unnecessary
	*/

	// another example of if else with err

	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Println("Num: ", num)
	} else {
		fmt.Println("Error: ", err)
	}
}

// Basic error handling flow with if statement
func doSomething(flag bool) error {
	if !flag {
		return errors.New("Something went wrong")
	}
	return nil
}
