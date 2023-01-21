package middleware

import (
	"github.com/gin-gonic/gin"
)

var Recovery = gin.CustomRecovery(func(c *gin.Context, err any) {
	c.JSON(500, "服务端发生错误")
	c.Abort()
})
