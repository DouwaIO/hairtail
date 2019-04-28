package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"

	"github.com/DouwaIO/hairtail/src/router"
	//"github.com/DouwaIO/hairtail/src/store"
	"github.com/DouwaIO/hairtail/src/store/datastore"
	"github.com/DouwaIO/hairtail/src/router/middleware"
	"log"
	task_service "github.com/DouwaIO/hairtail/src/service"
	yaml_pipeline "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/task"
	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/task/queue"
)

func main() {
	app := cli.NewApp()
	app.Name = "hairtail"
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
	//if c.Bool("debug") {
	//} else {
	//}
	store_ := datastore.New(
		c.String("db-url"),
	)
	task.Queue = model.WithTaskStore(queue.New(), store_)

	handler := router.Load(
		middleware.Store(c, store_),
		//middleware.Task(c, store_),
	)

	//启动数据库里面的service
	services, _ := store_.GetServiceAllList()
	log.Printf("Received a message: %s", services)
	for _, service := range services {
		parsed, err := yaml_pipeline.ParseString(service.Data)
		if err != nil {
		       return nil
		}

		if len(parsed.Services) > 0 {
		       for _, service2 := range parsed.Services {
			if service2.Name == service.Name && service2.Type == "MQ" {
				//log.Printf("Received a message: %s", service)
				log.Printf("run service:", service.Name)

				task_service.Service(service2, parsed.Pipeline, service.ID, store_)

			}
		       }
		}
	}

	if !c.Bool("lets-encrypt") {
		return http.ListenAndServe(
			c.String("server-addr"),
			handler,
		)
	}

    return nil
}
