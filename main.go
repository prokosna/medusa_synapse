package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fgrosse/goldi"
	"github.com/labstack/echo"
	"github.com/prokosna/medusa_synapse/app"
	"github.com/prokosna/medusa_synapse/route"
	"github.com/urfave/cli"
)

var (
	port    uint
	brokers []string
)

func main() {
	ap := cli.NewApp()
	ap.Name = "Medusa Synapse"
	ap.Flags = []cli.Flag{
		cli.UintFlag{
			Name:        "port, P",
			Value:       8080,
			Usage:       "Running port",
			EnvVar:      "MS_PORT",
			Destination: &port,
		},
		cli.StringSliceFlag{
			Name:   "brokers, B",
			Value:  &cli.StringSlice{"http://localhost:8000"},
			Usage:  "Brokers list of Kafka",
			EnvVar: "MS_BROKERS",
		},
	}
	ap.Action = func(c *cli.Context) error {
		// Config
		brokers = c.StringSlice("brokers")

		// DI
		registry := goldi.NewTypeRegistry()
		config := map[string]interface{}{}
		container := goldi.NewContainer(registry, config)
		container.RegisterType("Medusa", app.NewMedusa)
		container.RegisterType("RootRoute", route.NewRootRoute)
		container.RegisterType("MedusaRoute", route.NewMedusaRoute, "@Medusa")

		// Server routing
		e := echo.New()
		root := container.MustGet("RootRoute").(*route.RootRoute)
		root.InitRoutes(e)
		mg := e.Group("/medusa")
		medusa := container.MustGet("MedusaRoute").(*route.MedusaRoute)
		medusa.InitRoutes(mg)

		// Server start
		return e.Start(":" + fmt.Sprint(port))
	}
	err := ap.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
