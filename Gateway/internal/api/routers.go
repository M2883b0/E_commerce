package api

import (
	"Gateway/internal/services"
	"github.com/gin-gonic/gin"
)

// 设置项目的根路径，所有接口都加上api开头
const (
	rootPath = "/api"
)

func CmsRouter(r *gin.Engine) {
	//cms对象实例化
	cmsApp := services.NewCmsApp()
	//创建JWT中间件实例
	jwtAuth := NewJWTAuth()

	////资源监控中间件，上报数据到prometheus
	//r.Use(prometheusMiddleware())
	////链路追踪中间件，上报数据到zipkin
	//r.Use(opentracingMiddleware())

	//创建【路由组/cms/】
	cmsGroup := r.Group(rootPath + "/cms").Use(jwtAuth.Auth) //使用jwt中间件
	{
		// 商品
		//路径/api/cms/content/create
		cmsGroup.POST("/content/create", cmsApp.ContentCreate)
		//路径/api/cms/content/update
		cmsGroup.POST("/content/update", cmsApp.ContentUpdate)
		//路径/api/cms/content/delete
		cmsGroup.POST("/content/delete", cmsApp.ContentDelete)
		//路径/api/cms/content/find
		cmsGroup.POST("/content/find", cmsApp.ContentFind) //搜索框
		//路径/api/cms/content/get
		cmsGroup.POST("/content/get", cmsApp.ContentGet) //id查找商品信息

		// 用户
		//路径/api/cms/user/create
		//cmsGroup.POST("/user/create", cmsApp.UserCreate)
		//路径/api/cms/content/update
		cmsGroup.POST("/user/update", cmsApp.UserUpdate)
		//路径/api/cms/content/delete
		cmsGroup.POST("/user/delete", cmsApp.UserDelete)
		//路径/api/cms/content/find
		cmsGroup.POST("/user/find", cmsApp.UserFind)

		// order
		cmsGroup.POST("/order/place_order", cmsApp.PlaceOrder)
		cmsGroup.POST("/order/list_order", cmsApp.ListOrder)
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
