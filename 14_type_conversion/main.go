package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// * In GO: Type conversion is intentionally explicit and strict
// * GO doesn't implicit type conversion we do explicit type conversion using T(value) syntax
func main() {
	// Examples

	var a int = 10
	var b float32 = float32(a) // explicit type conversion

	fmt.Printf("Type of b: %T\n", b) // float32

	// no implicit conversion
	// var c float32 = a // compile error

	piFloat := 3.14
	piInt := int(piFloat)
	fmt.Println(piInt) // 3 (truncation, not rounding)

	// strings -> number conversion

	strNum := "123"
	// num := int(strNum) // Compile Error

	// better option use strconv package
	num, error := strconv.Atoi(strNum)

	if error != nil {
		fmt.Printf("Error in parsing numStr into Num: %#v", error.Error())
		os.Exit(1)
	}

	fmt.Println(num)

	fmt.Println(convert())

}

func convert() string {

	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	// convert smaller into a larger type: Always a safe operation
	m := fmt.Sprintf("int8 = %v -> int64 = %v\n", i8, int64(i8))

	// convert from int type to int8 : this cause integer overflow
	m += fmt.Sprintf("int = %v -> int8 = %v\n", i, int8(i))

	// convert int8 into float64: this doesn't cause an overflow but data is unchanged
	m += fmt.Sprintf("int8 = %v -> float64 = %v\n", i8, float64(i8))

	// convert float type into int type: All the decimal data will be lost
	m += fmt.Sprintf("float64 = %v -> int = %v\n", f64, int(f64))

	return m
}
