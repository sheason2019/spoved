package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/exceptions/exception"
)

func GetProps[T any](c *gin.Context) *T {
	var v T
	j, e := GetData(c)
	if e != nil {
		e.Panic()
	}

	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		exception.New(err).Panic()
	}

	return &v
}
