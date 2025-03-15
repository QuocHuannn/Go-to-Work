package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `gorm:"uniqueIndex;size:255" json:"email"`
	Password  string    `gorm:"size:255" json:"-"`
	FullName  string    `gorm:"size:255" json:"full_name"`
	LastLogin time.Time `json:"last_login"`
	Status    int       `gorm:"default:1" json:"status"` // 1: active, 0: inactive
}
