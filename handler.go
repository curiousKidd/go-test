package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type EmployeeHandler struct {
	repository *EmployeeRepository
}

func NewEmployeeHandler(repository *EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{repository: repository}
}

type Message struct {
	Error string `json:"error"`
}

type Pagination struct {
	Offset int `form:"offset" binding:"gte=0"`
	Limit  int `form:"limit,default=20" binding:"gte=0"`
}

func (h *EmployeeHandler) HandleEmployees() gin.HandlerFunc {
	return func(c *gin.Context) {
		var page Pagination
		err := c.ShouldBindQuery(&page)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Message{Error: err.Error()})
			return
		}

		employees, err := h.repository.FindAllEmployees(page.Offset, page.Limit)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, Message{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, employees)
	}
}

type EmployeeParam struct {
	EmpNo int `uri:"id" binding:"required,gte=1"`
}

func (h *EmployeeHandler) HandleEmployeeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var empParam EmployeeParam
		err := c.ShouldBindUri(&empParam)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, Message{Error: err.Error()})
			return
		}

		employees, err := h.repository.FindEmployeeById(empParam.EmpNo)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, Message{Error: err.Error()})
			return
		}
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, Message{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, employees)
	}
}
