package main

import (
	"regexp"
)

// ParseEvent 解析提醒事件
func ParseEvent(content string) string {
	re := regexp.MustCompile(`^.+[提醒|通知|告诉|叫]+[我们大家]{0,2}(.+)$`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

// ParseFrequency 解析提醒频率
func ParseFrequency(content string) {

}

// ParseTrigger 解析提醒触发时间
func ParseTrigger(content string) {

}
