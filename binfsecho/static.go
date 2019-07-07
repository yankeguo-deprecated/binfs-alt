package binfsecho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.guoyk.net/binfs"
	"net/http"
	"net/url"
	"strings"
)

type StaticConfig struct {
	Skipper middleware.Skipper
	Root    string
	Index   []string
}

var (
	DefaultStaticConfig = StaticConfig{
		Skipper: middleware.DefaultSkipper,
		Root:    "public",
		Index:   []string{"index.html"},
	}
)

func Static(root string) echo.MiddlewareFunc {
	return StaticWithConfig(StaticConfig{Root: root})
}

func StaticWithConfig(cfg StaticConfig) echo.MiddlewareFunc {
	if cfg.Skipper == nil {
		cfg.Skipper = DefaultStaticConfig.Skipper
	}
	if cfg.Root == "" {
		cfg.Root = DefaultStaticConfig.Root
	}
	if cfg.Index == nil {
		cfg.Index = DefaultStaticConfig.Index
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			if cfg.Skipper(ctx) {
				return next(ctx)
			}
			p := ctx.Request().URL.Path
			if p, err = url.PathUnescape(p); err != nil {
				return
			}
			n := binfs.Find(strings.Split(cfg.Root+"/"+p, "/")...)
			if n != nil && n.Chunk != nil {
				http.ServeContent(ctx.Response(), ctx.Request(), n.Name, n.ModTime(), n.ReadSeeker())
				return nil
			}
			for _, index := range cfg.Index {
				n = binfs.Find(strings.Split(cfg.Root+"/"+p+"/"+index, "/")...)
				if n != nil && n.Chunk != nil {
					http.ServeContent(ctx.Response(), ctx.Request(), n.Name, n.ModTime(), n.ReadSeeker())
					return nil
				}
			}
			return next(ctx)
		}
	}
}
