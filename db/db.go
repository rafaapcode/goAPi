package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
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

func AddStudent() {
	db := Init()

	studante := Student{
		Name:   "Rafael",
		CPF:    123456788,
		Email:  "Rafa@gmail.com",
		Age:    20,
		Active: true,
	}

	if res := db.Create(&studante); res.Error != nil {
		log.Fatal(res.Error)
	}

	fmt.Print("Student created")

}
