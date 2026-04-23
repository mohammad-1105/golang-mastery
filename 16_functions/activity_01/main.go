package main

import "fmt"

// calculation of the working hours of employees

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type WeekDay int

const (
	Sunday WeekDay = iota // starts from zero
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day WeekDay, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total := 0

	for _, v := range d.WorkWeek {
		total += v
	}

	return total
}

func main() {

	d := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "Tony",
			LastName:  "Stark",
		},
		HourlyRate: 10,
	}

	d.LogHours(Monday, 10)
	d.LogHours(Tuesday, 11)

	fmt.Println("Hours worked on Monday : ", d.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday : ", d.WorkWeek[Tuesday])
	fmt.Println("Hours worked this week : ", d.HoursWorked())
}
