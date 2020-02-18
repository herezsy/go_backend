package auth

import (
	"../../managers/dbmanager"
	"../../params/authparams"
	"../../utils/randworker"
	"./supports"
	"errors"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Account int

func init() {
	log.SetReportCaller(true)
}

func (account *Account) AuthAndGetToken(secret *authparams.AuthSecret, res *authparams.ResWithToken) error {
	// auth info
	uid, pt, pl, nk, err := auth(secret)
	if err != nil {
		meetError("auth", err)
		return err
	}
	// make token
	token, err := supports.MakeToken(uid, pt, pl, nk)
	if err != nil {
		meetError("makeToken", err)
		return err
	}
	// assign
	res.Uid = uid
	res.Nickname = nk
	res.PrivilegeType = pt
	res.PrivilegeLevel = pl
	res.Token = token
	// NOTE! CAN NOT assign the pointer by another object!
	//res = &authparams.ResWithToken{
	//	Uid:            uid,
	//	Nickname:       nk,
	//	PrivilegeType:  pt,
	//	PrivilegeLevel: pl,
	//	Token:          token,
	//}
	return nil
}

func (account *Account) Auth(secret *authparams.AuthSecret, res *authparams.ResWithoutToken) error {
	// auth info
	uid, pt, pl, nk, err := auth(secret)
	if err != nil {
		meetError("auth", err)
		return err
	}
	// assign
	res.Uid = uid
	res.Nickname = nk
	res.PrivilegeType = pt
	res.PrivilegeLevel = pl
	return nil
}

func (account *Account) AuthToken(secret *authparams.AuthSecret, res *authparams.ResWithToken) error {
	// confirm code type
	if secret.CodeType != "token" {
		return errors.New("codeType wrong")
	}
	// get message from token
	uid, pt, pl, nk, t, err := supports.DecodeToken(secret.Code)
	if err != nil {
		meetError("decodeToken", err)
		return err
	}
	// check whether token is old or not
	var token string
	if tDead := time.Now(); tDead.After(t) {
		return errors.New("token expired")
	} else if tOld := tDead.Add(10 * 24 * time.Hour); tOld.After(t) {
		token, err = supports.MakeToken(uid, pt, pl, nk)
		if err != nil {
			return err
		}
	} else {
		token = secret.Code
	}
	// assign
	res.Uid = uid
	res.Nickname = nk
	res.PrivilegeType = pt
	res.PrivilegeLevel = pl
	res.Token = token
	return nil
}

func (account *Account) Echo(str *string, res *string) error {
	*res = *str
	return nil
}

func (account *Account) SendCode(secret *authparams.AuthSecret, res *authparams.ResWithoutToken) error {
	// confirm account type
	if secret.AccountType != "phone" {
		return errors.New("accountType wrong")
	}
	// get send type
	uid, _, _, _, _, err := getInfo(secret)
	if err != nil && err.Error() != "uid not found" {
		meetError("auth", err)
		return err
	}
	var str string
	if err != nil && err.Error() == "uid not found" {
		// no-register type
		str = "auth&phone=" + secret.Account
	} else {
		// register type
		str = "auth&uid=" + strconv.FormatInt(uid, 10)
	}
	err = nil
	// set cache
	_, err = dbmanager.SetCacheWithPX(str, randworker.GetNumbersString(4), 300000)
	if err != nil {
		meetError("SetCacheWithPX", err)
	}
	return err
}

func (account *Account) FindUid(secret *authparams.AuthSecret, b *int64) error {
	// get info
	uid, _, _, _, _, err := getInfo(secret)
	if err != nil && err.Error() != "uid not found" {
		meetError("getInfo", err)
		return err
	}
	*b = uid
	return nil
}

func (account *Account) Register(secret *authparams.AuthSecret, res *authparams.ResWithToken) error {
	if secret.AccountType == "phone" && secret.CodeType == "code" {
		// phone & code Register type
		// get code cache
		code, err := dbmanager.GetCache("auth&phone=" + secret.Account)
		if err != nil {
			meetError("dbmanager.GetCache", err)
			return err
		}
		// auth code
		c := secret.Code
		if c == "" || code != c {
			// wrong code
			err = errors.New("auth error")
			return err
		} else {
			// correct code
			// delete code cache
			_, err = dbmanager.DelCache("auth&phone=" + secret.Account)
			if err != nil {
				meetError("dbmanager.DelCache", err)
				return err
			}
		}
	} else if secret.AccountType == "username" && secret.CodeType == "password" {
		// username & password Register type
		// empty action
	} else {
		// invalid Register type
		return errors.New("process not accepted")
	}
	// create account
	err := supports.CreateAccount(secret)
	if err != nil {
		meetError("supports.CreateAccount", err)
		return err
	}
	// make token
	uid, _, pt, pl, nk, err := getInfo(secret)
	if err != nil {
		meetError("getInfo", err)
		return err
	}
	token, err := supports.MakeToken(uid, pt, pl, nk)
	if err != nil {
		meetError("makeToken", err)
		return err
	}
	// assign
	res.Uid = uid
	res.Nickname = nk
	res.PrivilegeType = pt
	res.PrivilegeLevel = pl
	res.Token = token
	return err
}

func (account *Account) GetNickname(secret *authparams.AuthSecret, res *authparams.ResWithoutToken) error {
	// get info
	_, _, _, _, nk, err := getInfo(secret)
	if err != nil && err.Error() != "uid not found" {
		meetError("auth", err)
		return err
	}
	if nk == "" {
		return nil
	}
	// hide real-nickname
	rnk := []rune(nk)
	var snk = "*"
	if len(rnk) > 1 {
		snk += string(rnk[1:])
	}
	res.Nickname = snk
	return nil
}

func meetError(action string, err error) {
	log.WithFields(log.Fields{
		"action": action,
		"error":  err,
	}).Warn()
}

func auth(secret *authparams.AuthSecret) (uid int64, pt string, pl int64, nk string, err error) {
	// fetch auth info
	var password string
	uid, password, pt, pl, nk, err = getInfo(secret)
	if err != nil {
		return
	}
	// auth
	switch secret.CodeType {
	case "code":
		// code auth
		// get code from cache
		var code string
		code, err = dbmanager.GetCache("auth&uid=" + strconv.FormatInt(uid, 10))
		c := secret.Code
		if c == "" || code != c {
			// invalid
			err = errors.New("auth error")
		} else {
			// valid
			// delete cache
			_, err = dbmanager.DelCache("auth&uid=" + strconv.FormatInt(uid, 10))
		}
	case "password":
		// password auth
		c := secret.Code
		if c == "" || password != c {
			err = errors.New("auth error")
		}
	case "wxopenid":
		// wxOpenid auth
		// NOTE! Only internally available! it's Danger!
		c := secret.Code
		if c == "" {
			err = errors.New("auth error")
		}
	case "stuid":
		// stuid auth
		// privilege type must be student
		c := secret.Code
		if c == "" || pt != "student" {
			err = errors.New("auth error")
		}
	default:
		err = errors.New("codeType not exist")
	}
	return
}

func getInfo(secret *authparams.AuthSecret) (uid int64, password string, pt string, pl int64, nk string, err error) {
	switch secret.AccountType {
	case "phone":
		uid, password, pt, pl, nk, err = supports.QueryAuth("phone", secret)
	case "username":
		uid, password, pt, pl, nk, err = supports.QueryAuth("username", secret)
	case "stuid":
		uid, password, pt, pl, nk, err = supports.QueryAuth("stuid", secret)
	case "wxopenid":
		uid, password, pt, pl, nk, err = supports.QueryAuth("wxopenid", secret)
	default:
		err = errors.New("accountType not exist")
	}
	return
}
