package main

import (
	"log"
	"os"
	"restsample/config"
	"restsample/instance"
	"restsample/runner"

	"github.com/urfave/cli/v2"
)

func main() {

	config.Load()   // intialize all the configuration
	instance.Init() // intialize the databse

	app := cli.NewApp()
	app.Name = "Rest-Application"
	app.Usage = "Applications to start service"
	app.Commands = []*cli.Command{
		{
			Name:        "start",
			Usage:       "Start the main app",
			Description: "To start the server",
			Action: func(c *cli.Context) error {
				runner.Start()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
