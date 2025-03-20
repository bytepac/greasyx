package middleware

import (
	"strings"

	"github.com/bytepac/greasyx/gina"
	"github.com/bytepac/greasyx/libs/auth"
	"github.com/bytepac/greasyx/libs/xerror"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			gina.Fail(ctx, xerror.NeedLogin)
			ctx.Abort()
			return
		}

		claims, err := auth.ParseJwtToken(tokenString[7:])
		if err != nil {
			gina.Fail(ctx, xerror.NeedLogin)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
