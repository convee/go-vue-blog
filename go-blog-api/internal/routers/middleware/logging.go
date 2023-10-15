package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/convee/go-blog-api/pkg/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		//reg := regexp.MustCompile("(/v1/user|/login)")
		//if !reg.MatchString(path) {
		//	return
		//}

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// The basic informations.
		method := c.Request.Method
		ip := c.ClientIP()

		//log.Debugf("New request come in, path: %s, Method: %s, body `%s`", path, method, string(bodyBytes))
		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		headers := map[string]string{}
		for name, values := range c.Request.Header {
			headers[name] = strings.Join(values, "|")
		}

		logger.Info("access_log", zap.Duration("latency", latency), zap.String("ip", ip), zap.String("method", pad.Right(method, 5, "")), zap.String("path", path), zap.Int("status", blw.Status()), zap.Any("headers", headers), zap.Any("body", json.RawMessage(bodyBytes)), zap.Any("response", json.RawMessage(blw.body.Bytes())))
	}
}
