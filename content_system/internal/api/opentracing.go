package api

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	reporter "github.com/openzipkin/zipkin-go/reporter/http"
	"net/http"
)

// 链路追踪中间件
// 上报中间件 上报地址：http://localhost:9411/api/v2/spans
func opentracingMiddleware() gin.HandlerFunc {
	// 创建zipkin的上报节点Reporter：初始化，上报节点，指定zipkin的服务节点地址
	report := reporter.NewReporter("http://localhost:9411/api/v2/spans")
	// 创建本地节点,以及指定网关层的地址
	endpoint, err := zipkin.NewEndpoint("content_system", "localhost:8080")
	if err != nil {
		panic(err)
	}
	// 创建Zipkin Tracer
	tracer, err := zipkin.NewTracer(report,
		zipkin.WithLocalEndpoint(endpoint), //指定网关层的地址(本地节点)
		zipkin.WithTraceID128Bit(true))     //TraceID,使用128位
	if err != nil {
		panic(err)
	}
	//zipkin-go-opentracing：把zipkin节点，转化为opentracing节点（opentracing桥接转化）
	zipTracer := zipkinot.Wrap(tracer)
	//配置opentracing
	opentracing.SetGlobalTracer(zipTracer)
	// 创建中间件
	return ginhttp.Middleware(zipTracer, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return r.URL.Path
	}))
}
