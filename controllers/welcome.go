package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello, world")
}
