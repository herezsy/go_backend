package account

import (
	"../../../interface/base"
	"../../../managers/rpcmanager"
	"../../../params/authparams"
	"github.com/gin-gonic/gin"
	"net/rpc"
)

func SendCode(c *gin.Context, secret *authparams.AuthSecret, res *authparams.ResWithoutToken) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnSendCode(c, conn, secret, res)
	return
}

func ConnSendCode(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, res *authparams.ResWithoutToken) (err error) {
	err = conn.Call("Account.SendCode", secret, res)
	if err != nil {
		base.ServeError(c, "Account.SendCode", err)
	}
	return
}

func AuthAndGetToken(c *gin.Context, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnAuthAndGetToken(c, conn, secret, res)
	return
}

func ConnAuthAndGetToken(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	err = conn.Call("Account.AuthAndGetToken", &secret, &res)
	if err != nil {
		base.ServeError(c, "Account.AuthAndGetToken", err)
	}
	return
}

func Auth(c *gin.Context, secret *authparams.AuthSecret, res *authparams.ResWithoutToken) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnAuth(c, conn, secret, res)
	return
}

func ConnAuth(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, res *authparams.ResWithoutToken) (err error) {
	err = conn.Call("Account.Auth", &secret, &res)
	if err != nil {
		base.ServeError(c, "Account.Auth", err)
	}
	return
}

func AuthToken(c *gin.Context, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnAuthToken(c, conn, secret, res)
	return
}

func ConnAuthToken(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	err = conn.Call("Account.AuthToken", &secret, &res)
	if err != nil {
		base.ServeError(c, "Account.AuthToken", err)
	}
	return
}

func Register(c *gin.Context, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnRegister(c, conn, secret, res)
	return
}

func ConnRegister(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, res *authparams.ResWithToken) (err error) {
	err = conn.Call("Account.Register", &secret, &res)
	if err != nil {
		base.ServeError(c, "Account.Register", err)
	}
	return
}

func FindUid(c *gin.Context, secret *authparams.AuthSecret, b *int64) (err error) {
	conn, err := rpcmanager.Get("account")
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	err = ConnFindUid(c, conn, secret, b)
	return
}

func ConnFindUid(c *gin.Context, conn *rpc.Client, secret *authparams.AuthSecret, b *int64) (err error) {
	err = conn.Call("Account.FindUid", &secret, &b)
	if err != nil {
		base.ServeError(c, "Account.FindUid", err)
	}
	return
}
