package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

// 把 server 要 Use 的 handler func 写在这里
// 写在这个 builder 的 build 方法里
func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
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

		// 验证
		// sess 不可能没有，因为已经在前面创建过 seesion 的中间件了（mysession 那个）
		sess := sessions.Default(ctx)
		id := sess.Get("user_id")
		if id == nil {
			// 没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		// 多次 set option 一次 save 生效
		sess.Set("user_id", id)
		sess.Options(sessions.Options{
			MaxAge: 60,
		})
		now := time.Now().UnixMilli()
		if updateTime == nil {
			// 说明是第一次登陆，还没有刷新时间
			// 那么就设置刷新时间
			sess.Set("update_time", now)
			_ = sess.Save()
			return
		}

		// 如果有 updateTime
		// 这里通过 Get 方法取出来的 updateTime 是一个 interface
		// 需要断言一下才能用
		updateTimeVal, _ := updateTime.(int64)

		if now-updateTimeVal > 60*1000 {
			sess.Set("update_time", now)
			_ = sess.Save()
			return
		}
	}
}
