package param_valid

import (
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
)

func Check(params []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, param := range params {
			if c.DefaultQuery(param, "") == "" {
				c.JSON(400, gin.H{"message": "query param (" + param + ") is invalid"})
				c.Abort()
				return
			}
		}
	}
}
