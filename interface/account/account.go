package account

import (
	"../../managers/rpcmanager/account"
	"../../params/authparams"
	"../../utils/regexp"
	"../base"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendPhoneCode(c *gin.Context) {
	phone := c.PostForm("phone")
	if !regexp.RegexpPhone(phone) {
		base.ServeError(c, "params error", errors.New("params error"))
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
	if !regexp.RegexpUsername(username) || !regexp.RegexpPassword(password) {
		base.ServeError(c, "params error", errors.New("params error"))
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
	if !regexp.RegexpToken(token) {
		base.ServeError(c, "params error", errors.New("params error"))
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
	if !regexp.RegexpPhone(phone) || !regexp.RegexpCode(code) {
		base.ServeError(c, "params error", errors.New("params error"))
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

func GetNickname(c *gin.Context) {
	stuid := c.PostForm("stuid")
	if !regexp.RegexpStuid(stuid) {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	secret := &authparams.AuthSecret{
		Account:     stuid,
		AccountType: "stuid",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.ResWithoutToken{}
	err := account.GetNickname(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":    "success",
		"nickname": res.Nickname,
	})
}
