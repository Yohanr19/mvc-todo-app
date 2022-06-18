package psqlStore

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Text     string
	IsActive bool
}
