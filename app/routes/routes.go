package routes

import (
	"os"

	"checkout-backend/app"
	"checkout-backend/app/controller/number"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes godoc
func Routes(server *app.Server) {
	webV1Prefix := "/checkout/v1/"
	numberController := number.Init(server)
	server.Echo.Use(middleware.Logger())
	if os.Getenv("ENV") != "production" {
		server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
		server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
			AllowOrigins:     []string{"*"},
		}))
	}
	server.Echo.PUT(webV1Prefix+"number", numberController.Update)
	server.Echo.GET(webV1Prefix+"number", numberController.GetNumber)

}
