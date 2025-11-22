package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Printf("name =%v,age=%v,id=%v\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	e := Employee{Person{
		"Jack",
		3,
	}, "1"}
	e.PrintInfo()

}
