package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type requestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator() *requestValidator {
	return &requestValidator{validator.New()}
}

func (rv *requestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
