package utils

import (
	"database/sql/driver"
	"github.com/iris-contrib/schema"
	_ "github.com/iris-contrib/schema"
	"reflect"
	"time"
	_ "unsafe"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

type timeDecoder struct{}

var (
	TimeDecoder *timeDecoder
	nt          = LocalTime{}
)

func (t *LocalTime) Converter(s string) reflect.Value {
	if t, err := time.Parse(TimeFormat, s); err == nil {
		return reflect.ValueOf(LocalTime(t))
	}
	return reflect.Value{}
}

func (t *timeDecoder) RegisterLocalTimeDecoder() {
	schema.Form.RegisterConverter(nt, nt.Converter)
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String(TimeFormat) == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String(TimeFormat string) string {
	return time.Time(t).Format(TimeFormat)
}
