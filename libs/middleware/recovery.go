package middleware

import (
	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context) {
	defer func() {
		err := recover().(error)
		if err == nil {
			return
		}
		c.JSON(500, gin.H{"message": err.Error(), "code": -1})
	}()
	c.Next()
}
