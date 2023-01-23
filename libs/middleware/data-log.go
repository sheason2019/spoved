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
		fmt.Printf("\n\033[32mREQUEST QUERY\033[0m\n%s\n", j)
	} else {
		c.BindJSON(&arg)
		j = JsonStr(arg)
		fmt.Printf("\n\033[32mREQUEST BODY\033[0m\n%s\n", j)
	}

	c.Set("data", j)

	blw := CustomResponseWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw
	c.Next()

	j = ShowBodyString(blw.body.Bytes())
	fmt.Printf("\033[32mRESPONSE DATA\033[0m\n%s\n", j)
}

func JsonStr(v any) string {
	j, e := json.Marshal(v)
	if e != nil {
		exception.New(e).Panic()
	}

	var buf bytes.Buffer
	_ = json.Indent(&buf, j, "", "  ")
	return buf.String()
}

func ShowBodyString(body []byte) string {
	m := make(map[string]any)
	err := json.Unmarshal(body, &m)
	if err != nil {
		exception.New(err).Panic()
	}

	return JsonStr(m)
}

func GetData(c *gin.Context) (string, *exception.Exception) {
	j, exist := c.Get("data")
	if !exist {
		return "", exception.New(errors.New("Context中不存在Data"))
	}

	return j.(string), nil
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
