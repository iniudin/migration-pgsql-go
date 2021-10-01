package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name   string
	Email  string
	Course []Course
}
