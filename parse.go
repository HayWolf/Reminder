package main

import (
	"regexp"
)

// ParseEvent 解析提醒事件
func ParseEvent(content string) string {
	// 移除空格、换行符
	re := regexp.MustCompile(`[\s\n]+`)
	content = re.ReplaceAllString(content, "")
	// 通过「提醒」等关键词匹配目标事件
	re = regexp.MustCompile(`^.+(提醒|通知|告诉|叫)+[我们大家]{0,2}(.+)$`)
	match := re.FindStringSubmatch(content)
	if len(match) > 2 {
		return match[2]
	}
	return ""
}

// ParseFrequency 解析提醒频率
func ParseFrequency(content string) {

}

// ParseTrigger 解析提醒触发时间
func ParseTrigger(content string) {

}
