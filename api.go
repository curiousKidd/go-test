package main

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	router *gin.Engine
}

func NewAPIServer(empRepo *EmployeeRepository) *APIServer {
	return &APIServer{
		router: newRouter(empRepo),
	}
}

func (s *APIServer) Run() error {
	return s.router.Run()
}

func newRouter(empRepo *EmployeeRepository) *gin.Engine {
	r := gin.Default()

	empHandler := NewEmployeeHandler(empRepo)

	r.GET("/employees", empHandler.HandleEmployees())
	r.GET("/employees/:id", empHandler.HandleEmployeeById())

	return r
}
