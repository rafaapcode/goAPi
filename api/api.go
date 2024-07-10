package apí

import (
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
