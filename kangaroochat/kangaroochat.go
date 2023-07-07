package main

import (
	"context"
	"fmt"
	"html/template"

	"github.com/golang-jwt/jwt/v5"
	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/domain/entity"
	"github.com/howood/kangaroochat/infrastructure/custommiddleware"
	"github.com/howood/kangaroochat/interfaces/service/handler"
	"github.com/howood/kangaroochat/library/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// DefaultPort is default port of server
var DefaultPort = utils.GetOsEnv("SERVER_PORT", "8080")

func main() {
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

	jwtconfig := echojwt.Config{
		Skipper: custommiddleware.OptionsMethodSkipper,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(entity.JwtClaims)
		},
		SigningKey:  []byte(actor.TokenSecret),
		TokenLookup: "cookie:" + actor.RoomTokenKey,
		ContextKey:  actor.JWTContextKey,
	}
	e.GET("/websocket/:identifier", handler.WebSockerHandler{BroadCaster: braodcaster}.Request, echojwt.WithConfig(jwtconfig))
	e.GET("/client/:identifier", handler.ClientHandler{}.Request, echojwt.WithConfig(jwtconfig))

	go braodcaster.BroadcastMessages()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", DefaultPort)))

}
