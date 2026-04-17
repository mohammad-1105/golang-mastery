package main

import "fmt"

func main() {
	x := 2

	// Basic switch
	switch x {
	case 1:
		fmt.Println("one")
	case 2, 3, 4:
		fmt.Println("two, three, four")
	default:
		fmt.Println("Other")
	}

	// switch case without expression (if else like chain)
	y := 15

	switch {
	case y > 10:
		fmt.Println("Greater than 10")

	case y > 5:
		fmt.Println("Greater than 5")

	default:
		fmt.Println("Small number")
	}

	// fallthrough example
	n := 1

	switch n {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2 (forced)")
	}

	// In the above code `fallthrough` is used to force the execution to continue to the next case. For example in our code we already fulfilled case 1 but still case 2 ran because of `fallthrough` forced execution

}
