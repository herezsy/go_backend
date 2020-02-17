package rpcmanager

import (
	"errors"
	"net/rpc"
)

var accountPort = ":7744"

func Get(sys string) (client *rpc.Client, err error) {
	switch sys {
	case "account":
		client, err = rpc.DialHTTP("tcp", "127.0.0.1"+accountPort)
	default:
		client, err = nil, errors.New("system not exist")
	}
	return
}
