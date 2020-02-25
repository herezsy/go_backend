package cabinet

import (
	"../base"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Images   []map[string]interface{}
	Tooltips map[string]interface{}
}

func GetBingUrl(c *gin.Context) {
	resp, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN")
	if err != nil {
		base.ServeError(c, "http.Get", err)
		return
	}
	r := resp.Body
	if r == nil {
		base.ServeError(c, "resp.Body", errors.New("resp.body not exist"))
		return
	}
	j := json.NewDecoder(r)
	var m Message
	err = j.Decode(&m)
	if err != nil {
		base.ServeError(c, "j.Decode", err)
		return
	}
	s := m.Images[0]["url"].(string)
	str := "https://cn.bing.com/" + s
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"site":  str,
	})
}
