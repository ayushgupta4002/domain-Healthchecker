package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ayushgupta4002/check"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "web-healthchecker",
		Usage: "your domain healthchecker",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "domain for checking the health",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "port for checking the health",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := check.Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
