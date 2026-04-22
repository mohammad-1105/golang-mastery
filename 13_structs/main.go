package main

import "fmt"

/*
 * A Struct is a collection of fields
 * We define complex data models, organize information, define app data structure in struct
 * Struct are closing things that Go has to what are called classes in other languages, but the key difference is struct don't have any form of inheritance
 * The Designers of the GO feel that inheritance cause more problems that is solves in real-world code.
 */

// creating struct types and values
type User struct {
	Id      int
	Name    string
	age     int
	score   float32
	isAdult bool
}

// creating user structs
func getUsers() []User {

	// initialize struct using key-value notation
	u1 := User{
		Id:      101,
		Name:    "Mohammad Aman",
		age:     24,
		score:   99.99,
		isAdult: true,
	}

	// struct with few fields only
	u2 := User{
		Id:   102,
		Name: "John Doe",
	}

	// Initialize struct with values only
	u3 := User{
		103,
		"Jane",
		23,
		98.88,
		true,
	}

	// This var notation will create struct where all the fields were zero value (only false for bool)
	var u4 User

	fmt.Printf("User1: %v\n", u1) // User1: {101 Mohammad Aman 24 99.99 true}
	fmt.Printf("User2: %v\n", u2) // User2: {102 John Doe 0 0 false}
	fmt.Printf("User3: %v\n", u3) // User3: {103 Jane 23 98.88 true}
	fmt.Printf("User4: %v\n", u4) // User4: {0  0 0 false}

	// Now set the value using dot .
	u4.Id = 104
	u4.Name = "Michael Jordan"
	u4.age = 34
	u4.score = 98.99
	u4.isAdult = true

	fmt.Printf("User4: %v\n", u4) // User4: {104 Michael Jordan 34 98.99 true}

	return []User{u1, u2, u3, u4}
}

// Anonymous struct without name
func anonStruct() {

	employee := struct {
		id     string
		name   string
		salary float32
	}{
		id:     "11101",
		name:   "Jon Bonder",
		salary: 100000.00,
	}

	fmt.Println(employee.id)
	fmt.Println(employee.name)
	fmt.Println(employee.salary)
	fmt.Println(employee)

}

// comparing structs

type Point struct {
	x int
	y int
}

func compareStruct() (bool, bool) {

	// normal struct
	point1 := Point{
		x: 10,
		y: 20,
	}

	// anonymous struct
	point2 := struct {
		x int
		y int
	}{}
	point2.x = 10
	point2.y = 5

	// struct without key
	point3 := Point{
		10,
		20,
	}

	return point1 == point2, point1 == point3

}

func main() {

	users := getUsers()

	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
	// OUTPUT:
	/*
		0: main.User{Id:101, Name:"Mohammad Aman", age:24, score:99.99, isAdult:true}
		1: main.User{Id:102, Name:"John Doe", age:0, score:0, isAdult:false}
		2: main.User{Id:103, Name:"Jane", age:23, score:98.88, isAdult:true}
		3: main.User{Id:104, Name:"Michael Jordan", age:34, score:98.99, isAdult:true}
	*/

	// anonymous struct
	anonStruct()

	// comparing struct
	a, b := compareStruct()
	fmt.Printf("point1 == point2: %t\n", a) // false
	fmt.Printf("point1 == point3: %t\n", b) // true

	// loop over the dot slice
	dots := getDots()

	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}

	

}
