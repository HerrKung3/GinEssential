package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//预检请求响应
		//只要响应头里面的Origin里面本域（协议、域名、端口）一样，就可以接受到数据
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8989")
		//表示该响应的有效期，单位为秒。在有效时间内，浏览器无须为同一请求再次发起预检请求。
		//还有一点需要注意，该值要小于浏览器自身维护的最大有效时间，否则是无效的。
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//表示服务器允许客户端使用方法发起请求，可以一次设置多个，表示服务器所支持的所有跨域方法，而不单是当前请求那个方法，这样好处是为了避免多次预检请求。
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		//表示服务器允许请求中携带字段，也可以设置多个
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//Access-Control-Allow-Credentials 是否允许跨域使用cookies
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
