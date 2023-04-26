package errorCode

import "checkout-backend/app/utils"

var (
	// StatusInternalServerError godoc
	StatusInternalServerError = &utils.ErrorCode{Name: "StatusInternalServerError", Code: 10001, Message: "Sorry, something went wrong. Please try again later."}
)
