package handlers

import (
	"fmt"

	"github.com/fasthttp/router"
	"github.com/igorbelousov/fwc/foundation/database"
	"github.com/igorbelousov/fwc/foundation/web"
	"github.com/valyala/fasthttp"
)

func Api() *router.Router {

	r := router.New()

	r.GET("/", readiness)

	return r
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!")

}

func readiness(ctx *fasthttp.RequestCtx) {

	status := "ok"
	statusCode := fasthttp.StatusOK
	if err := database.DB.StatusCheck(); err != nil {
		status = "db not ready"
		statusCode = fasthttp.StatusInternalServerError
	}

	health := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	web.Respond(ctx, health, statusCode)

}
