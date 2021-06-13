package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"easygo/internal/common/config"
	"easygo/pkg/wrapper"
	"github.com/gin-gonic/gin"
)

// CopyBodyMiddleware Copy body
func CopyBodyMiddleware() gin.HandlerFunc {
	var maxMemory int64 = 64 << 20 // 64 MB
	if v := config.C.HTTP.MaxContentLength; v > 0 {
		maxMemory = int64(v)
	}

	return func(c *gin.Context) {
		if c.Request.Body == nil {
			c.Next()
			return
		}

		var requestBody []byte
		safe := &io.LimitedReader{R: c.Request.Body, N: maxMemory}
		requestBody, _ = ioutil.ReadAll(safe)

		c.Request.Body.Close()
		bf := bytes.NewBuffer(requestBody)
		c.Request.Body = http.MaxBytesReader(c.Writer, ioutil.NopCloser(bf), maxMemory)
		c.Set(wrapper.ReqBodyKey, requestBody)

		c.Next()
	}
}
