package dao

import (
	"gorm.io/gorm"
)

type CompileOrder struct {
	gorm.Model

	// 构建时使用的镜像名称
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

// 构建产物镜像的名称
func (order *CompileOrder) OutImageName() string {
	return order.Operator.Username + "/" + order.Project.ProjectName + ":" + order.Version
}
