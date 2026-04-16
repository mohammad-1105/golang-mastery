package main

import "fmt"

/*
	- Scope in Go defines where a variable is accessible (block-level, file-level, package-level)
	- Go uses Lexical scoping: visibility is determined by where code is written, not runtime
	- Shadowing happens when redeclare a variable in inner scope using :=, hiding the outer one, often a subtle bug source.
*/

func scopes() {

	// Basic scope
	x := 10 // scope: main function

	if true {
		y := 20        // scope: this if block
		fmt.Println(x) // 10
		fmt.Println(y) // 20
	}

	// fmt.Println(y) // compile error: undefined y

}

// function scope vs package scope

var global string = "I am global" // package level scope
func scopesTwo() {

	local := "I am local" // function level scope

	fmt.Println(global) // I am global
	fmt.Println(local)  // I am local
}

func another() {
	fmt.Println(global) // I am global
	// fmt.Println(local)  // Error: Undefined
}
