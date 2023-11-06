package common

import (
	"github.com/gin-gonic/gin"
)

const (
	cidKey = "x-cid"
)

func WithCid(ctx *gin.Context, cid string) {
	ctx.Set(cidKey, cid)
}

func GetCid(ctx *gin.Context) string {
	return ctx.GetString(cidKey)
}
