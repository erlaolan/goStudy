package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	//创建logrus实例
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logrus.WithFields(logrus.Fields{
			"status":      params.StatusCode,
			"method":      params.Method,
			"path":        params.Path,
			"ip":          params.ClientIP,
			"latency":     params.Latency,
			"user_agent":  params.Request.UserAgent(),
			"error":       params.ErrorMessage,
			"request_uri": params.Request.RequestURI,
			"timestamp":   params.TimeStamp.Format(time.RFC3339),
		}).Info("HTTP Request")
		return ""
	})
}

// 全局错误处理
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithFields(logrus.Fields{
					"error":  err,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("Panic Recovered")

				c.JSON(500, gin.H{
					"code":  500,
					"error": "Internal Server Error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
