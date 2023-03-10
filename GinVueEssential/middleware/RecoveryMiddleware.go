package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"herrkung.com/GinVueEssential/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()
		ctx.Next()
	}
}
