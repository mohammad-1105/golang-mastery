package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Panic and recover:
/*
* panic()  is Go's mechanisms for unrecoverable errors it stops normal execution and begins stack unwinding
* recover() is used only inside deferred functions to intercept a panic and resumes execution safely
 */

// Example 1:  panic and recover
func main() {

	safeCall()
	fmt.Println("programs continues")

	http.Handle("/", recoveryMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8000", nil)

}

func safeCall() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from: ", r)
		}
	}()
	danger()
	fmt.Println("This will not run")
}

func danger() {

	panic(errors.New("program crashed in Danger"))
}

// NOTE from Example: 1
/**	What happen actually
* when panic is called:
	* current function stops immediately
	* Deferred functions run (LIFO) order
	* stack unwinds upward
	* if no recover(), runtime prints stack + exits
*/

// Example: 2 (Real world) : Http server recovery middleware
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v\n", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	panic("Database connection lost") // simulate error
}
