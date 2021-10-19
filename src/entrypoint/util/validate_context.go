package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ValidateContext(context echo.Context, domain interface{}) error {
	if err := context.Bind(domain); err != nil {
		return err
	}
	if err := context.Validate(domain); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
