package utils

import (
	"strconv"
	"time"
)

// HowLongAgo 获取多久前时间
func HowLongAgo(timestamp int64) string {
	var second = time.Now().Unix() - timestamp
	var day = second/86400
	if day >= 1 {
		if day >= 365 {
			return strconv.FormatInt(day/365, 10) + "年前"
		}
		if day >= 30 {
			return strconv.FormatInt(day/30, 10) + "月前"
		}
		if day >= 7 {
			return strconv.FormatInt(day/7, 10) + "周前"
		}
		return strconv.FormatInt(day, 10) + "天前"
	}

	var hour = second/3600
	if hour >= 1 {
		return strconv.FormatInt(hour, 10) + "小时前"
	}

	var minute = second/60
	if minute >= 1 {
		return strconv.FormatInt(minute, 10) + "分钟前"
	}

	return strconv.FormatInt(second, 10) + "秒前"
}
