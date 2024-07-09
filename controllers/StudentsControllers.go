package studentController

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/goAPi/db"
)

func GetStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Pegando todos os estudantes")
}

func GetStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be returned", id)
	return c.String(http.StatusOK, student)
}

func CreateStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	db.AddStudent(&student)
	return c.String(http.StatusOK, "Criando estudantes")
}

func UpdateStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be updated", id)
	return c.String(http.StatusOK, student)
}

func DeleteStudents(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be deleted", id)
	return c.String(http.StatusOK, student)
}
