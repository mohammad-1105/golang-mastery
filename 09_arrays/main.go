package main

import (
	"fmt"
	"sort"
)

/*
  - Arrays is a fixed length collection of elements of the same type.
	- It stored elements in contiguous memory locations.
*/

func main() {

	// creating array
	var arr [4]int
	fmt.Println(arr)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	// arr[4] = 5 // Error: Out of bounds
	fmt.Println(arr)

	// create and assign value to array
	myArr := [3]int{10, 20, 30}
	fmt.Println(myArr)

	// accessing array elements
	fmt.Println(myArr[0])
	fmt.Println(myArr[1])
	fmt.Println(myArr[2])

	// Iterating over and array
	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

	// we could also use range here (the first iteration variable is index and second one is value)
	for _, v := range arr {
		fmt.Println(v)
	}

	// checking the array length
	fmt.Println(len(arr))

	fmt.Println()
	slicingInArray()
	comparingArray()
	arraySorting()
	arrayCopy()
	arrayDelete()
	arrayContains()
}

// Array Slicing: Slicing is a technique to access subset of an array. In most programming language the syntax is almost same like array[startIndex:endIndex]

func slicingInArray() {
	scores := [6]int{
		98, 99, 87, 67, 68, 54,
	}

	fmt.Println(scores[0:5]) // exclude the last one
	fmt.Println(scores[1:4]) // [99, 87, 67]
	fmt.Println(scores[:])   // print all [98, 99, 87, 67, 68, 54,]
	fmt.Println(scores[0:])  // print all [98, 99, 87, 67, 68, 54,]
	fmt.Println(scores[:2])  // from 0 to 1 index [98, 99,]
}

// Array comparison: Array can be compare using == operator
func comparingArray() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [4]int{1, 2, 3, 5}

	fmt.Println(arr1 == arr2) // true
	fmt.Println(arr1 == arr3) // false

}

func arraySorting() {
	arr := [4]int{4, 3, 2, 1}
	fmt.Println(arr)

	sort.Ints(arr[:]) // unsorted [4, 3, 2, 1]
	fmt.Println(arr)  // sorted [1, 2, 3, 4]

	// sorting array using for loop

	arr1 := [5]int{8, 5, 3, 2, 1}

	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i] > arr1[j] {
				arr1[i], arr1[j] = arr1[j], arr1[i]
			}
		}
	}

	fmt.Printf("arr1 after sorted: %v\n", arr1)
}

func arrayCopy() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 0}

	n := copy(arr2[:], arr1[:])

	fmt.Printf("arr2 after copied: %v\n", arr2)

	fmt.Printf("copy function return the number of copied value so n: %d\n", n)
}

func arrayDelete() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr1[1] = 0
	fmt.Println(arr1) // [1, 0, 2, 3, 4]
}

// sort function to check if array contains a value
func arrayContains() {
	arr := [10]int{33, 43, 41, 12, 45, 67, 87, 98, 56, 43}

	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= 45 })

	if i < len(arr) && arr[i] == 45 {
		fmt.Println("Found 45 at index ", i)
	}

}
