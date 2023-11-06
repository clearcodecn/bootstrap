package templates

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"tools/core/common"
	"tools/core/config"
)

func html(s string) template.HTML {
	return template.HTML(s)
}

func empty(args ...interface{}) string { return "" }

type stringFunc func(s string) string

func asset(ctx *gin.Context) stringFunc {
	return func(s string) string {
		cid := common.GetCid(ctx)
		h := config.GetHost(cid)
		return filepath.Join("static", h.Theme, s)
	}
}
