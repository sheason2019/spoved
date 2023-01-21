package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sheason2019/spoved/exceptions/exception"
)

// 展示请求参数
func DataLog(c *gin.Context) {
	arg := make(map[string]any)
	var j string
	if c.Request.Method == "GET" || c.Request.Method == "DELETE" {
		values := c.Request.URL.Query()
		for key := range values {
			arg[key] = c.Request.URL.Query().Get(key)
		}
		j = JsonStr(arg)
		fmt.Printf("REQUEST QUERY\n%s\n", j)
	} else {
		c.BindJSON(&arg)
		j = JsonStr(arg)
		fmt.Printf("REQUEST BODY\n%s\n", j)
	}

	c.Set("data", j)
	c.Next()
}

func JsonStr(v any) string {
	j, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}

	var buf bytes.Buffer
	_ = json.Indent(&buf, j, "", "  ")
	return buf.String()
}

func GetData(c *gin.Context) (string, *exception.Exception) {
	j, exist := c.Get("data")
	if !exist {
		return "", exception.New(errors.New("Context中不存在Data"))
	}

	return j.(string), nil
}
