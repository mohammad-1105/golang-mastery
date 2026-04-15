package main

import "fmt"

// Data Types in Go

/** Go has several built-in data types that can be categorized into the following groups:

1. Basic Types:
	- bool: Represents a boolean value (true or false).
	- string: Represents a sequence of characters.
	- int, int8, int16, int32, int64: Represents signed integers of various sizes.
	- uint, uint8, uint16, uint32, uint64: Represents unsigned integers of various sizes.
	- float32, float64: Represents floating-point numbers.
	- complex64, complex128: Represents complex numbers.
	- byte: An alias for uint8, represents a byte of data.
	- rune: An alias for int32, represents a Unicode code point.

2. Aggregate Types:
	- array: A fixed-size sequence of elements of the same type.
	- struct: A collection of fields that can be of different types.

3. Reference Types:
	- slice: A dynamically-sized, flexible view into the elements of an array.
	- map: A collection of key-value pairs where each key is unique.
	- channel: A conduit through which goroutines can communicate with each other.

4. Interface Types:
	- interface: A type that specifies a set of method signatures but does not implement them.

5. Function Types:
	- func: Represents a function with a specific signature.

6. Pointer Types:
	- *T: Represents a pointer to a value of type T.

7. Unsafe Pointer Types:
	- unsafe.Pointer: A special pointer type that can hold any pointer value and is used for low-level programming.

Each of these data types serves a specific purpose and can be used to create variables, constants, and data structures in Go. Understanding these data types is essential for writing efficient and effective Go code.

*/

func main() {

	// Basic data types

	// Boolean
	var isActive bool = true

	// integers (signed and unsigned)
	var a int = 10
	var b int = -10
	var c uint = 11

	// Floating Point
	var pi float64 = 3.14

	// Complex numbers
	var comp complex64 = 1 + 2i

	// string
	var name string = "Mohammad Aman"

	println(isActive, a, b, c, pi, comp, name)

	runeDataType()
}

func runeDataType() {
	// Special types runes:
	// Rune is an alias for int32 and represents Unicode Code Point. It's used when we want to work with characters (not bytes), especially non-ASCII texts.

	// ! Always remember: rune use single quote ''
	var ch rune = 'A'

	var unicodeCharRune rune = '😊'

	fmt.Println(ch)         // 65 (Unicode code point)
	fmt.Println(string(ch)) // "A"

	fmt.Println(unicodeCharRune)         // Unicode code point value
	fmt.Println(string(unicodeCharRune)) // "😊"

}
