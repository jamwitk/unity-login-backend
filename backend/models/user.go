package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model    `json:"gorm.Model"`
	Name          string    `json:"name"`
	Username      string    `json:"username,omitempty"`
	Password      string    `json:"password,omitempty"`
	Email         string    `json:"email,omitempty"`
	LastLoginTime time.Time `json:"lastLoginTime,omitempty"`
}
