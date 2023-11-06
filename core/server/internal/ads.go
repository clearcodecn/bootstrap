package internal

import (
	"github.com/gin-gonic/gin"
	"tools/core/common"
	"tools/core/config"
)

func AdsTxt(ctx *gin.Context) {
	cid := common.GetCid(ctx)
	host := config.GetHost(cid)
	ctx.String(200, host.AdsTxt)
}
