package handler

import (
	"gin-api/internal/logic"
	"gin-api/internal/pkg/response"
	"gin-api/internal/svc"
	"gin-api/internal/types"
	"github.com/gin-gonic/gin"
)

//
// init
// @Description: 初始化加载,需要注意手动init内容的加载顺序.例如配置文件读取
//
func init() {
	svcCtx = svc.NewServiceContext()
}

/*
统一返回示例
*/

func AddOne(c *gin.Context, params *types.AddOneRequest) (int, error) {
	return logic.AddOne(params.A)
}

// GetMap map返回示例
func GetMap(c *gin.Context, params *types.AddOneRequest) (map[string]string, error) {
	return logic.GetMap(params.A)
}

// GetSlice slice返回示例
func GetSlice(c *gin.Context, params *types.AddOneRequest) ([]int, error) {
	return logic.GetSlice(params.A)
}

// GetFloat float返回示例
func GetFloat(c *gin.Context, params *types.AddOneRequest) (float32, error) {
	return logic.GetFloat(params.A)
}

// GetInterface interface返回示例
func GetInterface(c *gin.Context, params *types.AddOneRequest) (interface{}, error) {
	return logic.GetInterface(params.A)
}

/*
以下是使用手动返回的示例
*/

var svcCtx *svc.ServiceContext

func AddUser(c *gin.Context, params *types.AddUserRequest) {
	l := logic.NewUserLogic(svcCtx)
	if err := l.AddUser(params); err != nil {
		response.NewResponseData().CustomizeResponse(c, 11001, "fail", err)
	} else {
		response.NewResponseData().SuccessResponse(c, "")
	}
}

func DelUser(c *gin.Context, params *types.DiffUserRequest) {
	l := logic.NewUserLogic(svcCtx)
	if err := l.DiffUser(params); err != nil {
		response.NewResponseData().CustomizeResponse(c, 11001, "fail", err)
	} else {
		response.NewResponseData().SuccessResponse(c, "")
	}
}
