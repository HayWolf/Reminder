package helper

import (
	"fmt"
	"time"
)

// GetNextWeekday
// 获取下一个星期x的具体日期 如果当日就是则返回当日
func GetNextWeekday(weekday time.Weekday, base time.Time) time.Time {
	daysUntil := (weekday - base.Weekday() + 7) % 7
	return base.AddDate(0, 0, int(daysUntil))
}

// GetNextDate
// 获取下一个该日期，如果当日就是则返回当日
func GetNextDate(day int, base time.Time) time.Time {
	year, month := base.Year(), base.Month()
	fmt.Println(year, month, base)

	// 如果当月已经包含该日期，则计算下一个月的日期
	if base.Day() > day {
		// 不能直接用base去AddDate，可能命中下个月没有base的日期则会出现跨两个月的问题
		newBase := time.Date(base.Year(), base.Month(), 1, 0, 0, 0, 0, base.Location())
		year = newBase.AddDate(0, 1, 0).Year()
		month = newBase.AddDate(0, 1, 0).Month()
	}

	// 构造下一个指定日的时间对象
	nextDate := time.Date(year, month, day, 0, 0, 0, 0, base.Location())
	fmt.Println(year, month, base, nextDate)

	// 如果下一个月不存在该日期，则继续计算下下个月的日期
	if nextDate.Day() != day || nextDate.Month() != month || nextDate.Year() != year {
		nextDate = nextDate.AddDate(0, 1, 0)
	}

	return nextDate
}
