package router

import (
	demo "web/src/controller"
	"web/src/tpl" // template

	//"text/template" //text template
	"net/http"
	echopprof "web/src/pprof"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware" //v4 use v4 middleware
)

//Init Router
func Init() *echo.Echo {

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash()) //request uri 最後面自動加 /

	t := tpl.Init()

	e.Renderer = t
	/*
		https://echo.labstack.com/guide/routing
		I would like to apply the same middlewares for all /admin/* URLs. Is it possible ?
		g := e.Group("/admin", <your-middleware>)
		g.GET("/secured", <your-handler>)
		Now route /admin/secured will execute your-middleware.
		https://github.com/labstack/echo/issues/613
	*/
	e.GET("/hello/:page/*", demo.Demo)

	e.GET("/hello/*", demo.Demo)

	//e.GET("/", api.Home)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//所有 public folder 的文件都可以被 access ,例如 public/robots.txt -> http://yourhost/robots.txt
	e.Static("/", "public")

	// EX:go tool pprof http://192.168.119.128/
	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	return e
}
