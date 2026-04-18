package main

import (
	"fmt"
	"slices"
)

/** Slices In GO:
- Slice is a dynamic, flexible view into an underlying array.
Unlike arrays slices can grow and shrink dynamically,
- Slices are more versatile and commonly used in Go. Most of the time we deal with slices instead of arrays.
*/

/* Slice structure
- A slice consists of three key components:
	* Pointer to the underlying array
	* Length of the slice
	* Capacity of the slice
*/

func main() {

	// 1. creating slice using slice literal
	fruits := []string{"apple", "banana", "guava"}
	fmt.Println(fruits)

	// 2. Using make() function
	numSlice1 := make([]int, 4)     // Length 4, capacity 4
	numSlice2 := make([]int, 5, 10) // length 5, capacity 10

	fmt.Println(numSlice1)
	fmt.Println(numSlice2)

	// difference between slice creation using literal and make function
	a := []int{1, 2, 4, 5, 6} // explicit values
	b := make([]int, 4)       // Zero values

	fmt.Println(a) // [1, 2, 4, 5, 6]
	fmt.Println(b) // [0, 0, 0, 0]
	/*
	* Literal --> Known data upfront (you control contents)
	* make -> you plan to fill it later (Go initializes with zero values)
	 */

	// buf := make([]byte, 1024) // buffer for reading file
	// fmt.Println(buf)

	// 3. Create slice from existing array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // creates slice from index 1 to 3
	fmt.Println(slice1)

	// slice properties
	fmt.Println(len(slice1)) // length of the slice
	fmt.Println(cap(slice1)) // capacity of the slice

	// empty slice with nil pointer
	var s []int
	fmt.Println(s)

	sliceLengthAndCapacity()
	copySlices()
	sliceIndexingAndSlicing()
	newSlice := insert([]int{1, 2, 3, 4}, 2, 99)
	fmt.Println(newSlice)

	newSliceAfterRemove := remove([]int{1, 2, 3, 4}, 1)
	fmt.Println("NewSliceAfterRemove: ", newSliceAfterRemove)

	playingWithSlicesUsingSlicesPackage()

	evenNumber := filterEven([]int{2, 3, 4, 5, 6, 7, 8})
	fmt.Println("Even Number: ", evenNumber)

	value, err := safeAccess([]int{1, 2, 3, 4, 5}, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}

/* len and cap in slice:
* len (length) = number of elements currently in the slice
* cap (capacity) = how many elements the slice can hold before reallocating
 */
func sliceLengthAndCapacity() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	s := arr[1:4] // elements [2, 3, 4]

	fmt.Println("slice: ", s)
	fmt.Println("len: ", len(s)) // 3
	fmt.Println("cap: ", cap(s)) // 6

	// here why cap = 6
	/*
	 * slice starts at index 1
	 * underlying array size 7
	 * capacity = from index 1 to end(7-1 = 6)
	 */

	// slice growth behavior
	newSlice := []int{1, 2, 3}

	fmt.Println(len(newSlice), cap(newSlice)) // 3, 3

	newSlice = append(newSlice, 4)
	fmt.Println(len(newSlice), cap(newSlice)) // 4, 6 (or more, depends on runtime)

	// what happened here:
	/*
	 * Original capacity was full -> Go allocates new array
	 * Copies old data -> appends new element
	 * Now slice points to new backing array
	 */

}

func copySlices() {
	original := []int{1, 2, 3, 4, 5}

	copied := make([]int, len(original))
	copy(copied, original)

	fmt.Println("original array: ", original)
	fmt.Println("copied slice: ", copied)
}

// slice manipulation
func sliceIndexingAndSlicing() {
	// slice indexing and slicing works same as we do in arrays

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s[:])         // [1, 2, 3, 4, 5]
	fmt.Println(s[1:3])       // [2, 3]
	fmt.Println(s[:len(s)-1]) // [1, 2, 3, 4]
	fmt.Println(s[2:])        // [3, 4, 5]

}

// slice modification techniques
// inserting elements
func insert(slice []int, index int, value int) []int {

	// slice = append(slice[:index], append([]int{value}, slice[index:]...)...)

	// return slice

	// More better way for explanation
	// Example:
	// slice = [1, 2, 3, 4]
	// index = 2
	// value = 99

	// step 1: Take left part (before index)
	left := slice[:index] // [1, 2]

	// step 2: Take right part (from index onward)
	right := slice[index:] // [3, 4]

	// step 3: Create a new slice with value at front
	// []int{value} -> [99]
	// append([99], right...) --> [99, 3, 4]
	newPart := append([]int{value}, right...)

	// step 4: Combine left + newPart
	// append([1, 2], [99, 3, 4]...)
	// -> [1, 2, 99, 3, 4]
	result := append(left, newPart...)

	return result
}

// deleting elements
func remove(slice []int, index int) []int {
	// return append(slice[:index], slice[index+1:]...)

	// step by step flow
	// Example:
	// slice = [10, 20, 30, 40, 50]
	// index = 2 (we want to remove 30)

	// step 1: left part (before index)
	// slice[:2] -> [10, 20]
	left := slice[:index]

	// step 2: right part (after index)
	// slice[3:] -> [40, 50]
	right := slice[index+1:]

	// step 3: merge left + right
	// append([10, 20], [40, 50]...)
	// [10 20, 40, 50]
	result := append(left, right...)
	return result
}

// Important Notes:
/*
* Go does not provides "rich collection API's" like (python/js) on Purpose:
 * Keeps the language simple and predictable
 * forces you to understand memory + performance
 * Avoids hidden allocations and abstractions
*/

// But in the GO version 1.21 + there is an standard package `slices` we can easily import it and use it for examples

func playingWithSlicesUsingSlicesPackage() {
	var slice = []int{1, 2, 3, 4, 5}
	clone := slices.Clone(slice)

	var slice2 = []int{1, 2, 3, 4, 5}

	fmt.Println("originalSlice: ", slice)
	fmt.Println("cloneSlice: ", clone)

	isEqual := slices.Equal(slice, clone)
	fmt.Println("is Equal slice and clone: ", isEqual)

	isSliceEqual := slices.Equal(slice, slice2)
	fmt.Println("is Equal slice and slice2: ", isSliceEqual)

	slices.Reverse(slice)
	fmt.Println("reversedSlice: ", slice)

}

// Advanced slice operations
// slice filtering

func filterEven(numbers []int) []int {
	var result []int

	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result

}

// error handling

func safeAccess(slice []int, index int) (int, error) {

	if index < 0 || index > len(slice) {
		return 0, fmt.Errorf("Index out of bound. You are trying to access index %d but slice length is %d\n", index, len(slice))
	}
	return slice[index], nil

}
