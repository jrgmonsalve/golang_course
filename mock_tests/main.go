package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

var GetPersonByDni = func(dni string) (Person, error) {
	var person Person
	person = Person{}
	time.Sleep(5 * time.Second)
	return person, nil
}

type Employee struct {
	Id       int
	Position string
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

type FullTimeEmployee struct {
	Employee
	Person
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDni(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p

	return ftEmployee, nil
}
