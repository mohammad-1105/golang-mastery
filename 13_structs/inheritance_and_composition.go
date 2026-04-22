package main

// ! Warning: GO doesn't have Inheritance feature
/*
	* The designers of the Go ignore inheritance to implement. it hurts the devs coming from the world of OOPs languages.
	* Why GO designers avoids inheritance in Golang:
		* Tight coupling (child depends heavily on the parent internals)
		* Fragile base class problem (changing parents breaks children)
		* Implicit behavior (harder to reason about large systems)

	? But Go have alternative better which is Composition
	 * Explicit structure: you see what is exactly used
	 * Looser coupling: swap components easily
	 * Better control over API surface
	 * Encourages small, focused types

	? But remember there is a massive difference between inheritance and composition
	 * Inheritance :-> "is a" relationship (child automatically gets behavior from parents)
	 * Composition :-> "has a" relationship (you build types by combining smaller pieces)

	-> That's why: Go avoids inheritance and uses composition to keep coupling low and behavior explicit.
*/

// struct embedding and initialization

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

// function that returns slices of dots

func getDots() []dot {

	// first dot instance with var notation
	var dot1 dot

	// dot2 initializing with zero values
	dot2 := dot{}
	dot2.name = "A"

	// using promoted fields here
	dot2.x = 10
	dot2.y = 12
	dot2.width = 8
	dot2.height = 12

	// initializing embedded types: We can't use promotion like above here
	dot3 := dot{
		name: "B",
		location: location{
			x: 10,
			y: 12,
		},
		size: size{
			width:  23,
			height: 24,
		},
	}

	// for dot4 we will use type names to set data
	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 102
	dot4.size.width = 23
	dot4.size.height = 24

	return []dot{dot1, dot2, dot3, dot4}

}
