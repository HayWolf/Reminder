package main

import (
	"github.com/HayWolf/Reminder/helper"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TimeUnit 时间单位
type TimeUnit string

// Frequency 频率结构体
type Frequency struct {
	Unit     TimeUnit // 重复单位
	Interval int      // 重复间隔
}

const (
	TimeUnitNone    TimeUnit = "none"    // 无
	TimeUnitMinute  TimeUnit = "minute"  // 分钟
	TimeUnitHour    TimeUnit = "hour"    // 小时
	TimeUnitDay     TimeUnit = "day"     // 天
	TimeUnitWeek    TimeUnit = "week"    // 周
	TimeUnitMonth   TimeUnit = "month"   // 月
	TimeUnitUnKnown TimeUnit = "unknown" // 未支持的单位
)

// ChToUnitMap 中文时间单位转TimeUnit定义
var ChToUnitMap = map[string]TimeUnit{
	"分":  TimeUnitMinute,
	"分钟": TimeUnitMinute,
	"小时": TimeUnitHour,
	"时":  TimeUnitHour,
	"钟头": TimeUnitHour,
	"钟":  TimeUnitHour,
	"天":  TimeUnitDay,
	"日":  TimeUnitDay,
	"周":  TimeUnitWeek,
	"星期": TimeUnitWeek,
	"礼拜": TimeUnitWeek,
	"月":  TimeUnitMonth,
}

// ParseEvent 解析提醒事件
func ParseEvent(content string) string {
	// 移除空格、换行符
	re := regexp.MustCompile(`[\s\n]+`)
	content = re.ReplaceAllString(content, "")
	// 通过「提醒」等关键词匹配目标事件
	re = regexp.MustCompile(`^.*(提醒|通知|告诉|叫)+(大家|我们|我)?(.+)$`)
	match := re.FindStringSubmatch(content)
	if len(match) > 3 {
		return match[3]
	}
	return ""
}

// ParseFrequency 解析提醒频率
func ParseFrequency(content string) Frequency {

	// 默认无重复无间隔
	ret := Frequency{
		Unit:     TimeUnitNone,
		Interval: 0,
	}

	// 暂时不支持分钟级的频率，避免发消息太频繁
	re := regexp.MustCompile(`^每(逢|当|到)?(隔)?([0123456789零一二两三四五六七八九十百千]*)?个?(小时|钟头|星期|礼拜|钟|时|天|日|周|月).*$`)
	match := re.FindStringSubmatch(content)
	if len(match) > 4 {
		// 匹配到了「每」，默认为不支持的频率
		ret.Unit = TimeUnitUnKnown

		// 确定频率单位
		if unit, ok := ChToUnitMap[match[4]]; ok {
			ret.Unit = unit
		}

		// 确定间隔
		if len(match[3]) > 0 {
			num := helper.StringToNumber(match[3])
			// 没有「隔」则数值要减1，「每两天」是间隔1天，「每隔两天」是间隔2天
			if len(match[2]) == 0 && num > 0 {
				num--
			}
			ret.Interval = int(num)
		}

	}

	return ret
}

// ParseTrigger 解析提醒触发时间
func ParseTrigger(content string, base time.Time) time.Time {

	// 匹配明天后天大后天
	re := regexp.MustCompile("^(明|后|大后)天")
	match := re.FindStringSubmatch(content)
	if len(match) > 1 && len(match[1]) > 0 {
		log.Printf("进入明后天匹配：%s", match[0])
		switch match[1] {
		case "明":
			base = base.Add(24 * time.Hour)
		case "后":
			base = base.Add(2 * 24 * time.Hour)
		case "大后":
			base = base.Add(3 * 24 * time.Hour)
		}
		log.Printf("base: %s", base)
	}

	// 匹配XX时间后，30分钟后，2小时后，20天后，过两个礼拜，1个月后
	re = regexp.MustCompile(`^过?([0123456789零一二两三四五六七八九十百千半]*)?个?(分钟|小时|钟头|星期|礼拜|分|钟|时|天|日|周|月)后?.*$`)
	match = re.FindStringSubmatch(content)
	// 量词和单位必须都满足
	if len(match) > 2 && len(match[1]) > 0 && len(match[2]) > 0 {
		log.Printf("进入xx时间后匹配: %s", match[0])
		num := helper.StringToNumber(match[1])
		unit, ok := ChToUnitMap[match[2]]
		// 支持半小时、半天、半个月
		if match[1] == "半" {
			if unit == TimeUnitHour {
				num, unit = 30, TimeUnitMinute
			} else if unit == TimeUnitDay {
				num, unit = 6, TimeUnitHour
			} else if unit == TimeUnitMonth {
				num, unit = 15, TimeUnitDay
			}
		}
		if ok && num > 0 {
			switch unit {
			case TimeUnitMinute:
				base = base.Add(time.Duration(num) * time.Minute)
			case TimeUnitHour:
				base = base.Add(time.Duration(num) * time.Hour)
			case TimeUnitDay:
				base = base.AddDate(0, 0, int(num))
			case TimeUnitWeek:
				base = base.AddDate(0, 0, int(num)*7)
			case TimeUnitMonth:
				base = base.AddDate(0, int(num), 0)
			}
		}
		log.Printf("base: %s", base)
	}

	// 匹配周x星期x
	re = regexp.MustCompile(`(这|下*)?(周|星期|礼拜)([1234567一二三四五六日天])`)
	match = re.FindStringSubmatch(content)
	if len(match) > 3 && len(match[2]) > 0 && len(match[3]) > 0 {
		log.Printf("进入星期x匹配：%s", match[0])
		num := helper.StringToNumber(match[3])
		if match[3] == "日" || match[3] == "天" {
			num = 7
		}
		base = helper.GetNextWeekday(time.Weekday(num), base)
		// 下周，下下周
		if next := strings.Count(match[1], "下"); next > 0 {
			base = base.AddDate(0, 0, next*7)
		}
		log.Printf("base: %s", base)
	}

	// 匹配xx月xx日 xx号
	re = regexp.MustCompile(`(([0123456789一二三四五六七八九十]+)月)?([0123456789一二三四五六七八九十]+)(日|号)`)
	match = re.FindStringSubmatch(content)
	if len(match) > 4 {
		log.Printf("进入x月x日匹配: %s", match[0])
		mouth := helper.StringToNumber(match[2])
		day := helper.StringToNumber(match[3])

		newBase := base
		// 目标月份跨年
		if mouth > 0 && time.Month(mouth) < base.Month() {
			newBase = time.Date(base.Year()+1, time.Month(mouth), 1, base.Hour(), base.Minute(), 0, 0, base.Location())
		} else if mouth > 0 {
			// 没有跨年
			newBase = time.Date(base.Year(), time.Month(mouth), 1, base.Hour(), base.Minute(), 0, 0, base.Location())
		}
		//log.Printf("mouth:%d day:%d newBase:%s", mouth, day, newBase)
		if day > 0 {
			base = helper.GetNextDay(int(day), newBase)
		}
		log.Printf("base: %s", base)
	}

	// 匹配时间格式 23:59分
	content = strings.Replace(content, "：", ":", -1)
	re = regexp.MustCompile(`(\d+):(\d+)`)
	match = re.FindStringSubmatch(content)
	if len(match) > 2 {
		log.Printf("进入xx:xx时间匹配: %s", match[0])
		hour, _ := strconv.Atoi(match[1])
		minute, _ := strconv.Atoi(match[2])
		// 伴随有「下午」「晚上」等词
		// TODO:支持晚上12点
		re = regexp.MustCompile(`下午|晚上|傍晚`)
		if hour < 12 && re.MatchString(content) {
			hour += 12
		}
		base = time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location())
		log.Printf("base: %s", base)
	}

	return base
}
