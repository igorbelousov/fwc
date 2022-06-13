package handlers

import (
	"fmt"

	"github.com/igorbelousov/foundation/database"
	routing "github.com/qiangxue/fasthttp-routing"
)

func Api(db *database.Database) *routing.Router {
	router := routing.New()

	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello, world!")
		return nil
	})

	return router
}
