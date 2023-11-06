package templates

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
	"path/filepath"
	"sync"
	"tools/core/common"
	"tools/core/config"
)

var themeStore sync.Map

func Html(ctx *gin.Context, name string, data gin.H) {
	cid := common.GetCid(ctx)
	cfg := config.Get()
	host := config.GetHost(cid)
	var (
		themeTemplate *template.Template
	)
	tr := ThemeRender{}
	theme, ok := themeStore.Load(host.Theme)
	if !ok {
		tplPath := filepath.Join(cfg.TemplatePath, host.Theme)
		tpl, err := template.New("").Funcs(tr.defaultFuncMaps()).ParseGlob(filepath.Join(tplPath, "*.gohtml"))
		if err != nil {
			panic("failed to parse html: " + err.Error())
		}
		if !cfg.Debug {
			themeStore.Store(host.Theme, tpl)
		}
		themeTemplate = tpl
	} else {
		themeTemplate = theme.(*template.Template)
	}

	if data == nil {
		data = gin.H{}
	}

	data["host"] = host

	ctx.Header("Content-Type", "text/html")
	if err := themeTemplate.Funcs(tr.realFuncMaps(ctx)).ExecuteTemplate(ctx.Writer, name, data); err != nil {
		logrus.Errorf("render failed: %+v", err)
	}
}

type ThemeRender struct{}

func (ThemeRender) defaultFuncMaps() template.FuncMap {
	return template.FuncMap{
		"html":  html,
		"asset": empty,
	}
}

func (ThemeRender) realFuncMaps(ctx *gin.Context) template.FuncMap {
	return template.FuncMap{
		"html":  html,
		"asset": asset(ctx),
	}
}

func CleanTheme() {
	themeStore.Range(func(key, value any) bool {
		themeStore.Delete(key)
		return true
	})
}
