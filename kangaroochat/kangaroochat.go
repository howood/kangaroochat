package main

import (
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/infrastructure/custommiddleware"
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
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrftoken",
	}))
	e.Renderer = renderer

	e.GET("/create", handler.AccountHandler{}.CreateGet)
	e.POST("/create", handler.AccountHandler{}.Create)
	e.GET("/login/:identifier", handler.AccountHandler{}.LoginGet)
	e.POST("/login/:identifier", handler.AccountHandler{}.Login)

	jwtconfig := middleware.JWTConfig{
		Skipper:     custommiddleware.OptionsMethodSkipper,
		Claims:      &entity.JwtClaims{},
		SigningKey:  []byte(actor.TokenSecret),
		TokenLookup: "cookie:" + actor.RoomTokenKey,
		ContextKey:  actor.JWTContextKey,
	}
	e.GET("/websocket/:identifier", handler.WebSockerHandler{BroadCaster: braodcaster}.Request, middleware.JWTWithConfig(jwtconfig))
	e.GET("/client/:identifier", handler.ClientHandler{}.Request, middleware.JWTWithConfig(jwtconfig))

	go braodcaster.BroadcastMessages()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", DefaultPort)))

}
