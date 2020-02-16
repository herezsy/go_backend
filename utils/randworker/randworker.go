package randworker

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GetNumbersString(l int) (res string) {
	a := new(big.Int).SetUint64(uint64(10))
	key := make([]byte, l)
	for range key {
		r, _ := rand.Int(rand.Reader, a)
		res += strconv.FormatInt(r.Int64(), 10)
	}
	return
}
