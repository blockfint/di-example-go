package utils

import (
	"path"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func getAbsoluteURL(relativeURL string) string {
	BASE_URL := viper.GetString("HTTP_SERVER.BASE_URL")

	return path.Join(BASE_URL, relativeURL)
}

func SetLocationHeader(c echo.Context, relativeURL string) {
	c.Response().Header().Set("Location", getAbsoluteURL(relativeURL))
}
