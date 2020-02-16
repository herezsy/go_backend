package account

import (
	"../../managers/rpcmanager"
	"../../params/authparams"
	"../base"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SendPhoneCode(c *gin.Context) {
	phone := c.PostForm("phone")
	if phone == "" {
		base.ServeError(c, "phone number empty", errors.New("phone number empty"))
	}
	conn, err := rpcmanager.Get()
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
		return
	}
	secret := authparams.AuthSecret{
		Account:     phone,
		AccountType: "phone",
	}
	var res = authparams.ResWithoutToken{}
	err = conn.Call("Account.SendCode", secret, &res)
	if err != nil {
		base.ServeError(c, "Account.SendCode", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		base.ServeError(c, "phone number empty", errors.New("phone number empty"))
		return
	}
	conn, err := rpcmanager.Get()
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
	}
	secret := authparams.AuthSecret{
		Account:     username,
		AccountType: "username",
		Code:        password,
		CodeType:    "password",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = authparams.ResWithToken{}
	err = conn.Call("Account.AuthAndGetToken", secret, &res)
	log.Info(res)
	if err != nil {
		base.ServeError(c, "Account.AuthAndGetToken", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"token": res.Token,
	})
}

func AuthToken(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		base.ServeError(c, "none token access", errors.New("none token access"))
		return
	}
	conn, err := rpcmanager.Get()
	defer conn.Close()
	if err != nil {
		base.ServeFatal(c, "rpcmanager.Get", err)
	}
	secret := authparams.AuthSecret{
		Code:     token,
		CodeType: "token",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = authparams.ResWithToken{}
	err = conn.Call("Account.AuthToken", secret, &res)
	log.Info(res)
	if err != nil {
		base.ServeError(c, "Account.AuthToken", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"token": res.Token,
	})
}
