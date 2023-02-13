package dao

import (
	"gorm.io/gorm"
)

type DeployOrder struct {
	gorm.Model
	Image      string
	StatusCode int

	Operator   User `gorm:"foreignKey:OperatorID"`
	OperatorID int

	CompileOrder   CompileOrder `gorm:"foreignKey:CompileOrderID"`
	CompileOrderID int
}
