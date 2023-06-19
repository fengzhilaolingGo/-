/*
! @Author: fengzhilaoling_Go fengzhilaoling_go@163.com
! @Date: 2023-06-17 10:15:37

	! @LastEditors: fengzhilaoling_Go fengzhilaoling_go@163.com
	! @LastEditTime: 2023-06-19 21:01:28

! @FilePath: \frne\utils\timeFunc.go
! @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
*/
package utils

import (
	"fmt"
	"log"
	"time"
)

func GetUTCTime(timeStr string) (time.Time, error) {
	date, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return time.Now(), err
	}
	date = date.UTC()
	return date, nil
}

func GetLosAngeles(t time.Time) string {
	lo, _ := t.Zone()
	if lo == time.UTC.String() {
		loc, _ := time.LoadLocation("America/Los_Angeles")
		pdtTime := t.In(loc)
		log.Println("pdttime", pdtTime)
		str := pdtTime.Format("2006-01-02 15:04:05 Z0700 MST")
		return str
	}
	return ""
}

func ConvertLosAngeles() string {
	t := time.Now()
	loc, _ := time.LoadLocation("America/Los_Angeles")
	pdtTime := t.In(loc)
	str := pdtTime.Format("2006-01-02 15:04:05")
	return str
}

func GetLocalTime(timeStr string) string {
	date, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return fmt.Sprintf(format string, a ...any)
	}
	str := date.Format("2006-01-02 15:04:05 Z0700 MST")
	return str
}
