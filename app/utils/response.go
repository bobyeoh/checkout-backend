package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
