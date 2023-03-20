package helper

import (
	"math"
	"strconv"
)

// from https://zhuanlan.zhihu.com/p/489654080

type chNumberLevel struct {
	Val int
	/*
		中文数字节位分权
			大权位： 万，亿……
			小权位； 十，百，千
	*/
	IsPower bool
}

// 权位对应表
var chPowerNumMap = map[string]*chNumberLevel{
	"亿": {
		int(math.Pow10(8)),
		true,
	},
	"万": {
		int(math.Pow10(4)),
		true,
	},
	"千": {
		int(math.Pow10(3)),
		false,
	},
	"百": {
		int(math.Pow10(2)),
		false,
	},
	"十": {
		int(math.Pow10(1)),
		false,
	},
}

// 数值对应表
var ChToArNumberMap = map[string]int{
	"零": 0,
	"一": 1,
	"二": 2,
	"三": 3,
	"四": 4,
	"五": 5,
	"六": 6,
	"七": 7,
	"八": 8,
	"九": 9,
}

// chPowerNumber 获取权位的信息
func chPowerNumber(inputStr string) (int, bool) {
	if powerNum, exist := chPowerNumMap[inputStr]; exist {
		return powerNum.Val, powerNum.IsPower
	}
	return 0, false
}

// convertChNumberToArNumber 中文数字转阿拉伯数字
func convertChNumberToArNumber(inputStrNum string) (ansNum int) {
	chNum := []rune(inputStrNum)
	var (
		CurNumber int
		SumNumber int
	)
	for index := 0; index < len(chNum); index++ {
		// 将中文转为阿拉伯数字
		var getNum = ChToArNumberMap[string(chNum[index])]
		// 如果转换失败 getNum = -1
		if getNum > 0 {
			// 处理九九八专用
			if CurNumber != 0 {
				CurNumber *= int(math.Pow10(1))
				CurNumber += getNum
			} else {
				CurNumber = getNum
			}

			// 如果队列结束，则终止循环，并将临时数据添加到ansNum
			if index == len(chNum)-1 {
				SumNumber += CurNumber
				ansNum += SumNumber
				break
			}
		} else {
			// getNum 等于 -1 或 0 进入这里
			// 没有零对应的权，所以这一步所得出的大概率是上面提到的权位信息
			powerNum, isPower := chPowerNumber(string(chNum[index]))

			// 如果是大权位，则准备下一个权区间
			if isPower {
				SumNumber = (SumNumber + CurNumber) * powerNum
				ansNum, SumNumber = ansNum+SumNumber, 0
			} else {
				// 如果是小权位，则将当前数字添加到临时数和中
				if CurNumber != 0 {
					SumNumber += CurNumber * powerNum
				} else {
					SumNumber += powerNum
				}
			}
			CurNumber = 0
			// 如果队列结束，则终止循环，并将临时数据添加到ansNum
			if index == len(chNum)-1 {
				ansNum += SumNumber
				break
			}
		}
	}
	return
}

// StringToNumber 字符串转整型，兼容阿拉伯和中文
//
// TODO:支持中文和阿拉伯混合场景
func StringToNumber(str string) int64 {
	if num, err := strconv.ParseInt(str, 10, 64); err == nil {
		return num
	}
	num := convertChNumberToArNumber(str)
	return int64(num)
}
