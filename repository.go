package main

import "gorm.io/gorm"

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) FindAllEmployees(offset, limit int) ([]Employee, error) {
	var employees []Employee
	err := r.db.Offset(offset).Limit(limit).Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepository) FindEmployeeById(id int) (Employee, error) {
	employee := Employee{EmpNo: id}
	err := r.db.Find(&employee).Error
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}
