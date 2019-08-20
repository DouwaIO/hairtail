package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/DouwaIO/hairtail/src/router"
	"github.com/DouwaIO/hairtail/src/router/middleware"
	"github.com/DouwaIO/hairtail/src/store/datastore"
	"github.com/DouwaIO/hairtail/src/service"
	yaml "github.com/DouwaIO/hairtail/src/yaml/pipeline"
	"github.com/DouwaIO/hairtail/src/model"
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
			Usage:  "database url",
		},
		cli.StringFlag{
			EnvVar: "HTAIL_TARGET_DB_URL",
			Name:   "target-db-url",
			Usage:  "target database url",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// debug level if requested by user
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	store_ := datastore.New(
		c.String("db-url"),
	)

	handler := router.Load(
		middleware.Store(c, store_),
		//middleware.Task(c, store_),
	)

	targetDB, err := gorm.Open("postgres", c.String("target-db-url"))
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Target db connect failed")
		return err
	}
	defer targetDB.Close()
	log.Info("target database connected")

	err = targetDB.AutoMigrate(&model.RemoteData{},).Error
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("db auto migrate error")
		return err
	}

	//启动数据库里面的service
	go func() {
		pipelines, _ := store_.GetPipelines("")
		for _, pl := range pipelines {
			parsed, err := yaml.ParseString(pl.Config)
			if err != nil {
				log.WithFields(log.Fields{"error": err, "name": pl.Name}).Debug("pipeline parse error")
				continue
			}
			log.WithFields(log.Fields{"name": pl.Name}).Debug("start pipeline")

			for _, s := range parsed.Services {
				// log.WithFields(log.Fields{"name": s.Name}).Debug("run service")
				// forever := make(chan bool)
				// for i := 1; i <= 30; i++ {
				// for i := 1; i <= 1; i++ {
					go func(s *yaml.Task) {
						svc := service.Service{
							Name:     s.Name,
							Desc:     s.Desc,
							Type:     s.Type,
							Settings: s.Settings,
							Steps:    parsed.Steps,
							Store:    &store_,
							TargetDB: targetDB,
						}
						svc.Run()
					}(s)
				// }
				// <-forever
			}
		}
	}()

	return http.ListenAndServe(
		c.String("server-addr"),
		handler,
	)
}
