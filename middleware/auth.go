package middleware

import (
	"ea-competition-api/library"
	"ea-competition-api/library/constant"
	"ea-competition-api/library/response"
	"ea-competition-api/library/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户授权中间件
func UserAuthMiddleware(token *token.Token, skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//跳过路由不检测授权
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if t := library.GetToken(c); t != "" {
			if userId := library.ParseUserID(token, c, t); userId > 0 {
				c.Set(constant.UserId, userId)
				c.Next()
				return
			} else {
				response.Error(c, http.StatusUnauthorized, "unauthorized")
				return
			}
		}

		response.Error(c, http.StatusUnauthorized, "unauthorized")
	}
}
