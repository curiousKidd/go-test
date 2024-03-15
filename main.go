package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

const dsn = "root:TEST_USER@tcp(HOST:PORT)/employees?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}

	employeeRepository := NewEmployeeRepository(db)

	if err = NewAPIServer(employeeRepository).Run(); err != nil {
		log.Println(err)
		return
	}
}
