package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"index"`
	PasswordHash  string
	PasswordSalt  string
	Projects      []Project
	CompileOrders []CompileOrder
	DeployOrders  []DeployOrder
}
