package utils

import "github.com/gin-gonic/gin"

func GetProps[T any](c *gin.Context) *T {
	var v T
	if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
		c.BindQuery(&v)
	} else {
		c.BindJSON(&v)
	}

	return &v
}
