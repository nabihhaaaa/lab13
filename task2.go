// You can edit this code!
// Click here and start typing.
package main

//task 2

import "fmt"

type emp struct {
	name     string
	salary   int
	position string
}
type company struct {
	company_name string
	empls        []emp
}

func main() {
	var comp company

	e1 := emp{"Asif", 80000, "Cybersecurity Analyst"}
	e2 := emp{"Umer", 40000, "Data Scientist"}
	e3 := emp{"Hamza", 50000, "Content Creator"}

	emplys := []emp{e1, e2, e3}

	comp.company_name = "FAST Techz"
	comp.empls = emplys

	fmt.Println("Company: ", comp.company_name)
	fmt.Print("Employees are: \n\n")

	fmt.Print("Employee1: \n")
	fmt.Println("Name: ", comp.empls[0].name)
	fmt.Println("Salary: ", comp.empls[0].salary)
	fmt.Println("Position: ", comp.empls[0].position)

	fmt.Print("\nEmployee2: \n")
	fmt.Println("Name: ", comp.empls[1].name)
	fmt.Println("Salary: ", comp.empls[1].salary)
	fmt.Println("Position: ", comp.empls[1].position)

	fmt.Print("\nEmployee3 \n")
	fmt.Println("Name: ", comp.empls[2].name)
	fmt.Println("Salary: ", comp.empls[2].salary)
	fmt.Println("Position: ", comp.empls[2].position)

}
