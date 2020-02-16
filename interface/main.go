package main

import (
	"../managers/rpcmanager"
	"./account"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var prefix string = "/zsy"

func init() {
	conn, err := rpcmanager.Get()
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
	r := router.Group(prefix)
	{
		r.POST("/auth/getcode", account.SendPhoneCode)
		r.POST("/auth/login", account.Login)
		r.POST("/auth/token", account.AuthToken)
	}
	router.Run(":8000")
}
