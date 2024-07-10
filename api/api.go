package apí

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rafaapcode/goAPi/db"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDb := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDb,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.GetStudents)
	api.Echo.POST("/students", api.CreateStudents)
	api.Echo.PUT("/students/:id", api.UpdateStudent)
	api.Echo.GET("/students/:id", api.GetStudent)
	api.Echo.DELETE("/students/:id", api.DeleteStudents)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) GetStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Nenhum estudante")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) GetStudent(c echo.Context) error {
	var id string = c.Param("id")
	student := fmt.Sprintf("Student %s will be returned", id)
	return c.String(http.StatusOK, student)
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
