package dbc

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/sheason2019/spoved/libs/dao"
	"github.com/sheason2019/spoved/libs/env"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open(env.DataRoot+"/spoved.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败:%w", err))
	}

	DB = db

	AutoMigrate()
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&dao.CompileOrder{},
		&dao.DeployOrder{},
		&dao.User{},
		&dao.Project{},
	)
}
