package handlers

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Api(db *sqlx.DB) *routing.Router {
	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

	return router
}
