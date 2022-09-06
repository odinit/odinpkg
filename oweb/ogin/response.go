package ogin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type code int64

const (
	Err code = -1 - iota

	ErrServer // 后台服务

	ErrCmdExecFail // 终端执行

	ErrTokenIsInvalid // 用户
	ErrTokenParseFail
	ErrUserNotLogin
	ErrUserHasNoPermission

	ErrPath // 文件
	ErrFile

	ErrData // 数据格式

	ErrRedis // 数据库
	ErrMySQL

	ErrRequest // web相关
	ErrWebSocket
	ErrParamInvalid
	ErrParamProcess

	ErrNet // 网络相关

	ErrHard // 硬件相关

	ErrDocker // docker相关
)

const (
	OK code = 200 + iota
)

var codeMsgMap = map[code]string{
	OK:  "成功",
	Err: "异常",

	// 后台服务
	ErrServer: "服务异常",

	// 终端执行
	ErrCmdExecFail: "终端命令执行异常",

	// 用户
	ErrTokenIsInvalid:      "非法token",
	ErrTokenParseFail:      "token解析异常",
	ErrUserNotLogin:        "用户未登录",
	ErrUserHasNoPermission: "用户没有权限",

	// 文件
	ErrPath: "路径异常",
	ErrFile: "文件异常",

	// 数据格式
	ErrData: "数据操作异常",

	// 数据库
	ErrRedis: "redis操作异常",
	ErrMySQL: "mysql操作异常",

	// web相关
	ErrRequest:      "请求异常",
	ErrWebSocket:    "websocket异常",
	ErrParamInvalid: "参数格式异常",
	ErrParamProcess: "参数处理异常",

	// 网络相关
	ErrNet: "网络相关异常",

	// 硬件相关
	ErrHard: "硬件相关异常",

	// docker相关
	ErrDocker: "docker相关异常",
}

type ResData = map[string]interface{}

func NewResData(code code, data interface{}, extra ...map[string]interface{}) (resData ResData) {
	resData = ResData{
		"code": code,
		"msg":  codeMsgMap[code],
		"data": data,
	}
	for i := range extra {
		for k, v := range extra[i] {
			resData[k] = v
		}
	}
	return
}

func (c_ code) Return(c *gin.Context) {
	resData := ResData{
		"code": c_,
		"msg":  codeMsgMap[c_],
		"data": nil,
	}
	c.JSON(http.StatusOK, resData)
}

func (c_ code) Extra(c *gin.Context, extra ...map[string]interface{}) {
	resData := ResData{
		"code": c_,
		"msg":  codeMsgMap[c_],
		"data": nil,
	}
	for i := range extra {
		for k, v := range extra[i] {
			resData[k] = v
		}
	}
	c.JSON(http.StatusOK, resData)
}

func (c_ code) Data(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResData{
		"code": c_,
		"msg":  codeMsgMap[c_],
		"data": data,
	})
}

func (c_ code) DataExtra(c *gin.Context, data interface{}, extra ...map[string]interface{}) {
	resData := ResData{
		"code": c_,
		"msg":  codeMsgMap[c_],
		"data": data,
	}
	for i := range extra {
		for k, v := range extra[i] {
			resData[k] = v
		}
	}
	c.JSON(http.StatusOK, resData)
}

func (c_ code) Msg(c *gin.Context, msg ...string) {
	c.JSON(http.StatusOK, ResData{
		"code": c_,
		"msg":  codeMsgMap[c_] + ": " + strings.Join(msg, ", "),
		"data": nil,
	})
}

func (c_ code) MsgExtra(c *gin.Context, msg string, extra ...map[string]interface{}) {
	resData := ResData{
		"code": c_,
		"msg":  msg,
		"data": nil,
	}
	for i := range extra {
		for k, v := range extra[i] {
			resData[k] = v
		}
	}
	c.JSON(http.StatusOK, resData)
}

func (c_ code) CodeMsg() string {
	return codeMsgMap[c_]
}
