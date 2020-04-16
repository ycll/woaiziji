package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

//传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string{
	md5str:=fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return md5str
}

func SwitchTimeStampToData(timestamp int64) string {
	timeNow := time.Unix(timestamp, 0)
	timeString := timeNow.Format("2006-01-02 15:04:05")
	return timeString
}