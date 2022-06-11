package main

import (
	"fmt"

	"github.com/igorbelousov/fwc/cmd/handlers"
	"github.com/igorbelousov/fwc/foundation/web"
)

func main() {

	log, err := web.InitLog("fast api core")
	if err != nil {
		fmt.Errorf("not loger init", err)
	}

	app := web.NewApp(handlers.Api(), log)
	app.Run()
}
