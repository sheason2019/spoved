package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/exceptions/exception"
)

func Recovery(c *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		switch err.(type) {
		case *exception.Exception:
			handleException(c, err)
		case error:
			handleError(c, err)
		}
	}()
	c.Next()
}

func handleException(c *gin.Context, err any) {
	e := err.(*exception.Exception)
	if e.Code == -1 {
		c.JSON(500, gin.H{"message": "服务端逻辑异常", "code": -1})
	} else {
		c.JSON(500, gin.H{"message": e.Error(), "code": e.Code})
	}
	fmt.Println(e.Print())
	c.Abort()
}

func handleError(c *gin.Context, err any) {
	e := err.(error)
	c.JSON(500, gin.H{"message": "服务端未知错误", "code": -1})
	fmt.Println(e.Error())
	c.Abort()
}
