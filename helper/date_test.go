package helper

import (
	"reflect"
	"testing"
	"time"
)

func TestGetNextWeekday(t *testing.T) {

	type args struct {
		weekday time.Weekday
		base    time.Time
	}
	tests := []struct {
		name  string
		input args
		want  time.Time
	}{
		{
			name: "same day",
			input: args{
				weekday: time.Sunday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Monday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Tuesday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Wednesday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Thursday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Friday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "same week day",
			input: args{
				weekday: time.Saturday,
				base:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 7, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross week day",
			input: args{
				weekday: time.Sunday,
				base:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross month day",
			input: args{
				weekday: time.Monday,
				base:    time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 2, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross year day",
			input: args{
				weekday: time.Friday,
				base:    time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := GetNextWeekday(tt.input.weekday, tt.input.base)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("GetNextWeekday rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}

func TestGetNextDay(t *testing.T) {

	type args struct {
		day  int
		base time.Time
	}
	tests := []struct {
		name  string
		input args
		want  time.Time
	}{
		{
			name: "same day",
			input: args{
				day:  1,
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "interval of 1 day",
			input: args{
				day:  2,
				base: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross 1 month",
			input: args{
				day:  1,
				base: time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross 2 month",
			input: args{
				day:  29,
				base: time.Date(2023, 1, 31, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 3, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "cross 1 year",
			input: args{
				day:  10,
				base: time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "in february",
			input: args{
				day:  30,
				base: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2023, 3, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret := GetNextDay(tt.input.day, tt.input.base)
			if !reflect.DeepEqual(tt.want, ret) {
				t.Errorf("GetNextDay rsp got = %v, want %v", ret, tt.want)
			}
		})
	}
}
