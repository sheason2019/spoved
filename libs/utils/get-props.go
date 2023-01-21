package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/libs/middleware"
)

func GetProps[T any](c *gin.Context) *T {
	var v T
	j, e := middleware.GetData(c)
	if e != nil {
		e.Panic()
	}

	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		panic(err)
	}

	return &v
}
