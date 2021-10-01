package entity

import "gorm.io/gorm"

type Classroom struct {
	gorm.Model
	Name        string
	Description string
	CourseID    uint
}
