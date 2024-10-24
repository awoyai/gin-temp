package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

const (
	MSG_TOKEN_INVALID       = "登录信息过期或无效"
	MSG_INTERNAL_ERROR      = "内部错误，请联系管理员"
	MSG_SCHEME_SEARCH_ERROR = "数据库查询失败"
	MSG_BAD_PARAM           = "参数错误"
)

func Result(code, status int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(status, Response{
		code,
		data,
		msg,
	})
	c.AbortWithStatus(status)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, http.StatusOK, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, http.StatusOK, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, http.StatusOK, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, http.StatusOK, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, http.StatusOK, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, http.StatusOK, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, http.StatusOK, data, message, c)
}

func FailWithUnauthorized(c *gin.Context) {
	Result(ERROR, http.StatusUnauthorized, map[string]interface{}{}, "权限校验失败", c)
}

func FailForbidden(c *gin.Context) {
	Result(ERROR, http.StatusForbidden, map[string]interface{}{}, "禁止访问", c)
}

func FailToken(c *gin.Context) {
	Result(ERROR, http.StatusForbidden, map[string]interface{}{}, "token失效，请重新登录", c)
}
