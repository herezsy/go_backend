package aescryption

import (
	"../../utils/randworker"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var key []byte

func init() {
	log.SetReportCaller(true)
	key = newKey()
}

func newKey() []byte {
	key = randworker.GetBytes(16)
	fmt.Println(key)
	return key
}

func Encrypt(str string) (res string, err error) {
	if str == "" {
		err = errors.New("empty encrypt string")
		return
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	b := []byte(str)
	l := len(b)
	var k int = l
	if l%aes.BlockSize != 0 {
		k = ((l / aes.BlockSize) + 1) * aes.BlockSize
		b = b[:k]
	}
	br := make([]byte, k)
	// NOTE! aes can ONLY Encrypt 16 bytes of data at once
	// the data MORE THAN limit will COME UP WITH 0!!!!!
	for i := 0; i < k; i += 16 {
		c.Encrypt(br[i:i+16], b[i:i+16])
	}
	// CAN NOT Truncate the []byte to its original length
	res = base64.StdEncoding.EncodeToString(br)
	return
}

func Decrypt(str string) (res string, err error) {
	if str == "" {
		err = errors.New("empty decrypt string")
		return
	}
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	k := len(b)
	if k%aes.BlockSize != 0 {
		err = errors.New("incorrect length")
		return
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	br := make([]byte, k)
	for i := 0; i < k; i += 16 {
		c.Decrypt(br[i:i+16], b[i:i+16])
	}
	for i := range br {
		if br[i] == 0 {
			br = br[:i]
			break
		}
	}
	res = string(br)
	return
}
