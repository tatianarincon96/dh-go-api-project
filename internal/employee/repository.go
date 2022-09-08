package employee

import (
	"errors"
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
)

type Repository interface {
	GetAll() []domain.Employee
	GetByID(id int) (domain.Employee, error)
	Create(p domain.Employee) (domain.Employee, error)
	Update(p domain.Employee) (domain.Employee, error)
	DeleteByID(id int) error
	UpdateField(id int, p domain.Employee) (domain.Employee, error)
}

type repository struct {
	list []domain.Employee
}

// NewRepository crea un nuevo repositorio
func NewRepository(list []domain.Employee) Repository {
	return &repository{list}
}

// GetAll devuelve todos los productos
func (r *repository) GetAll() []domain.Employee {
	return r.list
}

// GetByID busca un producto por su id
func (r *repository) GetByID(id int) (domain.Employee, error) {
	for _, employee := range r.list {
		if employee.Id == id {
			return employee, nil
		}
	}
	return domain.Employee{}, errors.New("employee not found")
}

// Update actualiza un employeeo
func (r *repository) Update(e domain.Employee) (domain.Employee, error) {
	originalEmployee, err := r.GetByID(e.Id)
	if err != nil {
		return e, err
	}
	index, err := r.getIndex(e.Id)
	if err != nil {
		return domain.Employee{}, err
	}
	r.list[index] = setEmployee(e, originalEmployee)
	return r.list[index], nil
}

// Create crea un nuevo employeeo
func (r *repository) Create(e domain.Employee) (domain.Employee, error) {
	e.Id = len(r.list) + 1
	r.list = append(r.list, e)
	return e, nil
}

func (r *repository) DeleteByID(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	index, err := r.getIndex(id)
	if err != nil {
		return err
	}
	r.list = removeItem(r.list, index)
	return nil
}

// UpdateField actualiza un campo de un employeeo
func (r *repository) UpdateField(id int, e domain.Employee) (domain.Employee, error) {
	originalEmployee, err := r.GetByID(id)
	if err != nil {
		return e, err
	}
	index, err := r.getIndex(e.Id)
	if err != nil {
		return domain.Employee{}, err
	}
	r.list[index] = setAttributeInEmployee(e, originalEmployee)
	return r.list[index], nil
}

func setEmployee(new domain.Employee, original domain.Employee) domain.Employee {
	original.Name = new.Name
	original.Active = new.Active
	return original
}

func setAttributeInEmployee(new domain.Employee, original domain.Employee) domain.Employee {
	if new.Name != "" {
		original.Name = new.Name
	}
	if new.Active != original.Active {
		original.Active = new.Active
	}
	return original
}

func (r *repository) getIndex(id int) (int, error) {
	for i, employee := range r.list {
		if employee.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("no such employee")
}

func removeItem(s []domain.Employee, i int) []domain.Employee {

	return append(s[:i], s[i+1:]...)

}
