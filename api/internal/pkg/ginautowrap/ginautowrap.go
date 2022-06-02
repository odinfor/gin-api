package ginautowrap

import (
	"errors"
	"fmt"
	"gin-api/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

//
// AutoBindWrap
// @Description: 自动绑定参数的包装方法减少bind参数垄余代码,例如: func NeedBindGet(c *gin.content, params *GetParams)
// @param ctrFunc 签名函数
// @return gin.HandlerFunc
//
func AutoBindWrap(handlerFunc interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取handler func函数参数struct,创建参数实例
		handlerFuncType := reflect.TypeOf(handlerFunc)
		handlerFuncValue := reflect.ValueOf(handlerFunc)

		// 检查参数类型,第一个
		if handlerFuncType.Kind() != reflect.Func {
			panic("not support type, handler must be a func type!")
			return
		}
		numIn := handlerFuncType.NumIn()
		if numIn != 2 {
			panic("not support params len. must need two params!")
			return
		}

		// bind参数
		handlerFuncParams := make([]reflect.Value, numIn)
		for i := 0; i < numIn; i++ {
			pt := handlerFuncType.In(i)
			// handle gin.content
			if pt == reflect.TypeOf(&gin.Context{}) {
				handlerFuncParams[i] = reflect.ValueOf(context)
				continue
			}
			// handle params, 根据请求方法bind
			if pt.Kind() == reflect.Ptr && pt.Elem().Kind() == reflect.Struct {
				pv := reflect.New(pt.Elem()).Interface()
				var err error
				switch context.Request.Method {
				case http.MethodGet:
					err = context.ShouldBindQuery(pv)
				default:
					err = context.ShouldBindJSON(pv)
				}
				if err != nil {
					response.NewResponseData().BindParamsErrorResponse(context, errors.New("bind params found error: "+err.Error()))
					return
				}
				handlerFuncParams[i] = reflect.ValueOf(pv)
			}
		}
		// 调用真实方法
		res := handlerFuncValue.Call(handlerFuncParams)

		// 不使用包装返回，handler需要手动调用返回,并且不能有返回内容。
		if res == nil {
			return
		}

		// handler返回验证,必须有2个返回
		// 一个返回data,作为response的返回数据
		// 一个error,不为空时作为错误内容返回
		if res != nil && len(res) != 2 {
			response.NewResponseData().CustomizeCodeResponse(context, response.HandlerReturnErrorCode, errors.New(response.HandlerReturnErrMessage))
			return
		}

		// 没有错误成功返回
		if res[1].IsNil() {
			switch res[0].Kind() {
			case reflect.String:
				response.NewResponseData().SuccessResponse(context, res[0].String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				response.NewResponseData().SuccessResponse(context, res[0].Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				response.NewResponseData().SuccessResponse(context, res[0].Uint())
			case reflect.Float32, reflect.Float64:
				response.NewResponseData().SuccessResponse(context, res[0].Float())
			case reflect.Bool:
				response.NewResponseData().SuccessResponse(context, res[0].Bool())
			default:
				response.NewResponseData().SuccessResponse(context, res[0].Interface())
			}
			return
		}

		// 出现错误,res[1]为error对象,统一使用通用错误返回
		response.NewResponseData().CommonFailResponse(context, fmt.Errorf("%v", res[1].Interface()))
	}
}
