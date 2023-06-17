package utils

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

/*
? @description: 进程check是否为空
? @param {*widget.Entry} input
? @return {*}
*/
func Check_input(input *widget.Entry) bool {
	if input.Text == "" {
		return false
	}

	return true
}

/*
? @description: 检查input中的输入内容是否符合时间字符串格式
? @param {string} str 时间字符串
? @return {*}  true表示符合格式  layout1满足条件之间返回结果,不在进行layout2(结果赋值为false)判断
*/
func Check_input_format(str string) (bool, bool) {
	layout1 := "2006-01-02 15:04:05"
	layout2 := "15:04:05"
	_, err := time.Parse(layout1, str)
	if err != nil {
		_, err := time.Parse(layout2, str)
		if err != nil {
			return false, false
		}
		return false, true
	}
	return true, false
}
