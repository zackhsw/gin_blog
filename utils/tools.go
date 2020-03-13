package utils

import (
	"crypto/md5"
	"fmt"
)

const (
	secret = "asdf@#$@#$@#$"
)
func MD5(str string) string{
	md5str := fmt.Sprintf("%x",md5.Sum(append([]byte(str),[]byte(secret)...)))
	return md5str
}
