package middleware

import (
	"strings"

	"github.com/bytepac/greasyx/console"
	"github.com/bytepac/greasyx/gina"
	"github.com/bytepac/greasyx/helper"
	"github.com/bytepac/greasyx/libs/auth"
	"github.com/bytepac/greasyx/libs/xerror"
	"github.com/gin-gonic/gin"
)

func Casbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.ToLower(ctx.Request.Method) == "OPTIONS" {
			ctx.Next()
			return
		}

		roleId := auth.GetTokenData[int64](ctx, "role_id")
		if roleId == 0 {
			console.Echo.Info("ℹ️ 提示: 无法使用 `Casbin` 权限校验, 请确保 `Token` 中包含了 `role_id`")
			ctx.Next()
			return
		}

		path := helper.ConvertToRestfulURL(strings.TrimPrefix(ctx.Request.URL.Path, "/api"))
		success, _ := gina.Casbin.Enforce(helper.Int64ToString(roleId), path, ctx.Request.Method)
		if success {
			ctx.Next()
		} else {
			gina.Fail(ctx, xerror.NoAuth)
			ctx.Abort()
			return
		}
	}
}
