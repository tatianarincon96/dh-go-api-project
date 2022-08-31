package employees

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Active bool
}

func GetActiveEmployees(isActive bool) any {
	var activeEmployees []Employee
	var noActiveEmployees []Employee
	employees := GetEmployeeList()
	for _, e := range employees {
		if e.Active {
			activeEmployees = append(activeEmployees, e)
		} else {
			noActiveEmployees = append(noActiveEmployees, e)
		}
	}
	if isActive {
		return activeEmployees
	} else {
		return noActiveEmployees
	}
}

func AddEmployee(e Employee) []Employee {
	employees := GetEmployeeList()
	employees = append(employees, e)
	return employees
}

func GetEmployeeList() []Employee {
	return []Employee{
		{Id: 1, Name: "John", Active: true},
		{Id: 2, Name: "Mary", Active: true},
		{Id: 3, Name: "Mike", Active: false},
		{Id: 4, Name: "Adam", Active: true},
		{Id: 5, Name: "Peter", Active: false},
	}
}

func GetEmployeeById(id int) (Employee, error) {
	employees := GetEmployeeList()
	for _, employee := range employees {
		if employee.Id == id {
			return employee, nil
		}
	}
	return Employee{}, fmt.Errorf("employee not found")
}
