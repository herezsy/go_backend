package auth

import (
	"../../managers/dbmanager"
	"../../params/authparams"
	"../../utils/randworker"
	"../aescryption"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
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
	token, err := makeToken(uid, pt, pl, nk)
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
	uid, pt, pl, nk, t, err := decodeToken(secret.Code)
	if err != nil {
		meetError("decodeToken", err)
		return err
	}
	// check whether token is old or not
	tn := time.Now().Add(20 * time.Minute)
	var token string
	if tn.After(t) {
		token, err = makeToken(uid, pt, pl, nk)
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
	uid, _, _, _, _, err := getInfo(secret)
	if err != nil && (err.Error() != "uid not found" || secret.AccountType != "phone") {
		meetError("auth", err)
		return err
	}
	var str string
	if err != nil && err.Error() == "uid not found" {
		str = "auth&phone=" + secret.Account
	} else {
		str = "auth&uid=" + strconv.FormatInt(uid, 10)
	}
	err = nil
	_, err = dbmanager.SetCacheWithPX(str, randworker.GetNumbersString(4), 300000)
	log.Warnf("%v\n", strconv.FormatInt(uid, 10))
	if err != nil {
		meetError("SetCacheWithPX", err)
	}
	return err
}

func (account *Account) FindUid(secret *authparams.AuthSecret, b *int64) error {
	uid, _, _, _, _, err := getInfo(secret)
	if err != nil && err.Error() != "uid not found" {
		return err
	}
	*b = uid
	return nil
}

func (account *Account) Register(secret *authparams.AuthSecret, res *authparams.ResWithToken) error {
	if secret.AccountType == "phone" && secret.CodeType == "code" {
		code, err := dbmanager.GetCache("auth&phone=" + secret.Account)
		if err != nil {
			return err
		}
		c := secret.Code
		if c == "" || code != c {
			err = errors.New("auth error")
		} else {
			_, err = dbmanager.DelCache("auth&phone=" + secret.Account)
		}
		if err != nil {
			return err
		}
	} else if secret.AccountType == "username" && secret.CodeType == "password" {
		// empty action
	} else {
		return errors.New("process not accepted")
	}
	// create account
	err := createAccount(secret)
	if err != nil {
		return err
	}
	// make token
	uid, _, pt, pl, nk, err := getInfo(secret)
	if err != nil {
		meetError("getInfo", err)
		return err
	}
	token, err := makeToken(uid, pt, pl, nk)
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
		var code string
		code, err = dbmanager.GetCache("auth&uid=" + strconv.FormatInt(uid, 10))
		c := secret.Code
		if c == "" || code != c {
			err = errors.New("auth error")
		} else {
			_, err = dbmanager.DelCache("auth&uid=" + strconv.FormatInt(uid, 10))
		}
	case "password":
		c := secret.Code
		if c == "" || password != c {
			err = errors.New("auth error")
		}
	case "wxopenid":
		c := secret.Account
		if c == "" {
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
		uid, password, pt, pl, nk, err = queryAuth("phone", secret)
	case "username":
		uid, password, pt, pl, nk, err = queryAuth("username", secret)
	case "stuid":
		uid, password, pt, pl, nk, err = queryAuth("stuid", secret)
	case "wxopenid":
		uid, password, pt, pl, nk, err = queryAuth("wxopenid", secret)
	default:
		err = errors.New("accountType not exist")
	}
	return
}

func queryAuth(id string, secret *authparams.AuthSecret) (uid int64, password string, pt string, pl int64, nk string, err error) {
	db, err := dbmanager.DialPG()
	if err != nil {
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT uid, password, privilegetype, privilegelevel, nickname FROM account WHERE " + id + "=$1")
	if err != nil {
		return
	}
	res, err := stmt.Query(secret.Account)
	if err != nil {
		return
	}
	if res.Next() {
		var nuid sql.NullInt64
		var npassword, npt, nnk sql.NullString
		var npl sql.NullInt64
		err = res.Scan(&nuid, &npassword, &npt, &npl, &nnk)
		uid = nuid.Int64
		password = npassword.String
		pt = npt.String
		pl = npl.Int64
		nk = nnk.String
		log.WithFields(log.Fields{
			"action":   "queryAuth",
			"uid":      uid,
			"password": password,
			"pt":       pt,
			"pl":       pl,
			"nk":       nk,
		}).Info()
		return
	} else {
		err = errors.New("uid not found")
		return
	}
}

func createAccount(secret *authparams.AuthSecret) (err error) {
	db, err := dbmanager.DialPG()
	if err != nil {
		return
	}
	defer db.Close()
	var res sql.Result
	var stmt *sql.Stmt
	if secret.AccountType == "phone" && secret.CodeType == "code" {
		stmt, err = db.Prepare("INSERT INTO account(registerprocess, registertype, nickname, phone) VALUES ('auth', 'phone', $1, $2);")
		if err != nil {
			return
		}
		res, err = stmt.Exec(secret.Account, secret.Account)
		if err != nil {
			return
		}
	} else if secret.AccountType == "username" && secret.CodeType == "password" {
		stmt, err = db.Prepare("INSERT INTO account(registerprocess, registertype, nickname, username, password) VALUES ('auth', 'password', $1, $2, $3)")
		if err != nil {
			return
		}
		res, err = stmt.Exec(secret.Account, secret.Account, secret.Code)
		if err != nil {
			return
		}
	} else if secret.AccountType == "wxopenid" {
		stmt, err = db.Prepare("INSERT INTO account(registerprocess, registertype, nickname, wxopenid) VALUES ('database', 'wxopenid', $1, $1);")
		if err != nil {
			return
		}
		res, err = stmt.Exec(secret.Account, secret.Account)
		if err != nil {
			return
		}
	} else {
		err = errors.New("wrong type")
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		return
	}
	if row != 1 {
		err = errors.New("error rows effect")
	}
	return
}

func decodeToken(token string) (uid int64, pt string, pl int64, nk string, t time.Time, err error) {
	var message string
	message, err = aescryption.Decrypt(token)

	set := strings.Split(message, "&")
	if len(set) < 5 {
		err = errors.New("token message wrong")
		return
	}
	uid, err = strconv.ParseInt(set[0], 10, 64)
	pt = set[1]
	pl, err = strconv.ParseInt(set[2], 10, 64)
	nk = set[3]
	tt := set[4]
	t, err = time.Parse(time.RFC3339, tt)
	log.WithFields(log.Fields{
		"action": "decodeToken",
		"uid":    uid,
		"pt":     pt,
		"pl":     pl,
		"nk":     nk,
		"time":   t,
	}).Info()
	return
}

func makeToken(uid int64, pt string, pl int64, nk string) (token string, err error) {
	limit := time.Now().Add(60 * time.Minute)
	message := strconv.FormatInt(uid, 10) + "&" + pt + "&" + strconv.FormatInt(pl, 10) + "&" + nk + "&" + limit.Format(time.RFC3339)
	token, err = aescryption.Encrypt(message)
	return
}

func meetError(action string, err error) {
	log.WithFields(log.Fields{
		"Action": action,
		"Error":  err,
	}).Warn()
}
