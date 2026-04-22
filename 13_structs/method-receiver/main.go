package main

import "fmt"

// * Method receiver in GO:
// * Use it when to define a method on a type (usually struct)

// examples

type User struct {
	Name string
}

// This method got tied to User struct type
func (u User) Greet() string {
	return "Hello " + u.Name
}

func (u User) isEmpty() bool {
	return u.Name == ""
}

// rename user name : we should pass the user reference
func (u *User) Rename(name string) string {
	u.Name = name
	return u.Name
}

func main() {

	// create struct instance and call directly Greet() method into it
	user := User{Name: "Mohammad Aman"}
	fmt.Println(user.Greet())

	// check if the name is empty
	fmt.Println(user.isEmpty()) // false

	// rename
	newName := user.Rename("John Doe")
	fmt.Println(newName)
}
