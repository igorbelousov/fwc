package web

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type App struct {
	router *routing.Router
	log    *zap.SugaredLogger
}

func NewApp(router *routing.Router, log *zap.SugaredLogger) *App {
	return &App{router, log}
}

func (a *App) Run() {
	a.log.Infow("start server on port 8080")
	fasthttp.ListenAndServe(":8080", a.router.HandleRequest)
}
