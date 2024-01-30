package main

import "fmt"

type Person struct {
	name string
}

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func newEmployee(id int, name string) *Employee {
	return &Employee{
		id:   id,
		name: name,
	}
}

func main() {
	e := Employee{}
	fmt.Println(e)
	e.id = 10
	e.name = "John"
	e.SetId(20)
	fmt.Println(e)

	e2 := Employee{
		id:   30,
		name: "Jane",
	}
	fmt.Println(e2)

	e3 := new(Employee)
	e3.id = 40
	e3.name = "Jim"

	fmt.Println(*e3)

	e4 := newEmployee(50, "Jack")
	fmt.Println(e4)
}
