package employee

import (
	"fmt"
	"github.com/tatianarincon96/dh-go-api-project/cmd/server/internal/domain"
)

func GetActiveEmployees(list []domain.Employee, checkActive bool) any {
	var activeEmployees []domain.Employee
	var noActiveEmployees []domain.Employee
	for _, e := range list {
		if e.Active {
			activeEmployees = append(activeEmployees, e)
		} else {
			noActiveEmployees = append(noActiveEmployees, e)
		}
	}
	if checkActive {
		return activeEmployees
	} else {
		return noActiveEmployees
	}
}

func AddEmployee(list []domain.Employee, e domain.Employee) []domain.Employee {
	list = append(list, e)
	return list
}

func GetEmployeeById(list []domain.Employee, id int) (domain.Employee, error) {
	for _, employee := range list {
		if employee.Id == id {
			return employee, nil
		}
	}
	return domain.Employee{}, fmt.Errorf("Employee not found")
}
