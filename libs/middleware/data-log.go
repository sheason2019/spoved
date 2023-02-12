package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		panic(e)
	}

	var buf bytes.Buffer
	_ = json.Indent(&buf, j, "", "  ")
	return buf.String()
}

// 展示返回值的内容
func ShowBodyString(body []byte) string {
	if len(body) == 0 {
		return "<nil>"
	}

	var buf bytes.Buffer
	err := json.Indent(&buf, body, "", "  ")
	if err != nil {
		return string(body)
	}

	return buf.String()
}

// 获取请求的参数
func GetData(c *gin.Context) (string, error) {
	j, exist := c.Get("data")
	if !exist {
		return "", errors.WithStack(errors.New("Context中不存在Data"))
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
