package account

import (
	"../../managers/rpcmanager/account"
	"../../params/authparams"
	"../base"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendPhoneCode(c *gin.Context) {
	phone := c.PostForm("phone")
	if phone == "" {
		base.ServeError(c, "phone number empty", errors.New("phone number empty"))
	}
	secret := &authparams.AuthSecret{
		Account:     phone,
		AccountType: "phone",
	}
	var res = &authparams.ResWithoutToken{}
	err := account.SendCode(c, secret, res)
	if err != nil {
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
	secret := &authparams.AuthSecret{
		Account:     username,
		AccountType: "username",
		Code:        password,
		CodeType:    "password",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.ResWithToken{}
	err := account.AuthAndGetToken(c, secret, res)
	if err != nil {
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
	secret := &authparams.AuthSecret{
		Code:     token,
		CodeType: "token",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.ResWithToken{}
	err := account.AuthToken(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"token": res.Token,
	})
}

func RegisterByPhone(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	if code == "" || phone == "" {
		base.ServeError(c, "empty params access", errors.New("empty params access"))
		return
	}
	secret := &authparams.AuthSecret{
		Account:     phone,
		AccountType: "phone",
		Code:        code,
		CodeType:    "code",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.ResWithToken{}
	err := account.Register(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"token": res.Token,
	})
}
