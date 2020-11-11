package util

import "time"

// 时间格式化
func FomatNowDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}
