package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

// 检查密码
func CheckPassword(inputPassword, userPassword, passwordStr string) (bool, error) {
	if inputPassword == "" {
		return false, errors.New("password is empty")
	}
	clientPassword := addSalt(inputPassword, passwordStr)
	if clientPassword != userPassword {
		return false, nil
	}

	return true, nil
}

// addSalt
func addSalt(inputPassword, passwordStr string) string {
	encryStr := Md5(inputPassword)[16:26] + passwordStr
	hashStr := Md5(encryStr)
	return HashSha256(hashStr)
}

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

func HashSha256(str string) string {
	//使用sha256哈希函数
	h := sha256.New()
	_, _ = h.Write([]byte(str))
	sum := h.Sum(nil)

	//由于是十六进制表示，因此需要转换
	s := hex.EncodeToString(sum)
	return s
}
