package account

import (
	"../../managers/rpcmanager/account"
	"../../params/authparams"
	"../../settings"
	"../../utils/regexp"
	"../base"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SendPhoneCode(c *gin.Context) {
	phone := c.PostForm("phone")
	if !regexp.RegexpPhone(phone) {
		base.ServeError(c, "params error", errors.New("params error"))
	}
	secret := &authparams.Params{
		Account:     phone,
		AccountType: "phone",
	}
	var res = &authparams.Params{}
	err := account.SendCode(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

func GetProcess(c *gin.Context) {
	u, b := c.Get("uid")
	if !b {
		base.ServeError(c, "token info error", errors.New("token info error"))
		return
	}
	uid := u.(int64)
	secret := &authparams.Params{
		Uid: uid,
	}
	res := &authparams.Params{}
	err := account.GetAuthProcess(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state":    "success",
		"username": res.Process["username"],
		"password": res.Process["password"],
		"stuid":    res.Process["stuid"],
		"phone":    res.Process["phone"],
		"wxopenid": res.Process["wxopenid"],
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if !regexp.RegexpUsername(username) || !regexp.RegexpPassword(password) {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	secret := &authparams.Params{
		Account:     username,
		AccountType: "username",
		Code:        password,
		CodeType:    "password",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
	err := account.AuthAndGetToken(c, secret, res)
	if err != nil {
		return
	}
	c.SetCookie("token", res.Token, 1036800, "/", settings.Domain, http.SameSiteLaxMode, false, true)
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
	secret := &authparams.Params{
		Account:     stuid,
		AccountType: "stuid",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
	err := account.AuthAndGetToken(c, secret, res)
	if err != nil {
		return
	}
	c.SetCookie("token", res.Token, 1036800, "/", settings.Domain, http.SameSiteLaxMode, false, true)
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
		secret := &authparams.Params{
			Token: token,
		}
		// NOTE! res MUST BE INSTANTIATION!
		var res = &authparams.Params{}
		err = account.AuthToken(c, secret, res)
		if err != nil {
			c.Set("token", false)
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

func ChangePassword(c *gin.Context) {
	u, b := c.Get("uid")
	if !b {
		base.ServeError(c, "token info error", errors.New("token info error"))
		return
	}
	token := c.PostForm("token")
	newPassword := c.PostForm("password")
	if !regexp.RegexpPassword(newPassword) {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	uid := u.(int64)
	secret := &authparams.Params{
		Uid:      uid,
		Code:     newPassword,
		CodeType: "password",
		Token:    token,
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
	err := account.ChangeAuth(c, secret, res)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

func GetPromiseByPassword(c *gin.Context) {
	u, b := c.Get("uid")
	if !b {
		base.ServeError(c, "token info error", errors.New("token info error"))
		return
	}
	password := c.PostForm("password")
	if !regexp.RegexpPassword(password) {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	uid := u.(int64)
	secret := &authparams.Params{
		Account:     strconv.FormatInt(uid, 10),
		AccountType: "uid",
		Code:        password,
		CodeType:    "password",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
	err := account.Auth(c, secret, res)
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
	secret := &authparams.Params{
		Account:     phone,
		AccountType: "phone",
		CodeType:    "code",
		Token:       code,
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
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
	secret := &authparams.Params{
		Account:     stuid,
		AccountType: "stuid",
	}
	// NOTE! res MUST BE INSTANTIATION!
	var res = &authparams.Params{}
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
	c.SetCookie("token", "", 0, "/", settings.Domain, http.SameSiteLaxMode, false, true)
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}
