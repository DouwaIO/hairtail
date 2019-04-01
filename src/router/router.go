package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load loads the router
func Load() http.Handler {

	e := gin.New()

	// e.Use(header.NoCache)
	// e.Use(header.Options)
	// e.Use(header.Secure)
	// e.Use(middleware...)
	// e.Use(session.SetUser())
	// e.Use(token.Refresh)

	// e.GET("/api/dig", server.Dig)
	// e.POST("/api/test", server.Test)

	// e.GET("/version", server.Version)
	// e.GET("/healthz", server.Health)

	return e
}
