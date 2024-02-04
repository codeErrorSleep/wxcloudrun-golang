package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 请求参数
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 把读过的body再放回去，因为http的body只能读一次

		// 替换 response writer
		writer := &ResponseWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = writer

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 响应体
		respBody := writer.body.String()

		// 日志格式
		fmt.Printf("请求方法: %s, 请求URL: %s, 请求IP: %s, 状态码: %d, 执行时间: %v, 请求参数: %s, 响应体: %s\n",
			reqMethod, reqUri, clientIP, statusCode, latencyTime, bodyBytes, respBody)
	}
}
