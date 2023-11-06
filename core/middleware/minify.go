package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/xml"
	"net/http"
	"regexp"
)

func Minify(g *gin.Engine) http.Handler {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	return m.Middleware(g)
}
