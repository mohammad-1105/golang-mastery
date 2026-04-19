package main

import (
	"fmt"
	"time"
)

func gettingPointer() {
	// declare a pointer using var keyword
	var count1 *int
	// fmt.Println(count1) // nil (default)

	// create a variable using new
	count2 := new(int) // new() function is intended to be used to get some memory for a type and return a pointer to that address.
	// fmt.Println(count2) // 0x311d9d76a0a8

	// we can't take the address of the literal number so we create a variable to hold the number
	countTemp := 10
	count3 := &countTemp

	// It’s possible to create a pointer from some types without a temporary variable. Here, we’re using our trusty time struct
	t := &time.Time{}

	// print all

	fmt.Printf("Count1: %#v\n", count1) // Count1: (*int)(nil)
	fmt.Printf("Count2: %#v\n", count2) // Count2: (*int)(0x25913800a0a8)
	fmt.Printf("Count3: %#v\n", count3) // Count3: (*int)(0x25913800a0b0)
	fmt.Printf("Time: %#v\n", t)        // Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

}

// ! Important NOTE:
// * To get value the pointer associated with it. We must dereference value using * in front of the variable name.
// * Dereferencing the zero or a nil pointer is a common bug in GO software as the compiler won't warning us about it. So, it’s always best practice to check that a pointer is not nil before dereferencing it unless you are certain it’s not nil.

func gettingValueFromPointer() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("Count1: %#v\n", *count1)
	}

	// No need to check — always non-nil
	fmt.Printf("Count2: %#v\n", *count2)
	fmt.Printf("Count3: %#v\n", *count3)
	fmt.Printf("Time: %#v\n", *t)
}

// pointer value swap

func swap(a *int, b *int) {
	*b, *a = *a, *b
}

type User struct {
	Name string
	Age  int
}

func updateAge(u *User, newAge int) {
	u.Age = newAge
}

// pointers helps use to avoid expensive copies
// Example
type BigPayload struct {
	Data [1024 * 1024]byte // 1MB
}

// Bad version (hidden cost)
func process(payload BigPayload) {
	// Huge 1MB copy created here
	fmt.Println(len(payload.Data))
}

// better way
func processPayload(payload *BigPayload) {
	fmt.Println(len(payload.Data))
}
