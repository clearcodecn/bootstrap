package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"tools/core/common"
	"tools/core/config"
)

func CidInject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cid := config.Get().DevId
		if cid == "" {
			cid = ctx.GetHeader("cid")
		}
		if cid == "" {
			ctx.AbortWithError(503, errors.New("cid not set"))
			return
		}
		common.WithCid(ctx, cid)
		ctx.Next()
	}
}
