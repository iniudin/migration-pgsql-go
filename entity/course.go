package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string
	Description string
	Classroom   Classroom
	Task        []Task
	StudentID   uint
	TeacherID   uint
	TaskID      uint
}
