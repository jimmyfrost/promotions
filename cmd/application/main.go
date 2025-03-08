package main

import (
	"fmt"
	"net/http"
	"promotions/internal/models"
	promoDBService "promotions/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	promoDBService.InitDatabase()
	defer promoDBService.DB.Close()

	r := gin.Default()
	r.LoadHTMLGlob("../../views/*")

	r.GET("/employee", GetEmployees)
	r.POST("/employee", CreateEmployee)
	// r.PUT("/employee", GetEmployees)
	r.DELETE("/employee/:id", DeleteEmployee)

	r.Run(":8080")
}

func GetEmployees(c *gin.Context) {
	employees := promoDBService.ReadEmployeesList()
	c.HTML(http.StatusOK, "base", gin.H{
		"employees": employees,
	})
}

func CreateEmployee(c *gin.Context) {
	var emp models.Employee

	emp.Shorthand = c.Request.FormValue("shorthand")
	emp.Name = c.Request.FormValue("name")
	emp.Title = c.Request.FormValue("title")

	_, err := promoDBService.CreateNewEmployee(emp.Shorthand, emp.Name, emp.Title)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employees := promoDBService.ReadEmployeesList()
	c.HTML(http.StatusOK, "display", gin.H{
		"employees": employees,
	})
}

func DeleteEmployee(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	err := promoDBService.DeleteEmployee(id)
	if err != nil {
		fmt.Print("Couldnt delete:", err)
	}
	c.HTML(http.StatusOK, "employee.html", gin.H{
		"deleted": id,
	})
}
