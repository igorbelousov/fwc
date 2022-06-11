package handlers

import (
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func Api() *routing.Router {
	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

	return router
}
