package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SuccessCode            = 200
	CommonFailCode         = 10000
	BindParamsErrorCode    = 10001
	TokenValidationErrCode = 10002
	HandlerReturnErrorCode = 10008 // handler 返回值的数量不符合要求返回码
	UnKnowErrorCode        = 10009

	SuccessMessage            = "success"
	FailMessage               = "fail"
	UnKnowErrMessage          = "unknown error"
	BindParamsErrMessage      = "bind params fail"
	TokenValidationErrMessage = "token validation fail"
	HandlerReturnErrMessage   = "the quantity of handler func return does not meet the requirements" // handler 返回值的数量不符合要求
)

type Response interface {
	SuccessResponse(ctx *gin.Context, data interface{})
	CommonFailResponse(ctx *gin.Context, data error)
	CustomizeCodeResponse(ctx *gin.Context, code int, data error)
	CustomizeResponse(ctx *gin.Context, code int, msg string, data error)
	UnKnowErrorResponse(ctx *gin.Context, data error)
	BindParamsErrorResponse(ctx *gin.Context, data error)
}

type responseData struct{}

func NewResponseData() *responseData {
	return &responseData{}
}

// SuccessResponse 成功返回
func (r *responseData) SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    SuccessCode,
		"message": SuccessMessage,
		"data":    data,
	})
}

// CommonFailResponse 公共失败返回
func (r *responseData) CommonFailResponse(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    CommonFailCode,
		"message": FailMessage,
		"data":    data.Error(),
	})
}

func (r *responseData) CommonFailResponse1(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    CommonFailCode,
		"message": FailMessage,
		"data":    data.Error(),
	})
}

// CustomizeCodeResponse 自定义失败返回码
func (r *responseData) CustomizeCodeResponse(ctx *gin.Context, code int, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": FailMessage,
		"data":    data.Error(),
	})
}

// CustomizeResponse 自定义失败返回
func (r *responseData) CustomizeResponse(ctx *gin.Context, code int, msg string, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
		"data":    data.Error(),
	})
}

// UnKnowErrorResponse 未知错误返回
func (r *responseData) UnKnowErrorResponse(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    UnKnowErrorCode,
		"message": UnKnowErrMessage,
		"data":    data.Error(),
	})
}

// BindParamsErrorResponse 绑定参数错误返回
func (r *responseData) BindParamsErrorResponse(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    BindParamsErrorCode,
		"message": BindParamsErrMessage,
		"data":    data.Error(),
	})
}

// TokenValidationErrorResponse token验证错误
func (r *responseData) TokenValidationErrorResponse(ctx *gin.Context, data error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    TokenValidationErrCode,
		"message": TokenValidationErrMessage,
		"data":    data.Error(),
	})
}
