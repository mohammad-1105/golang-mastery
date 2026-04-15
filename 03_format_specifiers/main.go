package main

import "fmt"

// GO format specifiers

type User struct {
	ID   int
	Name string
}

func main() {

	// 1. Most used
	var (
		name   string  = "Mohammad Aman"
		age    int     = 22
		price  float32 = 3.14324233
		active bool    = true
	)

	fmt.Printf("Name: %s\n", name)      // string
	fmt.Printf("Age: %d\n", age)        // integer
	fmt.Printf("Price: %.2f\n", price)  // float (2 decimal)
	fmt.Printf("Active: %t\n", active)  // boolean
	fmt.Printf("Any value: %v\n", name) // default

	// 2. Debugging (very useful)
	user := User{ID: 1, Name: "Aman"}

	fmt.Printf("User (default): %v\n", user)
	fmt.Printf("User (detailed): %+v\n", user)
	fmt.Printf("Type: %T\n", user)

	// 3. Real world logging style
	fmt.Printf("user_id=%d action=%s\n", user.ID, "login")

	// 4. Float control
	fmt.Printf("Raw: %f\n", price)
	fmt.Printf("Controlled: %.2f\n", price)

	// 5. Width and alignment (cli output)
	fmt.Printf("|%-10s|%d|\n", name, age)

	// 6. Flags
	fmt.Printf("%+d\n", age)  // show sign
	fmt.Printf("%06d\n", age) // zero padding
	fmt.Printf("%#x\n", 255)  // hex with prefix

	// 7. strings variants
	fmt.Printf("%s\n", "GO") // GO
	fmt.Printf("%q\n", "GO") // "GO"

	// 8. fmt.Sprintf() vs fmt.Printf()
	msg := fmt.Sprintf("user_id=%d", user.ID)
	fmt.Println(msg) //user_id=1

	// 9. Golden Rule
	fmt.Printf("Safe Default: %v\n", user)

}
