package rpcmanager

import (
	"../../settings"
	"errors"
	"net/rpc"
)

func Get(sys string) (client *rpc.Client, err error) {
	switch sys {
	case "account":
		client, err = rpc.DialHTTP("tcp", settings.Domain+settings.PortAccount)
	default:
		client, err = nil, errors.New("system not exist")
	}
	return
}
