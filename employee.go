package main

import "time"

type Employee struct {
	EmpNo     int       `json:"empNo" gorm:"primaryKey"`
	BirthDate time.Time `json:"birthDate"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	HireDate  time.Time `json:"hireDate"`
}
