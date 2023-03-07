package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Payload interface{} `json:"payload"`
}

const (
	InteralStatusSuccess = 0
	InteralStatusFailed  = 1
)

func Response200(internalCode int, msg string, payload interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		internalCode, msg, payload,
	})
}

func Success(msg string, payload interface{}, ctx *gin.Context) {
	Response200(InteralStatusSuccess, msg, payload, ctx)
}

func Fail(msg string, payload interface{}, ctx *gin.Context) {
	Response200(InteralStatusFailed, msg, payload, ctx)
}
