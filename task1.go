// You can edit this code!
// Click here and start typing.
package main

//task 1

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var p1 Person
	var p2 Person

	p1.name = "John"
	p1.age = 26
	p1.job = "Pilot"
	p1.salary = 60000

	p2.name = "Adam"
	p2.age = 34
	p2.job = "Photographer"
	p2.salary = 50000

	printS(p1)

}
func printS(P1 Person) {
	fmt.Println("Name: ", P1.name)
	fmt.Println("Age: ", P1.age)
	fmt.Println("Job: ", P1.job)
	fmt.Println("Salary: ", P1.salary)

}