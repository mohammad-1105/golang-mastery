package main

import (
	"errors"
	"fmt"
)

// Creating an application to calculate pay for the week

// creating errors variables
var (
	ErrHourlyRate  = errors.New("Invalid hourly rate")
	ErrHoursWorked = errors.New("Invalid hours worked per week")
)

// payday function
func payDay(hoursWorked, hourlyRate int) int {

	report := func() {
		fmt.Printf("HoursWorked: %d\nHourlyRate: %d\n", hoursWorked, hourlyRate)
	}

	defer report()

	if hourlyRate < 10 || hourlyRate < 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime

	}
	return hoursWorked * hourlyRate
}

func main() {

	pay := payDay(81, 50)
	fmt.Println(pay)

}
