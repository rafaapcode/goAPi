package apí

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/goAPi/db"
	"gorm.io/gorm"
)

func (api *API) GetStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Nenhum estudante")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) GetStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, fmt.Sprintf("Student with id %d, not found.", id))
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) CreateStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(&student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
	return c.String(http.StatusOK, "Criando estudantes")
}

func (api *API) UpdateStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be updated", id)
	return c.String(http.StatusOK, student)
}

func (api *API) DeleteStudents(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be deleted", id)
	return c.String(http.StatusOK, student)
}
