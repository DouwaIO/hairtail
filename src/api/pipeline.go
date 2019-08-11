package server

import (
	"github.com/gin-gonic/gin"

	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/store"
)

func GetPipelines(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	pipelines, err := store.FromContext(c).GetPipelines(search)
	if err != nil {
		c.JSON(400, gin.H{"message": "get pipelines error", "error": err})
		return
	}

	c.JSON(200, pipelines)
}

func CreatePipeline(c *gin.Context) {
	var pipeline model.Pipeline
	if err := c.ShouldBindJSON(&pipeline); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := store.FromContext(c).CreatePipeline(&pipeline)
	if err != nil {
		c.JSON(400, gin.H{"message": "create pipeline error", "error": err})
		return
	}

	c.JSON(200, pipeline)
}

func GetPipeline(c *gin.Context) {
	pipeline_id := c.Query("pipeline_id")

	pipeline, err := store.FromContext(c).GetPipeline(pipeline_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get pipeline error", "error": err})
		return
	}

	c.JSON(200, pipeline)
}

func UpdatePipeline(c *gin.Context) {
	var pipeline model.Pipeline
	if err := c.ShouldBindJSON(&pipeline); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := store.FromContext(c).UpdatePipeline(&pipeline)
	if err != nil {
		c.JSON(400, gin.H{"message": "create pipeline error", "error": err})
		return
	}

	// c.JSON(200, pipeline)
	c.JSON(200, gin.H{"message": "updated"})
}

func DeletePipeline(c *gin.Context) {
	pipeline_id := c.Query("pipeline_id")

	err := store.FromContext(c).DeletePipeline(pipeline_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get pipeline error", "error": err})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

func OpenPipeline(c *gin.Context) {
	pipeline_id := c.Query("pipeline_id")

	pipeline, err := store.FromContext(c).GetPipeline(pipeline_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get pipeline error", "error": err})
		return
	}
    pipeline.Activate = 1

	err = store.FromContext(c).UpdatePipeline(pipeline)
	if err != nil {
		c.JSON(400, gin.H{"message": "create pipeline error", "error": err})
		return
	}

	// c.JSON(200, pipeline)
	c.JSON(200, gin.H{"message": "opened"})
}

func ClosePipeline(c *gin.Context) {
	pipeline_id := c.Query("pipeline_id")

	pipeline, err := store.FromContext(c).GetPipeline(pipeline_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get pipeline error", "error": err})
		return
	}
    pipeline.Activate = 0

	err = store.FromContext(c).UpdatePipeline(pipeline)
	if err != nil {
		c.JSON(400, gin.H{"message": "create pipeline error", "error": err})
		return
	}

	// c.JSON(200, pipeline)
	c.JSON(200, gin.H{"message": "closed"})
}
