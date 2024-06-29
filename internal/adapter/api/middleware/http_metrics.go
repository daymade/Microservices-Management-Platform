package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	metricsmiddleware "github.com/slok/go-http-metrics/middleware"
	midgin "github.com/slok/go-http-metrics/middleware/gin"
)

// HttpMetrics 记录 HTTP 请求指标
func HttpMetrics(appName string) gin.HandlerFunc {
	mdlw := metricsmiddleware.New(metricsmiddleware.Config{
		Recorder:      prometheus.NewRecorder(prometheus.Config{}),
		Service:       appName,
		GroupedStatus: true,
	})

	// If there isn't predefined handler ID, the lib will set handlerID as each URL path, that's what we wanted.
	handlerID := ""
	return midgin.Handler(handlerID, mdlw)
}
