package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func AddStudent(student *Student) {
	db := Init()

	if res := db.Create(student); res.Error != nil {
		log.Fatal(res.Error)
	}

	fmt.Print("Student created")

}
