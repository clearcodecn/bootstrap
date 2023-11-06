package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"tools/core/config"
	"tools/core/middleware"
	"tools/core/server/internal"
)

func StartServer(ch chan struct{}) {
	cfg := config.Get()

	g := gin.Default()

	// 多站点注入器.
	g.Use(middleware.CidInject())

	// 多语言.
	//g.Use(middleware.I18n)

	// 缓存.
	if cfg.Cache {
		g.Use(middleware.Cache(middleware.CacheConfig{
			Directory:           cfg.CachePath,
			Duration:            time.Duration(cfg.CacheDuration) * time.Second,
			SkipPrefixes:        middleware.DefaultPrefixes,
			StaticPrefixes:      middleware.DefaultStaticPrefixes,
			StaticCacheDuration: cfg.StaticCacheDuration,
		}))
	}

	g.GET("/ads.txt", internal.AdsTxt)
	g.GET("/sitemap.xml", internal.SiteMapXml)

	registerRouter(g)

	fmt.Println("server run at: http://localhost" + cfg.ListenAddress)
	httpServer := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: middleware.Minify(g),
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Errorf("server stop: %+v", err)
		}
	}()

	<-ch
	httpServer.Shutdown(context.Background())
}
