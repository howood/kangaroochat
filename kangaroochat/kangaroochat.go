package main

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/interfaces/service/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// DefaultPort is default port of server
var DefaultPort = "8080"

func main() {
	if os.Getenv("SERVER_PORT") != "" {
		DefaultPort = os.Getenv("SERVER_PORT")
	}
	renderer := &handler.HTMLTemplate{
		Templates: template.Must(handler.LoadTemplate("*.html")),
	}
	braodcaster := actor.NewBroadCastMessanger(context.Background())
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Renderer = renderer

	e.GET("/websocket/:identifier", handler.WebSockerHandler{BroadCaster: braodcaster}.Request)
	e.GET("/client/:identifier", handler.ClientHandler{}.Request)

	go braodcaster.BroadcastMessages()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", DefaultPort)))

}
