package main

import (
	"fmt"
)

type Employee struct {
	Name   string
	Salary float64
	Address
	// CompanyAddress
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

// type CompanyAddress struct {
// 	Street     string
// 	City       string
// 	State      string
// 	PostalCode string
// }

func main() {

	employee := Employee{Name: "AAA"}
	employee.City = "Seoul"

	fmt.Println(employee)

}
