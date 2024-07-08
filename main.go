package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/students", getStudents)

	e.Logger.Fatal(e.Start(":8080"))
}

func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Pegando todos os estudantes")
}
