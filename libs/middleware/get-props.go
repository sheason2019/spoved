package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetProps[T any](c *gin.Context) *T {
	var v T

	// 使用body存储参数的情况下，读取存储在上下文中的Data并返回
	if c.Request.Method != "GET" && c.Request.Method != "DELETE" {
		j, err := GetData(c)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(j), &v)
		if err != nil {
			panic(err)
		}

		return &v
	} else {
		// 使用Query存储参数的情况下，直接使用gin的Bind方法绑定参数并使用
		err := c.BindQuery(&v)
		if err != nil {
			panic(err)
		}
		return &v
	}
}
