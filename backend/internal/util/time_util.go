package util

import "time"

func timeNowDate() string {
	return time.Now().Format("20060102")
}

// NowTime 返回当前时间。
func NowTime() time.Time { return time.Now() }

// TodayDate 返回 Asia/Shanghai 时区的今天日期（与 DB 会话 TimeZone=Asia/Shanghai 一致）。
func TodayDate() string {
	return time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02")
}

// StartOfToday 返回 Asia/Shanghai 时区今日 00:00:00 对应的时间点，用于“当日”重复性检查的左闭边界。
// 显式使用 CST 时区，避免应用进程时区与 DB 会话时区（Asia/Shanghai）不一致导致跨天判定错位。
func StartOfToday() time.Time {
	cst := time.FixedZone("CST", 8*3600)
	now := time.Now().In(cst)
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, cst)
}
