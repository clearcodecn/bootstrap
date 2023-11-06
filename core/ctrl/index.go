package ctrl

import (
	"github.com/gin-gonic/gin"
	"tools/core/templates"
)

func Index(ctx *gin.Context) {
	templates.Html(ctx, "index", gin.H{})
}
