package middleware

import (
	"errors"
	"gin-api/internal/pkg/jwt"
	"gin-api/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

//
// JWTAuth
// @Description: header中验证token，验证通过放行否则abort中断并且直接return
// @return gin.HandlerFunc
//
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jenkins 上报接口忽略token认证
		if c.Request.URL.String() == "/devops/publish/v2/jenkinsCallback" {
			c.Next()
			return
		}
		// k8s相关忽略token认证
		if strings.HasPrefix(c.Request.URL.String(), "/devops/kubernetes") {
			c.Next()
			return
		}
		// 通过http header中的token解析来认证
		token := c.Request.Header.Get("token")
		if token == "" {
			resp := response.NewResponseData()
			resp.TokenValidationErrorResponse(c, errors.New("缺少token"))
			c.Abort()
			return
		}

		// 初始化一个JWT对象,并根据结构体方法来解析token
		j := jwt.NewJWT()

		// 解析token中包含的相关信息
		if claims, err := j.ParserToken(token); err != nil {
			switch err {
			case jwt.ErrTokenInvalid:
				resp := response.NewResponseData()
				resp.TokenValidationErrorResponse(c, jwt.ErrTokenInvalid)
			case jwt.ErrTokenExpired:
				resp := response.NewResponseData()
				resp.TokenValidationErrorResponse(c, jwt.ErrTokenExpired)
			case jwt.ErrTokenMalformed:
				resp := response.NewResponseData()
				resp.TokenValidationErrorResponse(c, jwt.ErrTokenMalformed)
			default: // 其他错误
				resp := response.NewResponseData()
				resp.TokenValidationErrorResponse(c, err)
			}
			c.Abort()
			return
		} else {
			// 将解析后的有效载荷claims重新写入gin.context引用对象中
			c.Set("claims", claims)
			c.Next()
		}
	}
}
