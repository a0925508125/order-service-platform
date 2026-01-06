package base_controller

import (
	"fmt"
	"net/http"
	"order-service-platform/errcode"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c *BaseController) JsonResponse(ctx *gin.Context, code int, data interface{}, errorMessage ...interface{}) {
	message, ok := errcode.CodeMapMessage[int32(code)]

	if !ok {
		message = "未定義"
	}

	if errorMessage != nil && len(errorMessage) != 0 {
		message = errcode.CodeMapMessage[int32(code)] + fmt.Sprint(errorMessage...)
	}

	ctx.JSON(http.StatusOK, Resp{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//func NewPbResp(code commonpb.RespCode) *commonpb.Resp {
//	msg := code.String()
//	return &commonpb.Resp{
//		Code:    code.Enum(),
//		Message: &msg,
//	}
//}
