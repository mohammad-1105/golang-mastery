package main

import "fmt"

// * Closure are a form of anonymous function. Regular function cannot reference variables outside of themselves, however an anonymous function can reference variable external to their definition

func incrementer() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func decrementer(i int) func() int {
	return func() int {
		i--
		return i
	}
}

func main() {
	/*
		i := 0

		incrementor := func() int {
			i += 1
			return i
		}
		fmt.Println(incrementor()) // 1
		fmt.Println(incrementor()) // 2

		i += 10

		fmt.Println(incrementor()) // 13

	*/
	// Problem with the above code: we don't want i to be changed by anyone except the incrementor function but it is changed because it is in the main function

	increment := incrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	decrement := decrementer(10)
	fmt.Println(decrement())
	// fmt.Println(decrement())
	// fmt.Println(decrement())
	// fmt.Println(decrement())
	// fmt.Println(decrement())
	// fmt.Println(decrement())

	fmt.Println() // new line

	double := multiplayer(2)
	triple := multiplayer(3)

	fmt.Println(double(12))
	fmt.Println(triple(14))
}

// dependency injection using closure
func multiplayer(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
