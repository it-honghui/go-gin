package domain

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type R struct {
	Data interface{} `json:"data,omitempty"`
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, R{
		Data: data,
		Code: SUCCESS.Code,
		Msg:  SUCCESS.Msg,
	})
}

func Error(c *gin.Context, code uint, msg string) {
	c.JSON(http.StatusOK, R{
		Data: nil,
		Code: code,
		Msg:  msg,
	})
}

func Unauthorized(c *gin.Context, code uint, msg string) {
	c.JSON(http.StatusUnauthorized, R{
		Data: nil,
		Code: code,
		Msg:  msg,
	})
}
