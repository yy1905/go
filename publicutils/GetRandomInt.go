package publicutils

import (
	"sync"
	"time"
	"math/rand"
)

// 获取随机数[start,end]
var randomMutex sync.Mutex

func GetRandomInt(start, end int) int {
	randomMutex.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := start + r.Intn(end-start+1)
	randomMutex.Unlock()
	return n
}