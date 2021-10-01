package entity

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	CourseID    uint
	DateDue     *time.Time
}
