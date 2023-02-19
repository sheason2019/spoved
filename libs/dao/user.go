package dao

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `gorm:"index"`
	PasswordHash  string
	PasswordSalt  string
	Projects      []Project      `gorm:"foreignKey:CreatorID"`
	CompileOrders []CompileOrder `gorm:"foreignKey:OperatorID"`
	DeployOrders  []DeployOrder  `gorm:"foreignKey:OperatorID"`
}
