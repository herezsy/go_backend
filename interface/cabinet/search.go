package cabinet

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToSearch(c *gin.Context) {
	word := c.Query("q")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"query": word,
	})
}
