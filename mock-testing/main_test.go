package main

func TestGetFullTimeEmployeeById(t *testing.T){
	table := []struct{
		id int
		dni string
		mockFunc func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:1,
			dni: "1",
			mockFunc: func(){
				GetEmployeeById = func(id int)(Employee,error){
					return Employee {
						Id: 1,
						Position: "CFO",
					}, nil
				}
				GetPersonByDNI = func (id string) (Person, error){
					return Person{
						Name: "Jhon Doe",
						Age: 35,
						DNI: "1"
					}, nil
				}
			}
		}
	}
}