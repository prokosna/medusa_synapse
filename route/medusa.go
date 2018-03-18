package route

import (
	"github.com/labstack/echo"
	"github.com/prokosna/medusa_synapse/app"
)

type MedusaRoute struct {
	medusa *app.Medusa
}

func NewMedusaRoute(medusa *app.Medusa) *MedusaRoute {
	return &MedusaRoute{
		medusa: medusa,
	}
}

func (r *MedusaRoute) InitRoutes(g *echo.Group) {
	return
}
