package time

import (
	"fmt"
	"time"

	sqldriver "database/sql/driver"

)

const (
	defaultDateTimeFormat = "2006-01-02 15:04:05"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(defaultDateTimeFormat))
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
	loc, _ := time.LoadLocation("Asia/Shanghai")
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
