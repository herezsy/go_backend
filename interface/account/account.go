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
	c.SetCookie("token", res.Token, 1036800, "/", "localhost", http.SameSiteLaxMode, false, true)
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"token": res.Token,
	})
}

func LoginByStuid(c *gin.Context) {
	stuid := c.PostForm("stuid")
	if !regexp.RegexpStuid(stuid) {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	secret := &authparams.AuthSecret{
		Account:     stuid,
		AccountType: "stuid",
		Code:        stuid,
		CodeType:    "stuid",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.ResWithToken{}
	err := account.AuthAndGetToken(c, secret, res)
	if err != nil {
		return
	}
	c.SetCookie("token", res.Token, 1036800, "/", "localhost", http.SameSiteLaxMode, false, true)
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

func OpenAuthToken(c *gin.Context) {
	token, _ := c.Get("token")
	if token == true {
		nk, _ := c.Get("nk")
		pt, _ := c.Get("pt")
		uid, _ := c.Get("uid")
		c.JSON(http.StatusOK, gin.H{
			"state": "success",
			"nk":    nk,
			"pt":    pt,
			"uid":   uid,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "error",
		})
	}
}

func AuthTokenNotReject(c *gin.Context) {
	token, err := c.Cookie("token")
	if err == nil && regexp.RegexpToken(token) {
		secret := &authparams.AuthSecret{
			Code:     token,
			CodeType: "token",
		}
		// NOTE! res MUST BE INSTANTIATION!
		var res = &authparams.ResWithToken{}
		err = account.AuthToken(c, secret, res)
		if err != nil {
			return
		}
		c.Set("pt", res.PrivilegeType)
		c.Set("pl", res.PrivilegeLevel)
		c.Set("uid", res.Uid)
		c.Set("nk", res.Nickname)
		c.Set("token", true)
	} else {
		c.Set("token", false)
	}

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

func Logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "localhost", http.SameSiteLaxMode, false, true)
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}
