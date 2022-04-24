package main

import (
	"oop/employee"
)

func main() {
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Smith",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	// var e employee.Employee
	e := employee.New("Bob", "Smith", 30, 20)
	e.LeavesRemaining()
}
