package kdconv

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var timeZoneOffset time.Duration

func init() {
	_, offset := time.Now().Zone()
	timeZoneOffset = time.Duration(offset)
}

func BoolDef(data interface{}, def bool) bool {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case bool:
		return field.Bool()
	}
	panic(fmt.Errorf("Convert type failed"))
}

func Bool(data interface{}) bool {
	return BoolDef(data, false)
}

func IntDef(data interface{}, def int) int {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64:
		return int(field.Int())
	case float32, float64:
		return int(field.Float())
	case string:
		i, err := strconv.ParseInt(field.String(), 10, 32)
		if err != nil {
			panic(err)
		}
		return int(i)
	}
	panic(fmt.Errorf("Convert type failed"))
}

func Int(data interface{}) int {
	return IntDef(data, 0)
}

func Int64Def(data interface{}, def int64) int64 {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64:
		return field.Int()
	case float32, float64:
		return int64(field.Float())
	case string:
		i, err := strconv.ParseInt(field.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int64(i)
	}
	panic(fmt.Errorf("Convert type failed"))
}

func Int64(data interface{}) int64 {
	return Int64Def(data, 0)
}

func FloatDef(data interface{}, def float32) float32 {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64:
		return float32(field.Int())
	case float32, float64:
		return float32(field.Float())
	case string:
		f, err := strconv.ParseFloat(field.String(), 32)
		if err != nil {
			panic(err)
		}
		return float32(f)
	}
	panic(fmt.Errorf("Convert type failed"))
}

func Float(data interface{}) float32 {
	return FloatDef(data, 0)
}

func StringDef(data interface{}, def string) string {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprint(data)
	case string:
		return field.String()
	}
	panic(fmt.Errorf("Convert type failed"))
}

func String(data interface{}) string {
	return StringDef(data, "")
}

func TimeDef(data interface{}, def time.Time) time.Time {
	if data == nil {
		return def
	}
	field := reflect.ValueOf(data)
	switch field.Interface().(type) {
	case string:
		parsedtime, err := time.Parse("2006-01-02T15:04:05Z", field.String())
		if err == nil {
			return parsedtime
		}
		parsedtime, err = time.Parse("2006-01-02 03:04 PM", field.String())
		if err == nil {
			return parsedtime
		}
		parsedtime, err = time.Parse("2006-1-2 03:04 PM", field.String())
		if err == nil {
			return parsedtime
		}
		panic(err)
	}
	panic(fmt.Errorf("Convert type failed"))
}

func Time(data interface{}) time.Time {
	return TimeDef(data, time.Time{})
}

func TimeAddLocalZoneOffset(t time.Time) time.Time {
	return t.In(time.UTC).Add(timeZoneOffset * time.Second)
}

func TimeNowUTC() time.Time {
	return TimeAddLocalZoneOffset(time.Now())
}
