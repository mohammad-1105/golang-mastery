package main

import "fmt"

// * Go has rich feature support for functions. We can define our own function types in GO

// examples

type message func() // this is message func have to params and return types

type calc func(int, int) string // a calc function type having two int params and string return type

// creating a calculator using calc type
func add(x, y int) string {
	result := x + y
	return fmt.Sprintf("Added %d + %d = %d\n", x, y, result)
}

func sub(x, y int) string {
	result := x - y
	return fmt.Sprintf("Sub %d - %d = %d\n", x, y, result)
}
func mul(x, y int) string {
	result := x * y
	return fmt.Sprintf("Multiply %d * %d = %d\n", x, y, result)
}
func div(x, y int) string {
	result := float64(x / y)
	return fmt.Sprintf("Divided %d / %d = %.2f\n", x, y, result)
}

// calculator function
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}

// square function
func square(x int) func() int {
	f := func() int {
		return x * x
	}
	return f
}

// creating various functions to calculate salary
func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, workedHours int) int {
	return hourlyRate * workedHours
}

func main() {

	calculator(add, 10, 20)
	calculator(sub, 20, 10)
	calculator(mul, 2, 8)
	calculator(div, 60, 10)

	v := square(9)
	fmt.Println(v())
	fmt.Printf("Type of v : %T\n", v)

	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)
}
