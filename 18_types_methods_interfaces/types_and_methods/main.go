package main

import (
	"fmt"
	"time"
)

// types and methods in GO
// * GO is Statically Typed language with both built-in types and user-defined types. Like most another language. Go allows us to attach methods to types

// Example of types in GO

type Score int
type HighScore int
type Name string
type boxes []string
type TeamScores map[string]Score
type Converter func(num int) Score

// Methods : Like most modern languages. Go supports methods on User-defined Types.
// The methods for the types defined at the package block level

// Example
type Person struct {
	FirstName string
	LastName  string
	Age       int
	salary    float32
}

// methods for the Person struct

// this one is a value receiver
func (p Person) String() string {
	return fmt.Sprintf("FirstName: %s\nLastName: %s\nAge: %d\nSalary: %.2f\n", p.FirstName, p.LastName, p.Age, p.salary)
}

// pointer receiver (use this only when want to modify struct)
func (p *Person) updateAge(age int) {
	p.Age = age
}

// Another example :
type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d\nLastUpdatedAt: %v\n", c.total, c.lastUpdated)
}

// * GO allows calling method on nil pointers
type IntTree struct {
	val         int
	left, right *IntTree // pointers to others IntTree nodes
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}

	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// Embedding is not inheritance
type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner %d\n", val)
}

func (i Inner) Double() int {
	return i.A * 2
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	person1 := Person{
		FirstName: "Mohammad",
		LastName:  "Aman",
		Age:       25,
		salary:    120000.00,
	}

	fmt.Println(person1.String())

	person1.updateAge(22)
	fmt.Println(person1.Age) // 22

	c := Counter{
		total:       10,
		lastUpdated: time.Now(),
	}

	c.increment()
	c.increment()
	c.increment()
	c.increment()

	fmt.Println(c.String())
	// Total: 14
	// LastUpdatedAt: 2026-05-02 07:55:37.708071852 +0545 +0545 m=+0.000026363

	var tree *IntTree

	tree = tree.Insert(10)
	tree = tree.Insert(5)
	fmt.Println(tree.Contains(10)) // true

	root := &IntTree{val: 10}
	root.left = &IntTree{val: 12}
	root.right = &IntTree{val: 5}

	fmt.Println(root, root.left, root.right)

	// Type Declaration != inheritance
	var score1 Score = 100
	var score2 HighScore = 200
	// fmt.Println(score1 == score2) // Error

	fmt.Println(score1 == Score(score2)) // explicit conversion

	o := Outer{
		Inner: Inner{
			A: 10,
		},

		S: "Hello",
	}

	fmt.Println(o.Double())

}
