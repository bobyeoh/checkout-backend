package app

import (
	"checkout-backend/app/validation"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Server godoc
type Server struct {
	Echo  *echo.Echo
	MySQL *gorm.DB
}

// NewServer godoc
func NewServer(mysql *gorm.DB) *Server {
	e := echo.New()
	// e.Logger.SetLevel(log.INFO)
	// e.Logger.SetOutput(os.Stdout)
	e.Validator = validation.Init()
	return &Server{
		Echo:  e,
		MySQL: mysql,
	}
}

// Start web server godoc
func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
