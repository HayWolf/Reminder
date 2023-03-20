package helper

import (
	"time"
)

// GetNextWeekday
// 获取下一个星期x的具体日期 如果当日就是则返回当日
func GetNextWeekday(weekday time.Weekday, base time.Time) time.Time {
	daysUntil := (weekday - base.Weekday() + 7) % 7
	return base.AddDate(0, 0, int(daysUntil))
}

// GetNextDay
// 获取下一个x日的日期，如果当日就是则返回当日
func GetNextDay(day int, base time.Time) time.Time {
	year, month := base.Year(), base.Month()

	// 如果当月该日期已过，或则当前月不存在该日期，则计算下一个月的日期
	if base.Day() > day || !isValid(year, month, day) {
		// 不能直接用base去AddDate，可能命中下个月没有base的日期则会出现跨两个月的问题
		newBase := time.Date(year, month, 1, 0, 0, 0, 0, base.Location())
		year, month = newBase.AddDate(0, 1, 0).Year(), newBase.AddDate(0, 1, 0).Month()
		// 判断新的月份是否有该日期，没有则Add 2个月（实际不会出现一个日期连续两个月不存在，所以最多加2个月）
		if !isValid(year, month, day) {
			year, month = newBase.AddDate(0, 2, 0).Year(), newBase.AddDate(0, 2, 0).Month()
		}
	}

	// 构造下一个指定日的时间对象
	return time.Date(year, month, day, 0, 0, 0, 0, base.Location())
}

func isValid(year int, month time.Month, day int) bool {
	newDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return newDate.Day() == day && newDate.Month() == month && newDate.Year() == year
}
