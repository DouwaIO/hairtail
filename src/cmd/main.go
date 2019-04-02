package main

import (
	"fmt"
    "net/http"
	"os"

	"github.com/urfave/cli"

	"github.com/DouwaIO/hairtail/src/router"
)

func main() {
	app := cli.NewApp()
	app.Name = "drone-server"
	// app.Version = version.Version.String()
	app.Version = "0.1"
	app.Action = run
	// app.Before = before
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			EnvVar: "HTAIL_DEBUG",
			Name:   "debug",
			Usage:  "enable server debug mode",
		},
		cli.StringFlag{
			EnvVar: "HTAIL_SERVER_HOST,HTAIL_ADDR",
			Name:   "server-host",
			Usage:  "server fully qualified url (<scheme>://<host>)",
		},
		cli.StringFlag{
			EnvVar: "HTAIL_SERVER_ADDR,HTAIL_ADDR",
			Name:   "server-addr",
			Usage:  "server address",
			Value:  ":8000",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	// debug level if requested by user
	if c.Bool("debug") {
	} else {
	}

	handler := router.Load()

	// start the server without tls enabled
	if !c.Bool("lets-encrypt") {
		return http.ListenAndServe(
			c.String("server-addr"),
			handler,
		)
	}

    return nil
}
