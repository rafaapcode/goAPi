package db

import (
	"github.com/rafaapcode/goAPi/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	var dsn string = "root:Qweasd#2003@tcp(api-golang.cloqmyckev6q.us-east-1.rds.amazonaws.com:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize MYSQL: %s", err.Error())
	}

	db.AutoMigrate(&schemas.Student{})
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student *schemas.Student) error {
	if res := s.DB.Create(student); res.Error != nil {
		log.Error().Msg("Failed to create students")
		return res.Error
	}

	log.Info().Msg("Create student!")
	return nil
}

func (s *StudentHandler) GetStudents() (students []schemas.Student, err error) {
	students = []schemas.Student{}
	err = s.DB.Find(&students).Error

	return
}

func (s *StudentHandler) GetStudent(id int) (student schemas.Student, err error) {
	student = schemas.Student{}
	err = s.DB.First(&student, id).Error

	return
}

func (s *StudentHandler) UpdateStudent(newStudent *schemas.Student) error {
	return s.DB.Save(newStudent).Error
}

func (s *StudentHandler) DeleteStudent(student *schemas.Student) error {
	return s.DB.Delete(student).Error
}
