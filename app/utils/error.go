package utils

import (
	"github.com/labstack/echo/v4"
)

// ErrorCode godoc
// Definition of errorcode
type ErrorCode struct {
	Name    string `json:"name"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Remark  string `json:"remark"`
}

func Fail(c echo.Context, statusCode int, errorObject *ErrorCode, remark string) error {
	errorObject.Remark = remark
	return c.JSON(statusCode, errorObject)
}
