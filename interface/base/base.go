package base

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var last = time.Now()

func Echo(c *gin.Context) {
	ip := c.ClientIP()
	method := c.Request.Method
	p := c.Query("p")
	f := c.PostForm("f")
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"state":   "success",
		"ip":      ip,
		"method":  method,
		"time":    t,
		"last":    last,
		"name":    "it is interface.",
		"path(p)": p,
		"form(f)": f,
	})
	last = t
}

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
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
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
