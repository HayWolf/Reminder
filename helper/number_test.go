package helper

import (
	"reflect"
	"testing"
)

func TestStringToNumber(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int64
	}{
		{name: "零", input: "零", want: 0},
		{name: "一", input: "一", want: 1},
		{name: "二", input: "二", want: 2},
		{name: "三", input: "三", want: 3},
		{name: "四", input: "四", want: 4},
		{name: "五", input: "五", want: 5},
		{name: "六", input: "六", want: 6},
		{name: "七", input: "七", want: 7},
		{name: "八", input: "八", want: 8},
		{name: "九", input: "九", want: 9},
		{name: "十", input: "十", want: 10},
		{name: "十一", input: "十一", want: 11},
		{name: "十九", input: "十九", want: 19},
		{name: "二十", input: "二十", want: 20},
		{name: "二十一", input: "二十一", want: 21},
		{name: "二十九", input: "二十九", want: 29},
		{name: "三十", input: "三十", want: 30},
		{name: "三十一", input: "三十一", want: 31},
		{name: "三十九", input: "三十九", want: 39},
		{name: "四十", input: "四十", want: 40},
		{name: "四十一", input: "四十一", want: 41},
		{name: "四十九", input: "四十九", want: 49},
		{name: "五十", input: "五十", want: 50},
		{name: "五十一", input: "五十一", want: 51},
		{name: "五十九", input: "五十九", want: 59},
		{name: "六十", input: "六十", want: 60},
		{name: "六十一", input: "六十一", want: 61},
		{name: "六十九", input: "六十九", want: 69},
		{name: "七十", input: "七十", want: 70},
		{name: "七十一", input: "七十一", want: 71},
		{name: "七十九", input: "七十九", want: 79},
		{name: "八十", input: "八十", want: 80},
		{name: "八十一", input: "八十一", want: 81},
		{name: "八十九", input: "八十九", want: 89},
		{name: "九十", input: "九十", want: 90},
		{name: "九十一", input: "九十一", want: 91},
		{name: "九十九", input: "九十九", want: 99},
		{name: "一百", input: "一百", want: 100},
		{name: "一百零一", input: "一百零一", want: 101},
		{name: "一百九十九", input: "一百九十九", want: 199},
		{name: "一千", input: "一千", want: 1000},
		{name: "一千零一", input: "一千零一", want: 1001},
		{name: "一千九百九十九", input: "一千九百九十九", want: 1999},
		{name: "一万", input: "一万", want: 10000},
		{name: "一万零一", input: "一万零一", want: 10001},
		{name: "一万九千九百九十九", input: "一万九千九百九十九", want: 19999},
		{name: "十万", input: "十万", want: 100000},
		{name: "十万零一", input: "十万零一", want: 100001},
		{name: "十九万九千九百九十九", input: "十九万九千九百九十九", want: 199999},
		{name: "一百万", input: "一百万", want: 1000000},
		{name: "一百万零一", input: "一百万零一", want: 1000001},
		{name: "一百九十九万九千九百九十九", input: "一百九十九万九千九百九十九", want: 1999999},
		{name: "一千万", input: "一千万", want: 10000000},
		{name: "一千万零一", input: "一千万零一", want: 10000001},
		{name: "一千九百九十九万九千九百九十九", input: "一千九百九十九万九千九百九十九", want: 19999999},
		{name: "一亿", input: "一亿", want: 100000000},
		{name: "一亿零一", input: "一亿零一", want: 100000001},
		{name: "一亿九千九百九十九万九千九百九十九", input: "一亿九千九百九十九万九千九百九十九", want: 199999999},
		{name: "十亿", input: "十亿", want: 1000000000},
		{name: "十亿零一", input: "十亿零一", want: 1000000001},
		{name: "十九亿九千九百九十九万九千九百九十九", input: "十九亿九千九百九十九万九千九百九十九", want: 1999999999},
		{name: "一百亿", input: "一百亿", want: 10000000000},
		{name: "一百亿零一", input: "一百亿零一", want: 10000000001},
		{name: "一百九十九亿九千九百九十九万九千九百九十九", input: "一百九十九亿九千九百九十九万九千九百九十九", want: 19999999999},
		{name: "19999999999", input: "19999999999", want: 19999999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := StringToNumber(tt.input)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("StringToNumber rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}
