package middleware

import (
	"github.com/gin-gonic/gin"
)

func GetProps[T any](c *gin.Context) *T {
	var v T
	var err error

	if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
		// 使用body存储参数的情况下，读取存储在上下文中的Data并返回
		err = c.BindJSON(&v)
	} else {
		// 使用Query存储参数的情况下，直接使用gin的Bind方法绑定参数并使用
		err = c.BindQuery(&v)
	}

	if err != nil {
		panic(err)
	}

	return &v
}
