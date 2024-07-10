package apí

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rafaapcode/goAPi/db"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", GetStudents)
	api.Echo.POST("/students", CreateStudents)
	api.Echo.PUT("/students/:id", UpdateStudent)
	api.Echo.GET("/students/:id", GetStudent)
	api.Echo.DELETE("/students/:id", DeleteStudents)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func GetStudents(c echo.Context) error {
	students, err := db.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Nenhum estudante")
	}
	return c.JSON(http.StatusOK, students)
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
	if err := db.AddStudent(&student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}
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
