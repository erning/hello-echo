package main

import (
	"github.com/erning/hello-echo/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	g := e.Group("/api")
	g.GET("/misc.GetServerTime", api.GetServerTime)

	e.Logger.Fatal(e.Start(":5000"))
}
