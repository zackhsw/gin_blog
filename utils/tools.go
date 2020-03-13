package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

const (
	secret = "asdf@#$@#$@#$"
)
func MD5(str string) string{
	md5str := fmt.Sprintf("%x",md5.Sum(append([]byte(str),[]byte(secret)...)))
	return md5str
}

func SwitchTimeStampToStr(timeStamp int64) string{
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}