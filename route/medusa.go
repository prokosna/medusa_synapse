package route

import (
	"github.com/labstack/echo"
	"github.com/prokosna/medusa_synapse/app"
	"github.com/prokosna/medusa_synapse/domain"
	"github.com/prokosna/medusa_synapse/exception"
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
	g.POST("/:key/frames", r.sendFrame)
	return
}

func (r *MedusaRoute) sendFrame(c echo.Context) error {
	key := c.Param("key")
	var img domain.Image
	err := c.Bind(&img)
	if err != nil {
		return exception.NewBadRequestError(err.Error())
	}
	err = r.medusa.SendFrame(key, img)
	if err != nil {
		return err
	}
	return c.String(200, "")
}
