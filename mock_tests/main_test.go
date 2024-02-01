package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunck        func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "123456789",
			mockFunck: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDni = func(dni string) (Person, error) {
					return Person{
						DNI:  "123456789",
						Name: "John",
						Age:  30,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
				Person: Person{
					DNI:  "123456789",
					Name: "John",
					Age:  30,
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDni
	for _, test := range table {
		test.mockFunck()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Error(err)
		}
		if ft != test.expectedEmployee {
			t.Errorf("GetFullTimeEmployeeById(%d, %s) = %v; want %v", test.id, test.dni, ft, test.expectedEmployee)
		}
		GetEmployeeById = originalGetEmployeeById
		GetPersonByDni = originalGetPersonByDni
	}
}
