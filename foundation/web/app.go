package web

import (
	"os"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type App struct {
	*routing.Router
	log      *zap.SugaredLogger
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(router *routing.Router, log *zap.SugaredLogger) *App {
	return &App{router, log, nil, nil}
}

func (a *App) Run(port string) {
	a.log.Infof("Server startup on port : %s", port)
	fasthttp.ListenAndServe(port, a.HandleRequest)

}
