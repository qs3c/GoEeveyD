package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

// 把 server 要 Use 的 handler func 写在这里
// 写在这个 builder 的 build 方法里
func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {

		for _, path := range l.paths {
			if ctx.Request.RequestURI == path {
				return
			}
		}
		// 登陆以后访问的后续页面才需要登陆状态验证
		// 注册和登录这两个页面是不需要登陆状态验证的
		if ctx.Request.URL.Path == "users/signup" || ctx.Request.URL.Path == "users/login" {
			return
		}

		// 验证 jwt token
		// 前端返回的 jwt token 放在 authorization 中
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			// 没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		// 从签名的 token 中反解析回 token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("H823kgHYwvHm9BltzLty2ZFU0vxBPVpg"), nil
		})
		if err != nil {
			// token解析失败，没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			// token不合法，没登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
