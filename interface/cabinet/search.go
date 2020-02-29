package cabinet

import (
	"../../managers/dbmanager"
	"../../utils/randworker"
	"../base"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ToSearch(c *gin.Context) {
	word := c.Query("q")
	if word == "" {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	rand := randworker.GetAlnumString(32)
	go record(4, rand, word)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"query": word,
		"rand":  rand,
	})
}

func UpdateSearch(c *gin.Context) {
	process := c.PostForm("process")
	rand := c.PostForm("rand")
	if process == "" || rand == "" {
		base.ServeError(c, "params error", errors.New("params error"))
		return
	}
	flag := false
	switch process {
	case "google":
		go update(4, rand, "google")
		flag = true
	case "baidu":
		go update(4, rand, "baidu")
		flag = true
	default:
	}
	if !flag {
		base.ServeError(c, "invalid lastWord", errors.New("params error"))
	} else {
		c.JSON(http.StatusOK, gin.H{
			"state": "success",
		})
	}
}

func GetSearch(c *gin.Context) {
	res, err := query(4)
	if err != nil {
		base.ServeError(c, "query wrong", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
		"data":  res[:],
	})
}

func record(uid int64, rand string, word string) {
	conn, err := dbmanager.DialPG()
	if err != nil {
		base.LogError("dbmanager.DialPG()", err)
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare(`INSERT INTO search(randword, searchword, uid) VALUES($1, $2, $3);`)
	if err != nil {
		base.LogError("conn.Prepare()", err)
		return
	}
	res, err := stmt.Exec(rand, word, uid)
	if err != nil {
		base.LogError("stmt.Exec()", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		base.LogError("res.RowsAffected()", err)
		return
	}
	if row != 1 {
		base.LogError("resRows wrong", errors.New(strconv.Itoa(int(row))))
		return
	}
}

func update(uid int64, rand, process string) {
	conn, err := dbmanager.DialPG()
	if err != nil {
		base.LogError("dbmanager.DialPG()", err)
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare(`UPDATE search SET lastprocess = $1 WHERE randword = $2 AND uid = $3;`)
	if err != nil {
		base.LogError("conn.Prepare()", err)
		return
	}
	res, err := stmt.Exec(process, rand, uid)
	if err != nil {
		base.LogError("stmt.Exec()", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		base.LogError("res.RowsAffected()", err)
		return
	}
	if row != 1 {
		base.LogError("resRows wrong", errors.New(strconv.Itoa(int(row))))
		return
	}
}

func query(uid int64) (h []map[string]string, err error) {
	conn, err := dbmanager.DialPG()
	if err != nil {
		return
	}
	defer conn.Close()
	stmt, err := conn.Prepare(`SELECT randword, date, searchword, lastprocess FROM search WHERE uid = $1 ORDER BY search.date DESC LIMIT 100;`)
	if err != nil {
		return
	}
	res, err := stmt.Query(uid)
	if err != nil {
		return
	}
	//h = make([]map[string]string, 100)
	i := 0
	for res.Next() {
		var nRandWord, nSearchWord, nLastProcess sql.NullString
		var nDate sql.NullTime
		err = res.Scan(&nRandWord, &nDate, &nSearchWord, &nLastProcess)
		if err != nil {
			return
		}
		t := nDate.Time
		var tmp = make(map[string]string)
		tmp["rand"] = nRandWord.String
		tmp["word"] = nSearchWord.String
		tmp["time"] = strconv.Itoa(t.Year()) + "年" + strconv.Itoa(int(t.Month())) + "月" + strconv.Itoa(t.Day()) + "日 " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second())
		tmp["process"] = nLastProcess.String
		h = append(h, tmp)
		i += 1
	}
	return
}
