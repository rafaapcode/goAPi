package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rafaapcode/goAPi/db"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/students", getStudents)
	e.POST("/students", createStudents)
	e.PUT("/students/:id", updateStudent)
	e.GET("/students/:id", getStudent)
	e.DELETE("/students/:id", deleteStudents)

	e.Logger.Fatal(e.Start(":8080"))
}

func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Pegando todos os estudantes")
}

func getStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be returned", id)
	return c.String(http.StatusOK, student)
}

func createStudents(c echo.Context) error {
	db.AddStudent()
	return c.String(http.StatusOK, "Criando estudantes")
}

func updateStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be updated", id)
	return c.String(http.StatusOK, student)
}

func deleteStudents(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be deleted", id)
	return c.String(http.StatusOK, student)
}
