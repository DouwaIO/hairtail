package router

import (
    "net/http"

    // "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"

	"github.com/DouwaIO/hairtail/src/router/middleware/param_valid"
	"github.com/DouwaIO/hairtail/src/api"
)

// Load loads the router
func Load(middleware ...gin.HandlerFunc) http.Handler {
    r := gin.New()

    // Global middleware
    // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
    // By default gin.DefaultWriter = os.Stdout
    r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware...)
    // allow cors
    // r.Use(cors.Default())

    // index
	r.LoadHTMLGlob("views/*")
	r.GET("/", server.Dashboard)

	// r.POST("/api/schema", server.Schema)
	// r.POST("/api/pipeline", server.Pipeline)
	// r.POST("/api/data", server.PostData)
	// r.GET("/api/info", server.Info)
	// r.GET("/api/builds", server.GetBuilds)

    pipelines := r.Group("/api/pipelines/")
	{
		pipelines.GET("", server.GetPipelines)
		pipelines.POST("", server.CreatePipeline)
	}

    pipeline := r.Group("/api/pipeline/")
	pipeline.Use(param_valid.Check([]string{"pipeline_id"}))
	{
		pipeline.GET("", server.GetPipeline)
		pipeline.PUT("", server.UpdatePipeline)
		pipeline.DELETE("", server.DeletePipeline)

		pipeline.PUT("open/", server.OpenPipeline)
		pipeline.PUT("close/", server.ClosePipeline)
	}

	return r
}
