package cabinet

import (
	"../base"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Message struct {
	Images   []map[string]interface{}
	Tooltips map[string]interface{}
}

var lastTime time.Time
var str string

func GetBingUrl(c *gin.Context) {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	if !lastTime.After(n) {
		getBingUrl()
	}
	if str == "" {
		base.ServeError(c, "url is nil", errors.New("url is nil"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"site":  str,
	})
}

func getBingUrl() {
	resp, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN")
	if err != nil {
		panic(err)
		return
	}
	r := resp.Body
	if r == nil {
		panic("r is nil")
		return
	}
	j := json.NewDecoder(r)
	var m Message
	err = j.Decode(&m)
	if err != nil {
		panic(err)
		return
	}
	s := m.Images[0]["url"].(string)
	str = "https://cn.bing.com" + s
	lastTime = time.Now()
}
