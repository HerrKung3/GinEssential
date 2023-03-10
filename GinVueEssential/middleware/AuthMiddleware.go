package middleware

import (
	"net/http"
	"strings"

	"herrkung.com/GinVueEssential/response"

	"github.com/gin-gonic/gin"
	"herrkung.com/GinVueEssential/common"
	"herrkung.com/GinVueEssential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1 get authorization header
		tokenString := ctx.GetHeader("Authorization")

		//2 validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			//ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg":  "permission denied"})
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
			ctx.Abort()
			return
		}

		//3 get header payload
		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			//ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg":  "permission denied"})
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
			ctx.Abort()
			return
		}

		//4 get usrID from claims
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		//5 find user from database via userID
		DB.First(&user, userId)

		//6 user is exists?
		if user.ID == 0 {
			//ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg":  "权限不足"})
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "permission denied")
			ctx.Abort()
			return
		}
		//if user is exists
		ctx.Set("user", user)
		ctx.Next()
	}
}
