package main

func main() {

	variablesAndConstants()
}

func variablesAndConstants() {
	// There are many ways to declare variables in GO, like

	// 1. Using the var keyword and specifying the type
	var a int = 1

	// 2. Using the var keyword and letting GO infer the type
	var b = "Hello, world"

	// 3. Using the short variable declaration syntax (:=) which also infers the type
	c := 3.14

	// 4. Declaring multiple variables at once in one line
	var d, e int = 2, 3

	// 5. Declaring multiple variables at once without specifying the type (GO will infer the type)
	var f, g = "Hello", "World"

	// 6. Declaring multiple variables at once
	var (
		firstName string = "John"
		lastName  string = "Doe"
		age       int    = 30
	)

	// 7. Declaring a constant using the const keyword
	const pi = 3.14

	// 8. Declaring a constant with a specific type
	const euler float64 = 2.71828

	// 9. Declaring multiple constants at once
	const (
		lightSpeed            = 299792458   // in meters per second
		gravitationalConstant = 6.67430e-11 // in m^3 kg^-1 s^-2
	)

	// printing the variables and constants to the console
	println("a:", a)
	println("b:", b)
	println("c:", c)
	println("d:", d, "e:", e)
	println("f:", f, "g:", g)
	println("firstName:", firstName, "lastName:", lastName, "age:", age)
	println("pi:", pi)
	println("euler:", euler)
	println("lightSpeed:", lightSpeed)
	println("gravitationalConstant:", gravitationalConstant)

}
