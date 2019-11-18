package publicutils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// 获取字符串的md5值
func GetStrMD5(rawMsg string) string {
	data := []byte(rawMsg)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str1)
}