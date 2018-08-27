package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yilingfeng/apiserver/pkg/token"
	"github.com/yilingfeng/apiserver/handler"
	"github.com/yilingfeng/apiserver/pkg/errno"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		
		c.Next()
	}
}
