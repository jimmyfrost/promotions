package main

import (
	//"fmt"
	"net/http"
	//"promotions/internal/models"
	promoDBService "promotions/services"

	"github.com/gin-gonic/gin"
)

func main() {
	promoDBService.InitDatabase()
	defer promoDBService.DB.Close()

	r := gin.Default()
	r.LoadHTMLGlob("../../templates/*")

	r.GET("/", func(c *gin.Context) {
		employees := promoDBService.ReadEmployeesList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"employees": employees,
		})
	})

	r.Run(":8080")
}

/*
func testEmployee() models.Employee {
	suggestion := models.Suggestion{Type: 0, Title: "Take Course", Details: "Use WebDev Go Course"}
	s := []models.Suggestion{}
	s = append(s, suggestion)

	goal := models.Goal{Title: "Learn Go", Details: "Nail down fundementals of GO. NO AI!", TimeHorizonInMonths: 3, Suggestions: s}
	g := []models.Goal{}
	g = append(g, goal)

	emp := models.Employee{Shorthand: "Jimmy", Name: "Jimmy Frost", Title: "EM", Goals: g}

	return emp
}
*/
