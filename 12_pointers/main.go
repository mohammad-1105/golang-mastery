package main

import "fmt"

// Pointer: A pointer stores the memory address of a value.

// To understand pointers, we first need to understand pass-by-value vs.
// pass-by-reference concepts. See this article:
// https://medium.com/the-bug-shots/understanding-value-and-pointer-receivers-in-golang-82dd73a3eef9

// Go is always pass-by-value.
/*
   When we create a variable and pass it to a function, Go copies the value.
   Any changes made to the parameter inside the function do NOT affect the original variable.
   This is called "pass by value".

   If we want to modify the original variable, we pass its address (a pointer).
   The function still receives a copy—but this time it's a copy of the address,
   so both the caller and callee refer to the same underlying memory.
*/

// Example: Value semantics

// In this function, x is a new variable with its own memory
func increment(x int) {
	fmt.Printf("%d incremented to %d inside the increment function\n", x, x+1)
	x++
}

// Pointer semantics

// Here, y holds the address of b passed from main
// *y dereferences the pointer and modifies the original value
func incrementUsingPointer(y *int) {
	*y++
	fmt.Println(*y)
}

func main() {
	var a int = 10
	increment(a)
	fmt.Printf("%d remains the same after calling increment because it was passed by value\n", a)

	var b int = 12
	incrementUsingPointer(&b)
	fmt.Println(b)

	gettingPointer()

	gettingValueFromPointer()

	x, y := 10, 20
	swap(&x, &y)
	fmt.Println(x == 20, y == 10) // true true (if swap function worked well)

	// update age of user
	user := User{Name: "Mohammad Ali", Age: 22}
	updateAge(&user, 23)
	fmt.Println(user.Age)

	payload := BigPayload{}
	process(payload)
	processPayload(&payload)

}
