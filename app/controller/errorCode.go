package controller

import "checkout-backend/app/utils"

var (
	// ParametersError godoc
	ParametersError = &utils.ErrorCode{Name: "ParametersError", Code: 10001, Message: "Sorry, the request parameters are invalid or missing. Please check your inputs and try again."}
)
