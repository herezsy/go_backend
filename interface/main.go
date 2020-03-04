package main

import (
	"../managers/rpcmanager"
	"../settings"
	"./account"
	"./base"
	"./cabinet"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	conn, err := rpcmanager.Get("account")
	if err != nil {
		log.Info("Can not connect to account system")
		return
	}
	defer conn.Close()
	var a, b string
	a = "hey"
	err = conn.Call("Account.Echo", a, &b)
	if err != nil {
		log.Info("Can not echo to account system")
		return
	}
	if a != b {
		log.Info("Can not correct to account system")
		return
	}
	log.WithFields(log.Fields{
		"echo": b,
	}).Info()
}

func main() {
	router := gin.Default()
	// Only load all files at once or there will be errors of undefine.
	router.LoadHTMLFiles("./static/template/search.tmpl", "./static/template/list.tmpl")
	r := router.Group(settings.Prefix, account.AuthTokenNotReject)
	{
		r.POST("/auth/getcode", account.SendPhoneCode)
		r.POST("/auth/loginbypassword", account.LoginByPassword)
		r.POST("/auth/loginbycode", account.LoginByCode)
		r.POST("/auth/token", account.OpenAuthToken)
		r.POST("/auth/register", account.RegisterByPhone)
		r.POST("/auth/getnickname", account.GetNicknameAndProcess)
		r.POST("/auth/stuid", account.LoginByStuid)
		r.POST("/auth/getprocess", account.GetProcess)
		r.POST("/auth/logout", account.Logout)
		r.POST("/auth/changepassword", account.ChangePassword)
		r.POST("/auth/getpromise", account.GetPromiseByPassword)
		r.Any("/echo", base.Echo)
	}

	c := router.Group("/c")
	{
		c.GET("/wallpaper", cabinet.GetBingUrl)

		c.GET("/search", cabinet.ToSearch)
		c.GET("/search/list", cabinet.GetRecord)
		c.POST("/search/update", cabinet.UpdateSearch)
		c.POST("/search/delete", cabinet.DeleteRecord)
	}
	// This approach will report an error about MIME type if server is running on Windows
	// which is because that FileSystem type is different.
	// Usually, static resources deploy by Nginx rather than Go process.
	// router.Static("/qndxx", "./static")
	router.Run(settings.PortInterface)
}
