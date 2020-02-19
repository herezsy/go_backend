package main

import (
	"../managers/rpcmanager"
	"../settings"
	"./account"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	var a, b string
	a = "hey"
	err = conn.Call("Account.Echo", a, &b)
	if err != nil {
		panic(err)
	}
	if a != b {
		panic(errors.New("echo not equal"))
	}
	log.WithFields(log.Fields{
		"echo": b,
	}).Info()
}

func main() {
	router := gin.Default()
	r := router.Group(settings.Prefix, account.AuthTokenNotReject)
	{
		r.POST("/auth/getcode", account.SendPhoneCode)
		r.POST("/auth/login", account.Login)
		r.POST("/auth/token", account.OpenAuthToken)
		r.POST("/auth/register", account.RegisterByPhone)
		r.POST("/auth/getnickname", account.GetNickname)
		r.POST("/auth/stuid", account.LoginByStuid)
		r.POST("/auth/logout", account.Logout)
	}
	router.Run(settings.PortInterface)
}
