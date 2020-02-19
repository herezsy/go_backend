package main

import (
	"../settings"
	"./auth"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/rpc"
)

func main() {
	account := new(auth.Account)
	err := rpc.Register(account)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "rpc.Register",
			"Error":  err,
		}).Panic()
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(settings.PortAccount, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "http.ListenAndServe",
			"Error":  err,
		}).Panic()
	}
}
