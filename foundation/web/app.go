package web

import (
	"os"
	"syscall"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type App struct {
	*router.Router
	log      *zap.SugaredLogger
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(router *router.Router, log *zap.SugaredLogger, shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{router, log, shutdown, mw}
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Run(port string) {
	withCors := cors.DefaultHandler()
	a.log.Infof("Server startup on port : %s", port)
	a.log.Fatal(fasthttp.ListenAndServe(port, withCors.CorsMiddleware(a.Handler)))

}
