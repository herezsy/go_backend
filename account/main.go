package main

import (
	"./auth"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/rpc"
)

var port string = ":7744"

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
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"action": "http.ListenAndServe",
			"Error":  err,
		}).Panic()
	}
}
