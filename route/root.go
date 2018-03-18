package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type RootRoute struct{}

func NewRootRoute() *RootRoute {
	return &RootRoute{}
}
func (r *RootRoute) InitRoutes(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
}
