package main

import "fmt"

// Variadic function In GO: It's a function that accepts a variable number of argument values.
// syntax
// func f(parameterName ...Type)
// * the three dot ... is called pack operator. It tells Go to store all arguments of Type in param name

// example
func numbers(i ...int) {

	fmt.Println(i)

}

// ! Important NOTE: A variadic function can have multiple parameters. However if our function requires multiple parameters, the variadic function params must be the last in the function

// Incorrect example
// func nums(i ...int, name string){
// 	fmt.Println(i)
// 	fmt.Println(name)
// }

// correct example
func nums(name string, i ...int) {
	fmt.Println(name)

	for _, v := range i {
		fmt.Println(v)
	}

	fmt.Println() // new line

	fmt.Println(i)
	fmt.Printf("%T\n: ", i)         // []int
	fmt.Printf("Len: %d\n", len(i)) // 4
	fmt.Printf("Cap: %d\n", cap(i)) // 4
}

// example : sum function

func sum(nums ...int) int {

	total := 0
	for _, v := range nums {
		total += v
	}
	return total

}

func main() {

	numbers(12, 34, 5, 4, 564)
	numbers(12)
	numbers() // []

	// nums(12, 23, "Hello")
	nums("Hola", 3, 4, 5, 6)

	i := []int{1, 2, 3, 4}
	// numbers(i) // compile error

	// can be possible using unpack operator
	numbers(i...)

	// sum
	numSlice := []int{1, 2, 3, 4, 5, 6, 7}
	result := sum(numSlice...)
	fmt.Println("Total sum result: ", result)
	result2 := sum(23, 44, 45)
	fmt.Println("Total sum result2: ", result2)
}
