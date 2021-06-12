package wrapper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"easygo/pkg/logger"
	"github.com/gin-gonic/gin"
)

// 定义 gin 上下文中的键
const (
	prefix           = "insider"
	UserIDKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

// GetUserID 获取用户ID
func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

// SetUserID 设定用户ID
func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

// ParseJSON 解析请求JSON
func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": v,
	})
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *ResponseError

	if err != nil {
		if e, ok := err.(*ResponseError); ok {
			res = e
		} else {
			res = UnWrapResponse(ErrInternalServer)
			res.ERR = err
		}
	} else {
		res = UnWrapResponse(ErrInternalServer)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}

		if status := res.StatusCode; status >= 400 && status < 500 {
			logger.WithContext(ctx).Warnf(err.Error())
		} else if status >= 500 {
			logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	ResJSON(c, res.StatusCode, gin.H{
		"code": res.Code,
		"msg":  res.Message,
		"data": res.Error(),
	})
}

//ParseForm 解析请求表单
func ParseForm(c *gin.Context, obj interface{}) error {

	if err := c.ShouldBind(obj); err != nil {
		ctx := c.Request.Context()
		ctx = logger.NewTagContext(ctx, "__parseform__")
		return Wrap400Response(err, fmt.Sprintf("解析请求参数发生错误 - %s", err.Error()))
	}
	return nil
}
