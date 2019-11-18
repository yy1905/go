package publicutils

import (
	"math/rand"
	"sync"
	"time"
)

// 返回随机n个字符串
func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var r *rand.Rand
	var once = sync.Once{}
	once.Do(func() {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
	var result string
	for i := 0; i < n; i++ {
		result += string(str[r.Intn(len(str))])
	}
	return result
}