package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	studentController "github.com/rafaapcode/goAPi/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/students", studentController.GetStudents)
	e.POST("/students", studentController.CreateStudents)
	e.PUT("/students/:id", studentController.UpdateStudent)
	e.GET("/students/:id", studentController.GetStudent)
	e.DELETE("/students/:id", studentController.DeleteStudents)

	e.Logger.Fatal(e.Start(":8080"))
}
