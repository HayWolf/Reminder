package main

import (
	"reflect"
	"testing"
)

func TestParseEvent(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "这是一个无关键词的话语",
			input: "这是一个无关键词的话语",
			want:  "",
		},
		{
			name:  "下午6点叫去吃饭",
			input: "下午6点叫去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点通知去吃饭",
			input: "下午6点通知去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点提醒去吃饭",
			input: "下午6点提醒去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点叫我去吃饭",
			input: "下午6点叫我去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点叫我们去吃饭",
			input: "下午6点叫我们去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点叫大家去吃饭",
			input: "下午6点叫大家去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点通知我去吃饭",
			input: "下午6点通知我去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点通知我们去吃饭",
			input: "下午6点通知我们去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点通知大家去吃饭",
			input: "下午6点通知大家去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点通知我去吃饭",
			input: "下午6点通知我去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点提醒我们去吃饭",
			input: "下午6点提醒我们去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "下午6点提醒大家去吃饭",
			input: "下午6点提醒大家去吃饭",
			want:  "去吃饭",
		},
		{
			name:  "明天8点叫我提醒大家开晨会",
			input: "明天8点叫我提醒大家开晨会",
			want:  "开晨会",
		},
		{
			name:  "明天8点提醒我通知开晨会",
			input: "明天8点提醒我通知开晨会",
			want:  "开晨会",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := ParseEvent(tt.input)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("ParseEvent rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}
