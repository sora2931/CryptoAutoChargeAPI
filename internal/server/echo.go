package server

import (
	"coin/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/api/new", handlers.NewTransaction)
	e.GET("/api/check", handlers.CheckTransaction)

	return e
}