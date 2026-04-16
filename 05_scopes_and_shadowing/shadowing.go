package main

import "fmt"

func shadowing() {
	x := 10

	if true {
		x := 20        // shadowing x
		fmt.Println(x) // 20
	}

	fmt.Println(x) // 10 the outer x unchanged
}

// Real world bug here
func shadowingTwo() {
	x := 10

	x, y := 20, 30 // shadowing x and creating new y

	fmt.Println(x, y) // 20 30

	// z := 10
	// z := 20 // compile error: z redeclared in this block

	// hidden bug in if initialization
	a := 10

	if a := 5; a < 10 {
		fmt.Println(a) // 5, not 10
	}
	fmt.Println(a) // 10, outer a unchanged
}
