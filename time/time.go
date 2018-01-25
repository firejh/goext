package gxtime

import (
	"time"
	"fmt"
	"strings"
)

//获取当前时间是一天的第几分钟后组成string｛2018_01_01_100｝
func GetYMDM() string {
	t := time.Now()
	y, m, d := t.Date()
	h := t.Hour()
	min := t.Minute()
	dmin := min + h * 60

	return fmt.Sprintf("%d_%d_%d_%d", y, m, d, dmin)
}

//｛2018_01_01_100｝->timestamp：YMD h:m:00
func GettimestampByYMDM(YMDM string) string {
	s := strings.Split(YMDM, "_")
	if len(s) < 4 {
		return ""
	}

	return fmt.Sprintf("%s%s%s %s:%s:00" , s[0], s[1], s[2], s[3])
}

//当前时间timestamp
func GetTimeStamp() string {
	t := time.Now()
	return t.Format("20060102 15:04:05")
}

func GetTimeStamp1() string {
	t := time.Now()
	return t.Format("20060102 15:04:05.000") + "+08:00"
}
