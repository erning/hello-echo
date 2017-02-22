package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func GetServerTime(c echo.Context) error {
	now := time.Now().Format("2006-01-02 15:04:05.00 -0700")
	payload := &Payload{
		Status:  "ok",
		Payload: now,
	}
	return c.JSON(http.StatusOK, payload)
}
