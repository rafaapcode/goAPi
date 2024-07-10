package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() *gorm.DB {
	var dsn string = "root:Qweasd#2003@tcp(api-golang.cloqmyckev6q.us-east-1.rds.amazonaws.com:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student *Student) error {
	if res := s.DB.Create(student); res.Error != nil {
		return res.Error
	}

	fmt.Print("Student created")
	return nil
}

func (s *StudentHandler) GetStudents() (students []Student, err error) {
	students = []Student{}
	err = s.DB.Find(&students).Error

	return students, err
}
