package rpcmanager

import "net/rpc"

var port string = ":7744"

func Get() (client *rpc.Client, err error) {
	client, err = rpc.DialHTTP("tcp", "127.0.0.1"+port)
	return
}
