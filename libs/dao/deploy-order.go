package dao

import (
	"gorm.io/gorm"
)

type DeployOrder struct {
	gorm.Model
	// 部署时使用的镜像名称
	Image      string
	StatusCode int

	Operator   User `gorm:"foreignKey:OperatorID"`
	OperatorID int

	CompileOrder   CompileOrder `gorm:"foreignKey:CompileOrderID"`
	CompileOrderID int
}
