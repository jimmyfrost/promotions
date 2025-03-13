package main

import (
	"fmt"
	"log"
	"net/http"
	"promotions/internal/models"
	promoDBService "promotions/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Action string `json:"action"`
	ID     string `json:"data"`
}

func main() {
	promoDBService.InitDatabase()
	defer promoDBService.DB.Close()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := gin.Default()
	r.LoadHTMLGlob("../../views/*")

	r.GET("/employees", employeeGetRouter)
	r.POST("/employees", CreateEmployee)
	// r.PUT("/employee", GetEmployees)
	r.DELETE("/employees/:id", DeleteEmployee)

	//r.GET("/achievement", GetAchievements)
	//r.POST("/achievement", CreateAchievement)
	// r.PUT("/employee", GetEmployees)
	//r.DELETE("/achievement/:id", DeleteAchievement)

	r.Run(":8080")
}

func employeeGetRouter(c *gin.Context) {
	log.Print("Entered Router")
	var payload RequestPayload

	payload.Action = c.Query("action")
	payload.ID = c.Query("id")

	fmt.Print()

	switch payload.Action {
	case "getAchievements":
		id, _ := strconv.ParseInt(payload.ID, 10, 64)
		GetAchievements(c, id)
	default:
		GetEmployees(c)
	}
}

func GetEmployees(c *gin.Context) {
	employees := promoDBService.ReadEmployeesList()
	c.HTML(http.StatusOK, "base", gin.H{
		"employees": employees,
	})
}

func CreateEmployee(c *gin.Context) {
	log.Print("Creating Employee")
	var emp models.Employee

	emp.Email = c.Request.FormValue("email")
	emp.Name = c.Request.FormValue("name")
	emp.Title = c.Request.FormValue("title")
	emp.Track = c.Request.FormValue("track")

	_, err := promoDBService.CreateNewEmployee(emp.Email, emp.Name, emp.Title, emp.Track)
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
	c.HTML(http.StatusOK, "employees.html", gin.H{
		"deleted": id,
	})
}

func GetAchievements(c *gin.Context, id int64) {
	log.Print("Employee Clicked")
	achievements := promoDBService.GetAchievementsByEmployeeID(id)
	c.HTML(http.StatusOK, "achievements", gin.H{
		"achievements": achievements,
	})
}

func CreateAchievement(c *gin.Context) {
	var achievement models.Achievement

	achievement.Situation = c.Request.FormValue("situation")
	achievement.Task = c.Request.FormValue("task")
	achievement.Action = c.Request.FormValue("action")
	achievement.Result = c.Request.FormValue("result")
	empVal, err := strconv.ParseInt(c.Request.FormValue("employeeId"), 10, 64)
	if err != nil {
		log.Fatal("Could not bind employeeId")
	}

	achievement.EmployeeID = empVal

	_, err2 := promoDBService.CreateAchievement(achievement.Situation, achievement.Task, achievement.Action, achievement.Result, achievement.EmployeeID)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	employees := promoDBService.ReadEmployeesList()
	c.HTML(http.StatusOK, "display", gin.H{
		"employees": employees,
	})
}

func DeleteAchievement(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	err := promoDBService.DeleteEmployee(id)
	if err != nil {
		fmt.Print("Couldnt delete:", err)
	}
	c.HTML(http.StatusOK, "employees.html", gin.H{
		"deleted": id,
	})
}
