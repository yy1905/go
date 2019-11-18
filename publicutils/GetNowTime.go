package publicutils

import "time"

// 获取当前时间
func GetNowTime() (nowTime string) {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	return nowStr
}