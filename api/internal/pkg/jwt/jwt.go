package jwt

import (
	"errors"
	"gin-api/internal/config"
	jwtpkg "github.com/dgrijalva/jwt-go"
	"time"
)

//
// JWT
// @Description: JWT对象
//
type JWT struct {
	SignKey   []byte        // 秘匙
	MaxExpire time.Duration // 最大过期时间,固定30分钟过期,就不考虑刷新token了
}

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

//
// CustomClaims
// @Description: 自定义载荷,这里采用自定义的UserName和Password作为有效载荷的一部分
//
type CustomClaims struct {
	Username     string `json:"username"`
	ExpireAtTime int64  `json:"expireAtTime"`

	// StandardClaims结构体实现了Claims接口(Valid()函数)
	// JWT 规定了7个官方字段，提供使用:
	// - iss (issuer)：发布者
	// - sub (subject)：主题
	// - iat (Issued At)：生成签名的时间
	// - exp (expiration time)：签名过期时间
	// - aud (audience)：观众，相当于接受者
	// - nbf (Not Before)：生效时间
	// - jti (JWT ID)：编号
	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey:   []byte(config.JWTConf.SignKey),
		MaxExpire: time.Duration(config.JWTConf.MaxExpire) * time.Minute,
	}
}

//
// GenerateToken
// @Description: token 生成,使用jwt-go库生成token,指定编码的算法为jwt.SigningMethodHS256
// @receiver j
// @param username
// @return string
// @return error
//
func (j *JWT) GenerateToken(username string) (string, error) {
	// 构造用户claims
	expireAtTime := j.expireAtTime()
	claims := CustomClaims{
		username,
		expireAtTime,
		jwtpkg.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAtTime,
			Issuer:    config.Srv.Name,
		},
	}

	// 根据claims生成token
	if token, err := j.createToken(claims); err != nil {
		return "", err
	} else {
		return token, err
	}
}

//
// ParserToken
// @Description: token解析
// @receiver j
// @param tokenStr token字符串
// @return *CustomClaims 自定义的claims
// @return error
//
func (j *JWT) ParserToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwtpkg.ParseWithClaims(tokenStr, &CustomClaims{}, func(r *jwtpkg.Token) (interface{}, error) {
		return j.SignKey, nil
	})

	// 解析错误分析
	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwtCli.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwtpkg.ValidationError); ok {
			if ve.Errors&jwtpkg.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwtpkg.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
			return nil, ErrTokenInvalid
		}
	}

	// 对比解析出的claims与CustomClaims
	if token == nil {
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

//
// expireAtTime
// @Description: 过期时间
// @receiver j
// @return int64
//
func (j *JWT) expireAtTime() int64 {
	var expireTime int64
	now := time.Now()
	expire := time.Duration(expireTime) * time.Minute
	return now.Add(expire).Unix()
}

//
// createToken
// @Description: 生成token的内部方法
// @receiver j
// @param claims
// @return string
// @return error
//
func (j *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodES256, claims)
	return token.SignedString(j.SignKey)
}
