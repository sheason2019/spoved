package router

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
	compile_controller "github.com/sheason2019/spoved/libs/controller/compile"
	project_controller "github.com/sheason2019/spoved/libs/controller/project"
	"github.com/sheason2019/spoved/libs/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.Recovery)
	r.Use(middleware.DataLog)
	r.Use(middleware.UserMiddleware)

	// api service
	apiService := r.Group("/api")
	{
		account_controller.BindController(apiService)
		project_controller.BindController(apiService)
		compile_controller.BindController(apiService)
	}

	// 路由网关
	r.GET("/proxy", func(ctx *gin.Context) {})

	// 前端反向代理
	target := "http://root--spoved-fe-service"
	proxyUrl, _ := url.Parse(target)
	rp := httputil.NewSingleHostReverseProxy(proxyUrl)
	r.GET("/", func(ctx *gin.Context) {
		rp.ServeHTTP(ctx.Writer, ctx.Request)
	})

	return r
}
