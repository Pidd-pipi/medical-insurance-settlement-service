package util

import "time"

func timeNowDate() string {
	return time.Now().Format("20060102")
}

// NowTime 返回当前时间。
func NowTime() time.Time { return time.Now() }

// TodayRange 返回某时间所在自然日的起止时间。
func TodayRange(now time.Time) (time.Time, time.Time) {
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return start, start.AddDate(0, 0, 1)
}

// TodayDate 返回 Asia/Shanghai 时区的今天日期（与 DB 会话 TimeZone=Asia/Shanghai 一致）。
func TodayDate() string {
	return time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02")
}
