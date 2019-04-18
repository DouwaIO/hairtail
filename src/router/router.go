package router

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/DouwaIO/hairtail/src/server"
	"github.com/DouwaIO/hairtail/src/router/middleware/ginrus"
	"github.com/DouwaIO/hairtail/src/router/middleware/header"
)

// Load loads the router
func Load(middleware ...gin.HandlerFunc) http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	e.Use(header.NoCache)
	e.Use(header.Options)
	e.Use(header.Secure)
	e.Use(middleware...)

	e.POST("/api/schema", server.Schema)
	e.POST("/api/pipeline", server.Pipeline)
	e.POST("/api/data", server.PostData)

	return e
}
