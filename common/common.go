package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

// 成功返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Json{
		Message: "ok",
		Code:    http.StatusOK,
		Data:    data,
	})
}

// 失败返回
func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Json{
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
	})
}

// 用户没有登录或者授权
func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Json{
		Message: "forbidden",
		Code:    http.StatusForbidden,
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, Json{
		Message: "not found",
		Code:    http.StatusNotFound,
	})
}

func MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, Json{
		Message: "method not allowed",
		Code:    http.StatusMethodNotAllowed,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Json{
		Message: "unauthorized",
		Code:    http.StatusUnauthorized,
	})
}
