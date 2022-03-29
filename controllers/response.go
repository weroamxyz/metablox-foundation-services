package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusBadRequest, &ResponseData{
		code,
		code.Msg(),
		nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusBadRequest, &ResponseData{
		code,
		msg,
		nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		CodeSuccess,
		CodeSuccess.Msg(),
		data,
	})
}

func ResponseSuccessData(c *gin.Context, data []byte) {
	c.Data(http.StatusOK, gin.MIMEJSON, data)
}

func ResponseSuccessWithMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		CodeSuccess,
		msg,
		nil,
	})
}

func ResponseSuccessWithMsgAndData(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		CodeSuccess,
		msg,
		data,
	})
}
