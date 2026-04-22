package main

import (
	"fmt"
	"strings"
)

// GO FUNCTIONS (quick notes)

// Purpose:
// - Break complex logic
// - Avoid duplication (DRY)
// - Improve reuse & maintainability

// Features:
// - Multiple return values (unlike most languages)
// - First-class → assign, pass, return functions
// - Supports closures & anonymous functions

// Design rules:
// - Single responsibility (1 task)
// - Keep small (≈25 lines guideline)

// Syntax:
// func name(params) returnTypes {
// logic
// }

// Key parts:
// - func → keyword
// - name → camelCase (lower=private, Upper=public)
// - params → inputs (optional)
// - returns → outputs (optional, can be multiple)
// - body → logic + return

// Notes:
// - Shorthand params: (a, b int)
// - Signature = params + return types (keep stable)
// - Multiple return statements common (esp. errors)

// check num function (if the number is even or odd)

func checkNum() {
	for i := 1; i <= 30; i++ {
		if i%2 == 0 {
			fmt.Println("EVEN")
		} else {
			fmt.Println("ODD")
		}
	}
}

// another way
func checkNumbers(i int) (int, string) {
	switch {
	case i%2 == 0:
		return i, "Even"
	}
	return i, "Odd"
}

// function that print salesman expectation ratings from the number of items sold
func itemsSold() {
	items := make(map[string]int)

	items["john"] = 41
	items["jane"] = 107
	items["noah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations")
		} else if v > 100 {
			fmt.Println("exceeded expectations")
		}
	}
}

// example:
// Mapping index values to column headers
// function that takes the slice of column headers from a CSV file and it will print out the map of an index value of the header we are interested in

func csvHeaderColumn(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)

	for i, v := range header {
		v = strings.TrimSpace(v)

		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v

		}
	}

	fmt.Println(csvHeadersToColumnIndex)
}

func main() {
	checkNum()
	itemsSold()

	header1 := []string{"empId", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHeaderColumn(header1)
	header2 := []string{"employee", "empId", "hours worked", "address", "manager", "hourly rate"}
	csvHeaderColumn(header2)

	for i := 1; i <= 30; i++ {
		num, result := checkNumbers(i)
		fmt.Printf("Results: %d %s\n", num, result)
	}
}
