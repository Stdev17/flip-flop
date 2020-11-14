package logic

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m Metric) getEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}