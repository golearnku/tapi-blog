package middleware

import (
	"github.com/china-golang/tapi-blog/handler"
	"github.com/china-golang/tapi-blog/pkg/errno"
	"github.com/china-golang/tapi-blog/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
