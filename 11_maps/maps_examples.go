package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUserById(id string) (string, bool) {

	// get the users from getUsers()
	users := getUsers()

	user, ok := users[id]

	if !ok {
		fmt.Printf("User not found with ID : {%v}\n", id)
	}

	return user, ok

}

func getArgumentsFromTerminalAndPrintUser() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userId := os.Args[1]

	name, ok := getUserById(userId)

	if !ok {
		fmt.Printf("Passed User id : {%v} Not Found. \nUsers:\n", userId)

		for k, v := range getUsers() {
			fmt.Println("Id: ", k, "Name: ", v)
		}
		os.Exit(1)
	}

	fmt.Println("Name: ", name)

}

// slicing the week
func slicingTheWeek() []string {

	week := []string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}

	week = append(week[6:], week[:6]...)

	return week
}

// deleting elements from map

func deleteUser(id string) map[string]string {
	var users = map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	delete(users, id)

	return users
}
