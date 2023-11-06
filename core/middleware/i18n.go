package middleware

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"tools/lang"
)

var (
	bundle = i18n.NewBundle(language.SimplifiedChinese) // 简体中文.
)

func init() {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFileFS(lang.LocaleFS, "locale.zh.toml")
}

type I18nManager struct {
}

func I18n() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.Request.FormValue("lang")
		accept := ctx.Request.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(bundle, lang, accept)
		_ = localizer
		// TODO:: implement me
	}
}
