package hutils

import (
	"errors"
	"fmt"
	"time"
)

const (
	// DateLayout 日期序列化.
	DateLayout = "2006-01-02"
	// DateTimeLayout 时间序列化
	DateTimeLayout = "2006-01-02T15:04:05"
)

var TimeFormat = map[string]string{
	"yyyy-mm-dd HH:MM:SS": "2006-01-02 15:04:05",
	"yyyy-mm-dd HH:MM":    "2006-01-02 15:04",
	"yyyy-mm-dd HH":       "2006-01-02 15:04",
	"yyyy-mm-dd":          "2006-01-02",
	"yyyy-mm":             "2006-01",
	"mm-dd":               "01-02",
	"dd-mm-yy HH:MM:SS":   "02-01-06 15:04:05",
	"yyyy/mm/dd HH:MM:SS": "2006/01/02 15:04:05",
	"yyyy/mm/dd HH:MM":    "2006/01/02 15:04",
	"yyyy/mm/dd HH":       "2006/01/02 15",
	"yyyy/mm/dd":          "2006/01/02",
	"yyyy/mm":             "2006/01",
	"mm/dd":               "01/02",
	"dd/mm/yy HH:MM:SS":   "02/01/06 15:04:05",
	"yyyy":                "2006",
	"mm":                  "01",
	"HH:MM:SS":            "15:04:05",
	"MM:SS":               "04:05",
}

// Time Date time.Date的快捷方法，省略sec，nsec，loc.
func Time(year int, month time.Month, day, hour, min int) time.Time {
	return time.Date(year, month, day, hour, min, 0, 0, time.Local)
}

// DefaultParseTime 默认时间解析.
func DefaultParseTime(value string) (time.Time, error) {
	return ParseTime(DateTimeLayout, value)
}

// ParseTime 当地时区解析时间字符串.
func ParseTime(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

// ParseStrToTime 将str按照给定格式解析为time.Time（yyyy-mm-dd HH:MM:SS）。
func ParseStrToTime(str, format string) (time.Time, error) {
	v, ok := TimeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not found", format)
	}
	return time.ParseInLocation(v, str, time.Local)
}

// ParseTimeToStr time.Time转str。
func ParseTimeToStr(t time.Time, format string) string {
	return t.Format(TimeFormat[format])
}

// Period 时间区间.
type Period struct {
	Start, End time.Time
}

// ParseDateTimePeriod 解析时间区间.
func ParseDateTimePeriod(start, end string) (*Period, error) {
	if start == "" || end == "" {
		return nil, errors.New("查询区间不能为空")
	}
	startFrom, err := ParseTime(DateTimeLayout, start)
	if err != nil {
		return nil, errors.New("开始时间解析失败")
	}
	endTo, err := ParseTime(DateTimeLayout, end)
	if err != nil {
		return nil, errors.New("结束时间解析失败")
	}
	return &Period{Start: startFrom, End: endTo}, nil
}

// Tomorrow 明天同一时间.
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// Yesterday 昨天同一时间.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}
