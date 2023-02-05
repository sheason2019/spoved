package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context) {
	defer func() {
		e := recover()
		if e != nil {
			err := e.(error)
			fmt.Printf("[PANIC] %+v", err)
			c.JSON(500, gin.H{"message": err.Error(), "code": -1})
			return
		}
	}()
	c.Next()
}
