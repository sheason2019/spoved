package router

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	account_controller "github.com/sheason2019/spoved/libs/controller/account"
	compile_controller "github.com/sheason2019/spoved/libs/controller/compile"
	deploy_controller "github.com/sheason2019/spoved/libs/controller/deploy"
	project_controller "github.com/sheason2019/spoved/libs/controller/project"
	"github.com/sheason2019/spoved/libs/middleware"
	proxy_service "github.com/sheason2019/spoved/libs/service/proxy"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.UserMiddleware)

	// api service
	apiService := r.Group("/api")
	{
		account_controller.BindController(apiService)
		project_controller.BindController(apiService)
		compile_controller.BindController(apiService)
		deploy_controller.BindController(apiService)
	}

	// 路由网关
	r.Any("/proxy", func(ctx *gin.Context) {})

	reg := regexp.MustCompile(`^/api`)

	// 前端反向代理
	r.NoRoute(func(ctx *gin.Context) {
		if reg.Match([]byte(ctx.Request.URL.Path)) {
			ctx.String(404, "api not found")
			return
		}

		host, err := proxy_service.FindProxyHost(ctx, "root", "spoved-fe")
		if err != nil {
			panic(err)
		}

		proxyUrl, _ := url.Parse(fmt.Sprintf("http://%s", host))
		rp := httputil.NewSingleHostReverseProxy(proxyUrl)
		rp.ServeHTTP(ctx.Writer, ctx.Request)
	})

	return r
}
