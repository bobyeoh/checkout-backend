package number

import (
	"checkout-backend/app"
	"checkout-backend/app/controller"
	"checkout-backend/app/controller/number/errorCode"
	"checkout-backend/app/controller/number/request"
	"checkout-backend/app/controller/number/response"
	"checkout-backend/app/repository"
	"checkout-backend/app/service"
	"checkout-backend/app/utils"
	"checkout-backend/app/validation"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NumberHandler godoc
type NumberHandler struct {
	server *app.Server
}

// Init godoc
func Init(server *app.Server) *NumberHandler {
	return &NumberHandler{server: server}
}

// Get Number godoc
// @Tags Number
// @Summary Collect backtest order from bot
// @Description
// @Success 200
// @Failure 400 {object} utils.ErrorCode
// @Failure 500 {object} utils.ErrorCode
// @Router /checkout/v1/number [get]
func (handler *NumberHandler) GetNumber(c echo.Context) error {
	log := c.Logger()

	numberRepo := repository.InitNumber(handler.server.MySQL)
	count, err := numberRepo.GetNumber()
	resp := response.NumberResult{Count: count}
	if err != nil {
		log.Errorf("Failed to retrive number: %v", err.Error())
		utils.Fail(c, http.StatusInternalServerError, errorCode.StatusInternalServerError, "")
	}

	return utils.Success(c, resp)
}

// Update godoc
// @Tags Number
// @Summary Update count
// @Description Update count
// @Param params body request.Number true "Number"
// @Success 200
// @Failure 400 {object} utils.ErrorCode
// @Failure 500 {object} utils.ErrorCode
// @Router /checkout/v1/number [put]
func (handler *NumberHandler) Update(c echo.Context) error {
	log := c.Logger()
	requestBody := new(request.Number)
	if err := c.Bind(requestBody); err != nil {
		return err
	}
	if err := c.Validate(requestBody); err != nil {
		errorMessage := validation.ProcessError(err)
		return utils.Fail(c, http.StatusBadRequest, controller.ParametersError, errorMessage.Error())
	}
	j, _ := json.Marshal(requestBody)
	log.Info(string(j))
	numberService := service.InitNumber(handler.server.MySQL)
	if err := numberService.Update(requestBody.Count); err != nil {
		log.Errorf("Failed to update: %v", err.Error())
		return utils.Fail(c, http.StatusInternalServerError, errorCode.StatusInternalServerError, "update count failed")
	}
	return utils.Success(c, nil)
}
