package randworker

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

func GetBytes(l int) []byte {
	a := new(big.Int).SetUint64(uint64(255))
	key := make([]byte, l)
	for i, _ := range key {
		r, _ := rand.Int(rand.Reader, a)
		rr := r.Bytes()
		if rr != nil {
			key[i] = rr[0]
		} else {
			key[i] = 0
		}
	}
	return key
}

func GetNumbersString(l int) (res string) {
	a := new(big.Int).SetUint64(uint64(10))
	key := make([]byte, l)
	for range key {
		r, _ := rand.Int(rand.Reader, a)
		res += strconv.FormatInt(r.Int64(), 10)
	}
	return
}

func GetAlnumString(l int) (res string) {
	a := new(big.Int).SetUint64(uint64(62))
	key := make([]byte, l)
	for range key {
		r, _ := rand.Int(rand.Reader, a)
		rr := r.Int64()
		if rr < 10 {
			res += strconv.FormatInt(rr, 10)
		} else if rr < 36 {
			res += string(rr + 55)
		} else {
			res += string(rr + 61)
		}
	}
	return
}
