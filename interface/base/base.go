package base

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ServeSuccess(c *gin.Context, h *gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

func ServeError(c *gin.Context, action string, err error) {
	log.WithFields(log.Fields{
		"Action": action,
		"Error":  err,
	}).Warn()
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"state": "error",
		"error": action,
	})
}

func ServeFatal(c *gin.Context, action string, err error) {
	log.WithFields(log.Fields{
		"Action": action,
		"Error":  err,
	}).Warn()
	c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
		"state": "fatal",
		"error": action,
	})
}
