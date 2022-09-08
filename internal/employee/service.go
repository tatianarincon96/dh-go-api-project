package employee

import (
	"github.com/tatianarincon96/dh-go-api-project/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Employee, error)
	GetById(id int) (domain.Employee, error)
	Create(p domain.Employee) (domain.Employee, error)
	Update(p domain.Employee) (domain.Employee, error)
	DeleteByID(id int) error
	UpdateField(p domain.Employee) (domain.Employee, error)
}

type service struct {
	r Repository
}

// NewSerice crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

// GetAll devuelve todos los employeeos
func (s *service) GetAll() ([]domain.Employee, error) {
	return s.r.GetAll(), nil
}

// GetByID busca un employeeo por id
func (s *service) GetById(id int) (domain.Employee, error) {
	e, err := s.r.GetByID(id)
	if err != nil {
		return domain.Employee{}, err
	}
	return e, nil
}

// Create agrega un nuevo employeeo
func (s *service) Create(e domain.Employee) (domain.Employee, error) {
	employee, err := s.r.Create(e)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

// Update actualiza un nuevo employee
func (s *service) Update(e domain.Employee) (domain.Employee, error) {
	employee, err := s.r.Update(e)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

// DeleteByID elimina un employeeo por id
func (s *service) DeleteByID(id int) error {
	err := s.r.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}

// Patch actualiza un campo del employeeo
func (s *service) UpdateField(e domain.Employee) (domain.Employee, error) {
	employee, err := s.r.UpdateField(e.Id, e)
	if err != nil {
		return domain.Employee{}, err
	}
	return employee, nil
}

/*func GetActiveEmployees(list []domain.Employee, checkActive bool) any {
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
}*/
