package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"longmen-gateway/global"
	"net/http"
	"strings"
)

type ResponseCode int

type Response struct {
	Code ResponseCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data"`
}

func ErrorResponse(ctx *gin.Context, code ResponseCode, err error) {
	resp := &Response{Code: code, Msg: err.Error(), Data: nil}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, resp)
	} else {
		resp.Msg = RemoveTopStruct(errs.Translate(global.Trans))
		ctx.JSON(http.StatusBadRequest, resp)
	}

	response, _ := json.Marshal(resp)
	defer zap.S().Error(string(response))
	ctx.Abort()
}

// RemoveTopStruct 美化参数校验提示信息
func RemoveTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	resp := &Response{Code: http.StatusOK, Msg: nil, Data: data}
	ctx.JSON(200, resp)

	if global.ServerInfo.LogInfo.EnableSuccessResponseLog {
		response, _ := json.Marshal(resp)
		defer zap.S().Info(string(response))
	}
}
