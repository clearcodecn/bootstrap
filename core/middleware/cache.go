package middleware

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	DefaultPrefixes = []string{
		"/static",
		"/api",
		"/admin",
		"/favico.ico",
	}

	DefaultStaticPrefixes = []string{
		"/static",
		"/favico.ico",
	}
)

type CacheConfig struct {
	Directory    string
	Duration     time.Duration
	SkipPrefixes []string

	StaticPrefixes      []string
	StaticCacheDuration int
}

var (
	mu = sync.Mutex{}
)

func Cache(config CacheConfig) gin.HandlerFunc {
	if config.Directory == "" {
		config.Directory = "cache"
	}

	os.MkdirAll(config.Directory, 0777)

	clean = func() {
		mu.Lock()
		os.RemoveAll(config.Directory)
		os.MkdirAll(config.Directory, 0777)
		mu.Unlock()
	}

	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.Next()
			return
		}

		isStatic := lo.ContainsBy(config.StaticPrefixes, func(item string) bool {
			return strings.HasPrefix(ctx.Request.RequestURI, item)
		})

		if isStatic {
			ctx.Header("Cache-Control", "")
			ctx.Header("Access-Control-Max-Age", fmt.Sprintf("%d", config.StaticCacheDuration)) // 缓存20分钟
			ctx.Header("Cache-Control", "public")
			ctx.Header("Cache-Control", fmt.Sprintf("max-age=%d", config.StaticCacheDuration))
			ctx.Next()
			return
		}

		mu.Lock()
		hash := buildHash(ctx.Request)
		path := filepath.Join(config.Directory, hash)
		fi, err := os.Stat(path)

		hit := true
		if err != nil {
			if !os.IsNotExist(err) {
				mu.Unlock()

				ctx.AbortWithError(503, err)
				return
			}
			hit = false
		}

		if fi != nil {
			if time.Now().Sub(fi.ModTime()) > config.Duration {
				hit = false
			} else {
				hit = true
			}
		}

		if hit {
			mu.Unlock()
			ctx.Header("x-cache-id", strings.TrimSuffix(hash, ".html"))
			ctx.Header("Content-Type", "text/html")
			ctx.File(path)
			ctx.Abort()
			logrus.Infof("hit cache %s %s %s", path, ctx.Request.Host, ctx.Request.URL.String())
			return
		}
		mu.Unlock()

		w := &responseWriter{ResponseWriter: ctx.Writer}
		ctx.Writer = w

		ctx.Next()
		if w.status != 200 {
			w.realWrite()
			return
		}

		mu.Lock()
		ioutil.WriteFile(path, w.buf, 0777)
		mu.Unlock()

		w.realWrite()
	}
}

func buildHash(r *http.Request) string {
	m := md5.New()
	m.Write([]byte(r.Host + r.URL.String()))
	isApp := strings.HasPrefix(r.RequestURI, "/app")
	if isApp {
		return fmt.Sprintf("%x.json", m.Sum(nil))
	}
	return fmt.Sprintf("%x.html", m.Sum(nil))
}

type responseWriter struct {
	gin.ResponseWriter
	buf    []byte
	status int
}

func (r *responseWriter) Write(b []byte) (int, error) {
	r.buf = append(r.buf, b...)
	return len(b), nil
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.status = statusCode
}

func (r *responseWriter) realWrite() {
	r.ResponseWriter.WriteHeader(r.status)
	r.ResponseWriter.Write(r.buf)
}

var clean func()

func CleanDir() {
	if clean == nil {
		return
	}
	clean()
}
