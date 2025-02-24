package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

// 资源监控的中间件
// Prometheus官方只实现了几个指标，这里我们自定义一些指标，用中间件的方式，对接口进行监控
func prometheusMiddleware() gin.HandlerFunc {
	//创建Prometheus的Counter(计数器)实例：只增不减，不能重置
	//统计各个接口的请求次数
	requestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",          //这个中间件的名称
		Help: "Total number of http request", //这个中间件的解释
	}, []string{
		"method", //接口不同方法的统计
		"path",   //接口不同路径的统计
	})
	//统计各个接口响应的状态码的次数
	requestsCodeTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_code_total",
		Help: "Total number of http request code",
	}, []string{
		"method",
		"path",
		"code",
	})

	//Summary指标，请求时间用时的平均数和中位数等等
	requestDuration := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name: "http_request_duration_seconds", //每秒的请求周期
		Help: "Http request duration in seconds",
		Objectives: map[float64]float64{ //误差范围
			0.5:  0.05,
			0.90: 0.01,
			0.99: 001,
		},
	}, []string{
		"method",
		"path",
	})

	//把上述的几种自定义的监控指标，注册到prometheus中
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestsCodeTotal)

	return func(c *gin.Context) {
		//定义：开始的时间
		start := time.Now()
		//得到请求的类型：Get Post ....
		method := c.Request.Method
		//得到请求的路由路径：/api/user/xxx
		path := c.FullPath()
		//

		//把requestsTotal指标，上报到prometheus中
		requestsTotal.WithLabelValues(method, path).Inc() //Inc是加1操作，自加

		c.Next()

		//执行完请求后，结束的时间，得到时间差
		elapsed := time.Since(start).Seconds()
		//把requestDuration指标上报prometheus中
		requestDuration.WithLabelValues(method, path).Observe(elapsed)

		statusCode := c.Writer.Status() //这个方法可以拿到响应的状态码
		//把requestsCodeTotal指标上报prometheus中
		requestsCodeTotal.WithLabelValues(method, path, strconv.Itoa(statusCode)).Inc() //strconv.Itoa(statusCode)，把int类型转为string类型
	}
}
