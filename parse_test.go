package main

import (
	"fmt"
	"github.com/ldigit/config"
	"reflect"
	"testing"
	"time"
)

const ConfigPath = "config.yaml"

func init() {
	// 加载配置文件
	cfg, _ := config.Load(ConfigPath)
	config.SetGlobalConfig(cfg)
}

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

func TestParseFrequency(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  Frequency
	}{
		{
			name:  "这是一个无关键词的话语",
			input: "这是一个无关键词的话语",
			want: Frequency{
				Unit:     TimeUnitNone,
				Interval: 0,
			},
		},
		{
			name:  "半个小时候提醒我去领快递",
			input: "半个小时候提醒我去领快递",
			want: Frequency{
				Unit:     TimeUnitNone,
				Interval: 0,
			},
		},
		{
			name:  "这周六早上8:00提醒我送孩子去画画",
			input: "这周六早上8:00提醒我送孩子去画画",
			want: Frequency{
				Unit:     TimeUnitNone,
				Interval: 0,
			},
		},
		{
			name:  "每个钟头提醒我喝水",
			input: "每个钟头提醒我喝水",
			want: Frequency{
				Unit:     TimeUnitHour,
				Interval: 0,
			},
		},
		{
			name:  "每个钟提醒我喝水",
			input: "每个钟提醒我喝水",
			want: Frequency{
				Unit:     TimeUnitHour,
				Interval: 0,
			},
		},
		{
			name:  "每两个小时提醒我喝水",
			input: "每两个时提醒我喝水",
			want: Frequency{
				Unit:     TimeUnitHour,
				Interval: 1,
			},
		},
		{
			name:  "每天早上8:00叫我起床",
			input: "每天早上8:00叫我起床",
			want: Frequency{
				Unit:     TimeUnitDay,
				Interval: 0,
			},
		},
		{
			name:  "每隔一天中午12:00提醒我午休",
			input: "每隔一天中午12:00提醒我午休",
			want: Frequency{
				Unit:     TimeUnitDay,
				Interval: 1,
			},
		},
		{
			name:  "每隔两天下午7:20提醒我要洗头了",
			input: "每隔两天下午7:20提醒我要洗头了",
			want: Frequency{
				Unit:     TimeUnitDay,
				Interval: 2,
			},
		},
		{
			name:  "每三天下午7:20提醒我要洗头了",
			input: "每三天下午7:20提醒我要洗头了",
			want: Frequency{
				Unit:     TimeUnitDay,
				Interval: 2,
			},
		},
		{
			name:  "每周五下午5:30提醒我早点下班",
			input: "每周五下午5:30提醒我早点下班",
			want: Frequency{
				Unit:     TimeUnitWeek,
				Interval: 0,
			},
		},
		{
			name:  "每个礼拜五下午5:30提醒我早点下班",
			input: "每个礼拜五下午5:30提醒我早点下班",
			want: Frequency{
				Unit:     TimeUnitWeek,
				Interval: 0,
			},
		},
		{
			name:  "每两星期六10:00提醒我预定会议室",
			input: "每两星期六10:00提醒我预定会议室",
			want: Frequency{
				Unit:     TimeUnitWeek,
				Interval: 1,
			},
		},
		{
			name:  "每两周周六10:00提醒我预定会议室",
			input: "每两周周六10:00提醒我预定会议室",
			want: Frequency{
				Unit:     TimeUnitWeek,
				Interval: 1,
			},
		},
		{
			name:  "每个月1号9:58提醒我领优惠券",
			input: "每个月1号9:58提醒我领优惠券",
			want: Frequency{
				Unit:     TimeUnitMonth,
				Interval: 0,
			},
		},
		{
			name:  "每隔一个月15号8:30提醒我去爬山",
			input: "每隔一个月15号8:30提醒我去爬山",
			want: Frequency{
				Unit:     TimeUnitMonth,
				Interval: 1,
			},
		},
		{
			name:  "每两个月1号10:30提醒我要理发了",
			input: "每两个月1号10:30提醒我要理发了",
			want: Frequency{
				Unit:     TimeUnitMonth,
				Interval: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := ParseFrequency(tt.input)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("ParseFrequency rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}

func TestParseTrigger(t *testing.T) {

	type args struct {
		text string
		base time.Time
	}
	tests := []struct {
		name  string
		input args
		want  time.Time
	}{
		{
			name: "这是一个无关键词的话语",
			input: args{
				text: "这是一个无关键词的话语",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "半个小时后提醒我去领快递",
			input: args{
				text: "半个小时后提醒我去领快递",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 30, 0, 0, time.UTC),
		},
		{
			name: "这周六早上8:00提醒我送孩子去画画",
			input: args{
				text: "这周六早上8:00提醒我送孩子去画画",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 7, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "每个钟头提醒我喝水",
			input: args{
				text: "每个钟头提醒我喝水",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "每两个小时提醒我喝水",
			input: args{
				text: "每两个小时提醒我喝水",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "每天早上8:00叫我起床",
			input: args{
				text: "每天早上8:00叫我起床",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "每隔一天中午12:00提醒我午休",
			input: args{
				text: "每隔一天中午12:00提醒我午休",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name: "每隔两天下午7:20提醒我要洗头了",
			input: args{
				text: "每隔两天下午7:20提醒我要洗头了",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 19, 20, 0, 0, time.UTC),
		},
		{
			name: "每三天下午7:20提醒我要洗头了",
			input: args{
				text: "每三天下午7:20提醒我要洗头了",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 19, 20, 0, 0, time.UTC),
		},
		{
			name: "每周五下午5:30提醒我早点下班",
			input: args{
				text: "每周五下午5:30提醒我早点下班",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 6, 17, 30, 0, 0, time.UTC),
		},
		{
			name: "每个礼拜五下午5:30提醒我早点下班",
			input: args{
				text: "每个礼拜五下午5:30提醒我早点下班",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 6, 17, 30, 0, 0, time.UTC),
		},
		{
			name: "每两星期六10:00提醒我预定会议室",
			input: args{
				text: "每两星期六10:00提醒我预定会议室",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 7, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "每两周周六10:00提醒我预定会议室",
			input: args{
				text: "每两周周六10:00提醒我预定会议室",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 7, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "每个月1号9:58提醒我领优惠券",
			input: args{
				text: "每个月1号9:58提醒我领优惠券",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 9, 58, 0, 0, time.UTC),
		},
		{
			name: "每隔一个月15号8:30提醒我去爬山",
			input: args{
				text: "每隔一个月15号8:30提醒我去爬山",
				base: time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 2, 15, 8, 30, 0, 0, time.UTC),
		},
		{
			name: "每两个月1号10:30提醒我要理发了",
			input: args{
				text: "每两个月1号10:30提醒我要理发了",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 10, 30, 0, 0, time.UTC),
		},
		{
			name: "2月29号10:30提醒我订机票",
			input: args{
				text: "2月29号10:30提醒我订机票",
				base: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2024, 2, 29, 10, 30, 0, 0, time.UTC),
		},
		{
			name: "下下周二提醒我要理发了",
			input: args{
				text: "下下周二提醒我要理发了",
				base: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 17, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "明天早上10:00提醒我要跟同事开会",
			input: args{
				text: "明天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "后天早上10:00提醒我要跟同事开会",
			input: args{
				text: "后天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 3, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "大后天早上10:00提醒我要跟同事开会",
			input: args{
				text: "大后天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 4, 10, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := ParseTrigger(tt.input.text, tt.input.base)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("ParseTrigger rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}

func TestParseTriggerByGPT(t *testing.T) {
	type args struct {
		text string
		base time.Time
	}
	tests := []struct {
		name  string
		input args
		want  time.Time
	}{
		{
			name: "这是一个无关键词的话语",
			input: args{
				text: "这是一个无关键词的话语",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "半个小时后提醒我去领快递",
			input: args{
				text: "半个小时后提醒我去领快递",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 30, 0, 0, time.UTC),
		},
		{
			name: "这周六早上8:00提醒我送孩子去画画",
			input: args{
				text: "这周六早上8:00提醒我送孩子去画画",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 7, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "每个钟头提醒我喝水",
			input: args{
				text: "每个钟头提醒我喝水",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "每两个小时提醒我喝水",
			input: args{
				text: "每两个小时提醒我喝水",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "每天早上8:00叫我起床",
			input: args{
				text: "每天早上8:00叫我起床",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC),
		},
		{
			name: "每隔一天中午12:00提醒我午休",
			input: args{
				text: "每隔一天中午12:00提醒我午休",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		//{
		//	name: "每隔两天下午7:20提醒我要洗头了",
		//	input: args{
		//		text: "每隔两天下午7:20提醒我要洗头了",
		//		base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		//	},
		//	want: time.Date(2023, 1, 1, 19, 20, 0, 0, time.UTC),
		//},
		//{
		//	name: "每三天下午7:20提醒我要洗头了",
		//	input: args{
		//		text: "每三天下午7:20提醒我要洗头了",
		//		base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		//	},
		//	want: time.Date(2023, 1, 1, 19, 20, 0, 0, time.UTC),
		//},
		{
			name: "每周五下午5:30提醒我早点下班",
			input: args{
				text: "每周五下午5:30提醒我早点下班",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 6, 17, 30, 0, 0, time.UTC),
		},
		{
			name: "每个礼拜五下午5:30提醒我早点下班",
			input: args{
				text: "每个礼拜五下午5:30提醒我早点下班",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 6, 17, 30, 0, 0, time.UTC),
		},
		//{
		//	name: "每两星期六10:00提醒我预定会议室",
		//	input: args{
		//		text: "每两星期六10:00提醒我预定会议室",
		//		base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		//	},
		//	want: time.Date(2023, 1, 7, 10, 0, 0, 0, time.UTC),
		//},
		//{
		//	name: "每两周周六10:00提醒我预定会议室",
		//	input: args{
		//		text: "每两周周六10:00提醒我预定会议室",
		//		base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		//	},
		//	want: time.Date(2023, 1, 6, 10, 0, 0, 0, time.UTC),
		//},
		{
			name: "每个月1号9:58提醒我领优惠券",
			input: args{
				text: "每个月1号9:58提醒我领优惠券",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 9, 58, 0, 0, time.UTC),
		},
		{
			name: "每隔一个月15号8:30提醒我去爬山",
			input: args{
				text: "每隔一个月15号8:30提醒我去爬山",
				base: time.Date(2023, 1, 16, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 2, 15, 8, 30, 0, 0, time.UTC),
		},
		{
			name: "每两个月1号10:30提醒我要理发了",
			input: args{
				text: "每两个月1号10:30提醒我要理发了",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 3, 1, 10, 30, 0, 0, time.UTC),
		},
		//{
		//	name: "2月29号10:30提醒我订机票",
		//	input: args{
		//		text: "2月29号10:30提醒我订机票",
		//		base: time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		//	},
		//	want: time.Date(2024, 2, 29, 10, 30, 0, 0, time.UTC),
		//},
		{
			name: "下下周二提醒我要理发了",
			input: args{
				text: "下下周二提醒我要理发了",
				base: time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 17, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "明天早上10:00提醒我要跟同事开会",
			input: args{
				text: "明天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 2, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "后天早上10:00提醒我要跟同事开会",
			input: args{
				text: "后天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 3, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "大后天早上10:00提醒我要跟同事开会",
			input: args{
				text: "大后天早上10:00提醒我要跟同事开会",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 4, 10, 0, 0, 0, time.UTC),
		},
		{
			name: "明天下午六点20分提醒我去取快递",
			input: args{
				text: "明天下午六点20分提醒我去取快递",
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 2, 18, 20, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := ParseTriggerByGPT(tt.input.text, tt.input.base)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("ParseTriggerByGPT rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	//result := "转换成Y-m-d H:i:s格式后为：2023-01-06 17:30:00。"
	//re := regexp.MustCompile(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`)
	//match := re.FindStringSubmatch(result)
	//for i, s := range match {
	//	fmt.Println(i, s)
	//}
	base := time.Date(2013, 1, 1, 13, 0, 0, 0, time.UTC)
	fmt.Println(base.Format("2006年01月02日 15:04:05"))
}
