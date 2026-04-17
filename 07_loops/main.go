package main

import "fmt"

/*
Go has only one loop construct: `for` but it covers all the cases like traditional loops, while loops, and iteration with `range`
*/
func main() {

	// 1. Traditional Loop
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	// print table with traditional loop
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", 2, i, 2*i)
	}

	// iterating over strings

	// this one fails (only prints the ASCII)
	name := "Mohammad Aman"
	for i := 0; i < len(name); i++ {
		fmt.Println(string(name[i]))
	}

	// ch (character) is rune here (range handle UTF-8 very well)
	for _, ch := range name {
		fmt.Println(string(ch))
	}

	// still demonstrate with for loop with runes
	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
	}

	// 2. while loop style
	j := 0
	for j < 5 {
		fmt.Println("J:", j)
		j++
	}

	// 3. Infinity Loop
	k := 0
	for {
		if k == 5 {
			break
		}
		fmt.Println("K: ", k)
		k++
	}

	// 4. Range loop (arrays/slices)
	arr := [5]int{10, 20, 30, 40, 50}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// ignore index
	for _, v := range arr {
		fmt.Println("Value: ", v)
	}

	// ! Important Notes Here:
	/*
		`range` returns (index, value) not (value, index) like some languages returns like javascript

		_ is not optional syntax sugar , it's required when we want to ignore a value
	*/

}
