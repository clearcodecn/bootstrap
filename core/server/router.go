package server

import (
	"github.com/gin-gonic/gin"
	"tools/core/ctrl"
)

func registerRouter(g *gin.Engine) {
	g.GET("/", ctrl.Index)
}
