package server

import (
	"github.com/gin-gonic/gin"

	"github.com/DouwaIO/hairtail/src/model"
	"github.com/DouwaIO/hairtail/src/store"
)

func GetWorkflows(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	workflows, err := store.FromContext(c).GetWorkflows(search)
	if err != nil {
		c.JSON(400, gin.H{"message": "get workflows error", "error": err})
		return
	}

	c.JSON(200, workflows)
}

func CreateWorkflow(c *gin.Context) {
	var workflow model.Workflow
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := store.FromContext(c).CreateWorkflow(&workflow)
	if err != nil {
		c.JSON(400, gin.H{"message": "create workflow error", "error": err})
		return
	}

	c.JSON(200, workflow)
}

func GetWorkflow(c *gin.Context) {
	workflow_id := c.Query("workflow_id")

	workflow, err := store.FromContext(c).GetWorkflow(workflow_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get workflow error", "error": err})
		return
	}

	c.JSON(200, workflow)
}

func UpdateWorkflow(c *gin.Context) {
	var workflow model.Workflow
	if err := c.ShouldBindJSON(&workflow); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := store.FromContext(c).UpdateWorkflow(&workflow)
	if err != nil {
		c.JSON(400, gin.H{"message": "create workflow error", "error": err})
		return
	}

	// c.JSON(200, workflow)
	c.JSON(200, gin.H{"message": "updated"})
}

func DeleteWorkflow(c *gin.Context) {
	workflow_id := c.Query("workflow_id")

	err := store.FromContext(c).DeleteWorkflow(workflow_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get workflow error", "error": err})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

func OpenWorkflow(c *gin.Context) {
	workflow_id := c.Query("workflow_id")

	workflow, err := store.FromContext(c).GetWorkflow(workflow_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get workflow error", "error": err})
		return
	}
    workflow.Activate = 1

	err = store.FromContext(c).UpdateWorkflow(workflow)
	if err != nil {
		c.JSON(400, gin.H{"message": "create workflow error", "error": err})
		return
	}

	// c.JSON(200, workflow)
	c.JSON(200, gin.H{"message": "opened"})
}

func CloseWorkflow(c *gin.Context) {
	workflow_id := c.Query("workflow_id")

	workflow, err := store.FromContext(c).GetWorkflow(workflow_id)
	if err != nil {
		c.JSON(400, gin.H{"message": "get workflow error", "error": err})
		return
	}
    workflow.Activate = 0

	err = store.FromContext(c).UpdateWorkflow(workflow)
	if err != nil {
		c.JSON(400, gin.H{"message": "create workflow error", "error": err})
		return
	}

	// c.JSON(200, workflow)
	c.JSON(200, gin.H{"message": "closed"})
}
