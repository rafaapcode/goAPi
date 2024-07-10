package apí

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rafaapcode/goAPi/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (api *API) GetStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Nenhum estudante")
	}

	listOfStudents := map[string][]schemas.StudentResponse{
		"students": schemas.NewReponse(students),
	}

	return c.JSON(http.StatusOK, listOfStudents)
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
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validating struct")
		return c.String(http.StatusBadRequest, "Error validating student")
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		Email:  studentReq.Email,
		CPF:    studentReq.CPF,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(&student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
	return c.String(http.StatusOK, "Criando estudantes")
}

func (api *API) UpdateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	receivedStudent := schemas.Student{}
	if err = c.Bind(&receivedStudent); err != nil {
		return err
	}

	updateStudent, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, fmt.Sprintf("Student with id %d, not found.", id))
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to found student to update")
	}

	student := updateStudentInfo(receivedStudent, updateStudent)

	if err := api.DB.UpdateStudent(&student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) DeleteStudents(c echo.Context) error {
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

	if err := api.DB.DeleteStudent(&student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(receivedStudent, updateStudent schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		updateStudent.Name = receivedStudent.Name
	}
	if receivedStudent.Email != "" {
		updateStudent.Email = receivedStudent.Email
	}
	if receivedStudent.CPF > 0 {
		updateStudent.CPF = receivedStudent.CPF
	}
	if receivedStudent.Age > 0 {
		updateStudent.Age = receivedStudent.Age
	}
	if receivedStudent.Active != updateStudent.Active {
		updateStudent.Active = receivedStudent.Active
	}

	return updateStudent
}
