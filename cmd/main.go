package main

import (
	"fmt"
	"os"
	"time"

	"github.com/igorbelousov/fwc/internal/handlers"

	"github.com/ardanlabs/conf"
	"github.com/igorbelousov/fwc/foundation/web"
	"github.com/igorbelousov/fwc/internal/database"
)

var build = "develop"

func main() {
	var cfg struct {
		conf.Version
		Web struct {
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:5s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
		Auth struct {
			KeyID          string `conf:"default:54bb2165-71e1-41a6-af3e-7da4a0e1e2c1"`
			PrivateKeyFile string `conf:"default:./private.pem"`
			Algorithm      string `conf:"default:RS256"`
		}
		DB struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:0.0.0.0:5432"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:true"`
		}
	}
	if err := conf.Parse(os.Args[1:], "APP", &cfg); err != nil {
		fmt.Printf("Config parse Error %s", err)
		os.Exit(1)
	}
	cfg.Version.SVN = build
	cfg.Version.Desc = "Core for writing web services"

	err := database.Open(database.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		fmt.Println("Connect to db", err)
		os.Exit(1)
	}
	defer func() {
		fmt.Printf("main: Database Stopping : %s", cfg.DB.Host)
		database.DB.Close()
	}()

	log, err := web.InitLog("fast api core")
	if err != nil {
		fmt.Println("loger init false ", err)
	}

	app := web.NewApp(handlers.Api(), log, nil, nil)
	app.Run(cfg.Web.APIHost)
}
