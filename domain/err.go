package domain

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

type ERR struct {
	Error error
	Code  uint
}

func Panic(recode Recode, extraMsg string) {
	panic(New(recode, extraMsg))
}

func New(recode Recode, extraMsg string) ERR {
	var msg string
	if extraMsg == "" {
		msg = recode.Msg
	} else {
		msg = fmt.Sprintf("%s : %s", recode.Msg, extraMsg)
	}
	return ERR{Code: recode.Code, Error: errors.New(msg)}
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()

			if err, ok := r.(ERR); ok {
				Error(c, err.Code, err.Error.Error())
			}

			c.Abort()
		}
	}()
	c.Next()
}
