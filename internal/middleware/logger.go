package middleware

import (
	"easygo/internal/common/config"
	"mime"
	"net/http"
	"time"

	"easygo/pkg/logger"
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := c.Request.URL.Path
		method := c.Request.Method

		entry := logger.WithContextPureLogger(logger.NewTagContext(c.Request.Context(), "__requset__"))

		// 获取 req 的信息与开始时间
		start := time.Now()
		fields := []zap.Field{
			zap.String("ip", c.ClientIP()),
			zap.String("method", method),
			zap.String("url", c.Request.URL.String()),
			zap.String("proto", c.Request.Proto),
			zap.String("user_agent", c.GetHeader("User-Agent")),
			zap.Int64("content_length", c.Request.ContentLength),
		}

		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType != "multipart/form-data" {
				if v, ok := c.Get(wrapper.ReqBodyKey); ok {
					if b, ok := v.([]byte); ok && len(b) <= config.C.HTTP.MaxContentLength {
						fields = append(fields, zap.String("req_body", string(b)))
					}
				}
			}
		}

		c.Next()

		// 获得 res 的信息并计算结束时间
		timeConsuming := time.Since(start).Nanoseconds() / 1e6
		fields = append(fields, zap.Int("res_status", c.Writer.Status()))
		fields = append(fields, zap.Int("res_length", c.Writer.Size()))

		if v, ok := c.Get(wrapper.LoggerReqBodyKey); ok {
			if b, ok := v.([]byte); ok && len(b) <= config.C.HTTP.MaxLoggerLength {
				fields = append(fields, zap.String("log_body", string(b)))
			}
		}

		if v, ok := c.Get(wrapper.ResBodyKey); ok {
			if b, ok := v.([]byte); ok && len(b) <= config.C.HTTP.MaxLoggerLength {
				fields = append(fields, zap.String("res_body", string(b)))
			}
		}

		fields = append(fields, zap.String(logger.UserIDKey, wrapper.GetUserID(c)))
		entry.With(fields...).Sugar().Infof(
			"[http] %s-%s-%s-%d(%dms)",
			p,
			c.Request.Method,
			c.ClientIP(),
			c.Writer.Status(),
			timeConsuming,
		)
	}
}
