package dao

import (
	"gorm.io/gorm"
)

type CompileOrder struct {
	gorm.Model
	Image      string
	Version    string
	StatusCode int // 0表示执行中，1表示成功，-1表示失败
	Branch     string

	Operator   User `gorm:"foreignKey:OperatorID"`
	OperatorID int

	Project   Project `gorm:"foreignKey:ProjectID"`
	ProjectID int

	DeployOrders []DeployOrder
}
