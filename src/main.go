package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"

	"github.com/DouwaIO/hairtail/src/router"
	"github.com/DouwaIO/hairtail/src/router/middleware"
	"github.com/DouwaIO/hairtail/src/service"
	"github.com/DouwaIO/hairtail/src/store/datastore"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
)

func main() {
	app := cli.NewApp()
	app.Name = "hairtail"
	// app.Version = version.Version.String()
	app.Version = "0.1"
	app.Action = run
	// app.Before = before
	app.Flags = []cli.Flag{
		cli.BoolTFlag{
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
		cli.StringFlag{
			EnvVar: "HTAIL_DB_URL",
			Name:   "db-url",
			Usage:  "server address",
			Value:  "host=47.110.154.127 port=30172 user=postgres dbname=hairtail sslmode=disable password=huansi@2017",
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
        log.SetLevel(log.DebugLevel)
    }

	store_ := datastore.New(
		c.String("db-url"),
	)
	// pipeline.Queue = model.WithTaskStore(queue.New(), store_)

	handler := router.Load(
		middleware.Store(c, store_),
		//middleware.Task(c, store_),
	)

	//启动数据库里面的service
	pipelines, _ := store_.GetPipelines("")
	for _, pl := range pipelines {
		log.Debugf("start pipeline: %s", pl.Name)
		parsed, err := yaml.ParseString(pl.Config)
		if err != nil {
			return nil
		}

		for _, s := range parsed.Services {
			log.Debugf("run service:", s.Name)

			svc := service.Service{
				Name:     s.Name,
				Desc:     s.Desc,
				Type:     s.Type,
				Settings: s.Settings,
				Steps:    parsed.Steps,
				Store:    &store_,
			}
			err := svc.Run()
			if err != nil {
				return nil
			}
		}
	}

	return http.ListenAndServe(
		c.String("server-addr"),
		handler,
	)
}
