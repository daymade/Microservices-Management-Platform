package middleware

import (
	"fmt"
	"time"

	"github.com/VictoriaMetrics/metrics"
	"github.com/gin-gonic/gin"
)

// Metrics 记录请求的指标
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		metrics.GetOrCreateHistogram(fmt.Sprintf(`http_request_duration_seconds{method="%s", path="%s", status="%d"}`,
			c.Request.Method, c.FullPath(), c.Writer.Status())).Update(duration.Seconds())
	}
}
