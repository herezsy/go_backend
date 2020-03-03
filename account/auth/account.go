package auth

import (
	"../../params/authparams"
	"./supports"
	"errors"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

type Account int

func init() {
	log.SetReportCaller(true)
}

func (account *Account) AuthAndGetToken(secret *authparams.Params, res *authparams.Params) error {
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

func (account *Account) Auth(secret *authparams.Params, res *authparams.Params) error {
	// auth info
	uid, pt, pl, nk, err := auth(secret)
	if err != nil {
		meetError("auth", err)
		return err
	}
	// make secret
	token, err := supports.MakeSecret(uid)
	if err != nil {
		meetError("supports.MakeSecret", err)
		return err
	}
	// assign
	res.Uid = uid
	res.Nickname = nk
	res.PrivilegeType = pt
	res.PrivilegeLevel = pl
	res.Token = token
	return nil
}

func (account *Account) AuthConfirm(secret *authparams.Params, res *authparams.Params) (err error) {
	// auth secret
	err = supports.CheckSecret(secret.Uid, secret.Token)
	if err != nil {
		meetError("supports.CheckSecret", err)
	}
	return
}

func (account *Account) AuthToken(secret *authparams.Params, res *authparams.Params) error {
	// get message from token
	uid, pt, pl, nk, t, err := supports.DecodeToken(secret.Token)
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

func (account *Account) ChangeAuth(secret *authparams.Params, res *authparams.Params) (err error) {
	// auth secret
	err = supports.CheckSecret(secret.Uid, secret.Token)
	if err != nil {
		meetError("supports.CheckSecret", err)
		return
	}
	// apply change
	switch secret.AccountType {
	case "username":
		err = supports.ChangeAccount(secret.Uid, "username", secret.Account)
	case "phone":
		err = supports.CheckCodeWithPhone(secret.Account, secret.Token)
		if err != nil {
			return
		}
		err = supports.ChangeAccount(secret.Uid, "phone", secret.Account)
	case "stuid":
		err = supports.ChangeAccount(secret.Uid, "stuid", secret.Account)
	case "":
	default:
		err = errors.New("wrong type")
	}
	switch secret.CodeType {
	case "password":
		err = supports.ChangeAccount(secret.Uid, "password", secret.Code)
	case "wxopenid":
		err = supports.ChangeAccount(secret.Uid, "wxopenid", secret.Code)
	case "":
	default:
		err = errors.New("wrong type")
	}
	return
}

func (account *Account) Echo(str *string, res *string) error {
	*res = *str
	return nil
}

func (account *Account) SendCode(secret *authparams.Params, res *authparams.Params) error {
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
	if err != nil && err.Error() == "uid not found" {
		// no-register type
		err = nil
		err = supports.SendCodeWithPhone(secret.Account)
	} else {
		// register type
		err = supports.SendCodeWithUid(uid)
	}
	return err
}

func (account *Account) FindUid(secret *authparams.Params, b *int64) error {
	// get info
	uid, _, _, _, _, err := getInfo(secret)
	if err != nil && err.Error() != "uid not found" {
		meetError("getInfo", err)
		return err
	}
	*b = uid
	return nil
}

func (account *Account) Register(secret *authparams.Params, res *authparams.Params) error {
	if secret.AccountType == "phone" && secret.CodeType == "code" {
		// phone & code Register type
		// auth code
		err := supports.CheckCodeWithPhone(secret.Account, secret.Token)
		if err != nil {
			return err
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

func (account *Account) GetNickname(secret *authparams.Params, res *authparams.Params) error {
	// get info
	uid, _, _, _, nk, err := getInfo(secret)
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
	res.Uid = uid
	return nil
}

func (account *Account) GetAuthProcess(secret *authparams.Params, res *authparams.Params) error {
	uid := secret.Uid
	if uid == 0 {
		return errors.New("uid is empty")
	}
	err := supports.QueryProcess(uid, &(res.Process))
	if err != nil {
		return err
	}
	return nil
}

func meetError(action string, err error) {
	log.WithFields(log.Fields{
		"action": action,
		"error":  err,
	}).Warn()
}

func auth(secret *authparams.Params) (uid int64, pt string, pl int64, nk string, err error) {
	// fetch auth info
	var password string
	uid, password, pt, pl, nk, err = getInfo(secret)
	if err != nil {
		return
	}
	// auth
	c := secret.Account
	if c == "" {
		err = errors.New("auth error")
	}
	switch secret.AccountType {
	case "wxopenid":
		// wxOpenid auth
		// NOTE! Only internally available! it's Danger!
	case "stuid":
		// stuid auth
		// privilege type must be student
		if pt != "student" {
			err = errors.New("auth error")
		}
		switch secret.CodeType {
		case "password":
			if password != "" {
				err = checkPassword(secret.Code, password)
			}
		default:
			err = errors.New("auth error")
		}
	case "phone":
		switch secret.CodeType {
		case "password":
			err = checkPassword(secret.Code, password)
		case "code":
			err = supports.CheckCodeWithPhone(secret.Account, secret.Code)
		default:
			err = errors.New("auth error")
		}
	case "username":
		switch secret.CodeType {
		case "password":
			err = checkPassword(secret.Code, password)
		case "code":
			err = supports.CheckCodeWithPhone(secret.Account, secret.Code)
		default:
			err = errors.New("auth error")
		}
	case "uid":
		switch secret.CodeType {
		case "password":
			err = checkPassword(secret.Code, password)
		case "code":
			err = supports.CheckCodeWithPhone(secret.Account, secret.Code)
		default:
			err = errors.New("auth error")
		}
	default:
		err = errors.New("auth error")
	}
	return
}

func getInfo(secret *authparams.Params) (uid int64, password string, pt string, pl int64, nk string, err error) {
	switch secret.AccountType {
	case "uid":
		uid, password, pt, pl, nk, err = supports.QueryAuth("uid", secret.Account)
	case "phone":
		uid, password, pt, pl, nk, err = supports.QueryAuth("phone", secret.Account)
	case "username":
		uid, password, pt, pl, nk, err = supports.QueryAuth("username", secret.Account)
	case "stuid":
		uid, password, pt, pl, nk, err = supports.QueryAuth("stuid", secret.Account)
	case "wxopenid":
		uid, password, pt, pl, nk, err = supports.QueryAuth("wxopenid", secret.Account)
	default:
		err = errors.New("accountType not exist")
	}
	return
}

func checkPassword(origin, crypto string) error {
	if origin == "" || origin != crypto {
		return errors.New("auth error")
	}
	return nil
}
