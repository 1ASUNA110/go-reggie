package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

func Ok(data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS().Code,
		Msg:  SUCCESS().Msg,
		Data: data,
	})
}

func Fail(resultCodeEnum ResultCode, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: resultCodeEnum.Code,
		Msg:  resultCodeEnum.Msg,
		Data: nil,
	})
}
