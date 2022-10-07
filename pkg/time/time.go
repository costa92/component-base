package time

import (
	sqldriver "database/sql/driver"
	"fmt"
	"time"
)

const (
	defaultDateTimeFormat = "2006-01-02 15:04:05"
	defaultDateLocal      = "Asia/Shanghai"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%q\"", t.Format(defaultDateTimeFormat))
	return []byte(formatted), nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t Time) Value() (sqldriver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func ToTime(str string) (Time, error) {
	var jt Time
	loc, _ := time.LoadLocation(defaultDateLocal)
	value, err := time.ParseInLocation(defaultDateTimeFormat, str, loc)
	if err != nil {
		return jt, err
	}
	return Time{
		Time: value,
	}, nil
}

func Now() Time {
	return Time{
		Time: time.Now(),
	}
}

// FormatTime 时间戳格式化到上海时间
func FormatTime(t int64) string {
	loc, _ := time.LoadLocation(defaultDateLocal)
	return time.Unix(t, 0).In(loc).Format(defaultDateTimeFormat)
}

// UnixToStr 时间戳转时间
func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

// DateFormat 字符串日期格式化成time.Time
func DateFormat(strDate string) (time.Time, error) {
	timeLayout := defaultDateTimeFormat           // 转化所需模板
	loc, _ := time.LoadLocation(defaultDateLocal) // 获取时区
	return time.ParseInLocation(timeLayout, strDate, loc)
}
