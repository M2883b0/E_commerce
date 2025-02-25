package api

import (
	"content_system/internal/services"
	"github.com/gin-gonic/gin"
)

// 设置项目的根路径，所有接口都加上api开头
const (
	rootPath = "/api"
)

func CmsRouter(r *gin.Engine) {
	//cms对象实例化
	cmsApp := services.NewCmsApp()
	//创建sessionID中间件实例
	sessionAuth := NewSessionAuth()

	////资源监控中间件，上报数据到prometheus
	//r.Use(prometheusMiddleware())
	////链路追踪中间件，上报数据到zipkin
	//r.Use(opentracingMiddleware())

	//创建【路由组/cms/】
	cmsGroup := r.Group(rootPath + "/cms").Use(sessionAuth.Auth) //使用Use()方法，为这个组的接口，加入鉴权Auth中间件
	//cmsGroup := r.Group(rootPath + "/cms").Use(JwtToken()) //使用jwt中间件
	{
		//路径/api/cms/ping
		cmsGroup.GET("/ping", cmsApp.Hello)
		//路径/api/cms/content/create
		cmsGroup.POST("/content/create", cmsApp.ContentCreate)
		//路径/api/cms/content/update
		cmsGroup.POST("/content/update", cmsApp.ContentUpdate)
		//路径/api/cms/content/delete
		cmsGroup.POST("/content/delete", cmsApp.ContentDelete)
		//路径/api/cms/content/find
		cmsGroup.POST("/content/find", cmsApp.ContentFind)

		////路径/api/cms/user/create
		//cmsGroup.POST("/user/create", cmsApp.UserCreate)
		//路径/api/cms/content/update
		cmsGroup.POST("/user/update", cmsApp.UserUpdate)
		//路径/api/cms/content/delete
		cmsGroup.POST("/user/delete", cmsApp.UserDelete)
		//路径/api/cms/content/find
		cmsGroup.POST("/user/find", cmsApp.UserFind)
	}
	//创建不需要鉴权的【路由组】例如注册、登录
	noAuthGroup := r.Group(rootPath + "/noauth")
	{
		//注册路径/api/noauth/register
		noAuthGroup.POST("/register", cmsApp.Register)
		//登录路径/api/noauth/register
		noAuthGroup.POST("/login", cmsApp.Login)
	}
	//资源监控Prometheus的配置文件yml，设置了8080。这里我们要给它提供一个接口
	// http://localhost:8080/metrics
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
