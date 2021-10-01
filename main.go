package main

import (
	"classroom-example/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=172.17.0.2 user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(
		&entity.Student{},
		&entity.Teacher{},
		&entity.Classroom{},
		&entity.Course{},
		&entity.Task{},
	)

}
