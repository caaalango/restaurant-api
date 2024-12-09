package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "github.com/calango-productions/api/cmd/http"
	"github.com/calango-productions/api/internal/adapters"
	"github.com/calango-productions/api/internal/adapters/config"
	"github.com/calango-productions/api/internal/adapters/connections"
	"github.com/calango-productions/api/internal/controllers/docsctl"
	"github.com/calango-productions/api/internal/controllers/healthy"
	"github.com/calango-productions/api/internal/controllers/menuctl"
	"github.com/calango-productions/api/internal/controllers/userctl"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.New()

	conn := connections.New()
	conn.ConnectCoreDatabase(conf)
	conn.ConnectRedis(conf)

	apt := adapters.New(conn)
	app := server.New(apt, conf)

	store := cookie.NewStore([]byte("secret"))
	app.Router.Use(sessions.Sessions("session", store))

	app.Router.LoadHTMLGlob("web/*.html")

	app.Router.Use(apt.Middlewares.CorsMiddleware.Execute())
	app.Router.Use(gin.Logger())
	app.Router.Use(gin.Recovery())

	app.RegisterController(
		docsctl.New(apt),
		healthy.New(apt),
		userctl.New(apt),
		menuctl.New(apt),
	)

	stopCh := setupSignalHandler()

	if err := app.Run(); err != nil {
		fmt.Printf("Application run error: %v\n", err)
		stopCh <- syscall.SIGTERM
	}

	<-stopCh

	shutdownApp(app, conn)
}

func setupSignalHandler() chan os.Signal {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	return stopCh
}

func shutdownApp(app *server.App, conn *connections.Connections) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn.Shutdown(ctx)
	app.Shutdown(ctx)

	err := app.Router.Run(app.Server.Addr)
	if err != nil {
		panic(fmt.Sprintf("Unable to listen server: %v", err))
	}
}
