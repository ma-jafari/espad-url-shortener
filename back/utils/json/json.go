package json

import (
	"github.com/kataras/iris/v12"
	utils "url-shortener/utils/errors"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error int64       `json:"error"`
	Msg   string      `json:"msg"`
}

func WriteError(ctx iris.Context, errorCode int64) {
	response := &Response{}
	response.Data = nil
	response.Error = errorCode
	response.Msg = utils.ErrorMessages[int(errorCode)]

	ctx.StatusCode(int(errorCode))
	ctx.JSON(response)
}

func WriteErrorWithMsg(ctx iris.Context, errorCode int64, msg string) {
	response := &Response{}
	response.Data = nil
	response.Error = errorCode
	response.Msg = msg

	ctx.StatusCode(int(errorCode))
	ctx.JSON(response)
}

func WriteResponse(ctx iris.Context, data interface{}) {
	response := &Response{}
	response.Data = data

	ctx.JSON(response)
}
