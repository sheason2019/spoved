package dbc

import (
	"fmt"

	"github.com/sheason2019/spoved/libs/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ins *gorm.DB

func initClient() {

	db, err := gorm.Open(sqlite.Open(env.DataRoot+"/spoved.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败:%w", err))
	}

	ins = db
}

// 懒汉单例避免重复创建
func GetClient() *gorm.DB {
	if ins == nil {
		initClient()
	}

	return ins
}
