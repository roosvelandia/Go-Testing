package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CFO",
					}, nil
				}
				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						Name: "Jhon Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "Jhon Doe",
				},
				Employee: Employee{
					Id:       1,
					Position: "CFO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error While getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error age expected %d got %d", test.expectedEmployee.Age, ft.Age)
		}

		if ft.Position != test.expectedEmployee.Position {
			t.Errorf("Error position expected %s got %s", test.expectedEmployee.Position, ft.Position)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}
