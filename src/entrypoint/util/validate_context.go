package util

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)
type (
CustomValidator struct {
Validator *validator.Validate
})

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}


func ValidateContext(context echo.Context, domain interface{}) error {
	if err := context.Bind(domain); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := context.Validate(domain); err != nil {
		return err
	}
	return nil
}
