package router

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	proxy_service "github.com/sheason2019/spoved/libs/service/proxy"
)

// IngressRouter
func SetupProxy() *gin.Engine {
	r := gin.Default()

	// /api/ 请求反向代理到 Spoved
	r.Any("/api/*endpoint", func(ctx *gin.Context) {
		ctx.Request.URL.Path = ctx.Param("endpoint")
		proxyTo(ctx, "root", "spoved")
	})

	// /proxy/ 请求反向代理到指定 Project
	group := r.Group("/proxy/")
	{
		group.Any(":username/:projName/*endpoint", func(ctx *gin.Context) {
			ctx.Abort()
		})
		group.Use(func(ctx *gin.Context) {
			ctx.AbortWithStatus(404)
		})
	}

	// No Router 请求反向代理到 Spoved-fe
	r.NoRoute(func(ctx *gin.Context) {
		proxyTo(ctx, "root", "spoved-fe")
	})

	return r
}

// 反向代理到数据库中记录的指定服务
func proxyTo(ctx *gin.Context, username, projName string) {
	host, err := proxy_service.GetHost(ctx, username, projName)
	if err != nil {
		panic(err)
	}

	proxyUrl, _ := url.Parse(fmt.Sprintf("http://%s", host))
	rp := httputil.NewSingleHostReverseProxy(proxyUrl)
	rp.ServeHTTP(ctx.Writer, ctx.Request)
}
