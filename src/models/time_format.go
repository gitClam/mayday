package model

import (
	"database/sql/driver"
	"time"
	"log"
    "github.com/iris-contrib/schema"
    _ "github.com/iris-contrib/schema"
    "reflect"
    _ "unsafe"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time
var nt = LocalTime{}
func (n *LocalTime) Converter(s string) reflect.Value {
    if t, err := time.Parse(TimeFormat, s); err == nil {
        return reflect.ValueOf(LocalTime(t))
    }
    return reflect.Value{}
}

func RegisterLocalTimeDecoder() {	
	log.Print("123123")
	var nt = LocalTime{}
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

func (t LocalTime) String(Timeformat string) string {
	return time.Time(t).Format(Timeformat)
}