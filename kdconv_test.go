package kdconv

import (
	"testing"
	"time"
)

func Test_Bool(t *testing.T) {
	convData := Bool(nil)
	if false != convData {
		t.Error("convData:", convData, "except:", false)
	}

	dataBool := bool(true)
	convData = Bool(dataBool)
	if dataBool != convData {
		t.Error("convData:", convData, "except:", dataBool)
	}

	defBool := true
	convData = BoolDef(nil, defBool)
	if defBool != convData {
		t.Error("convData:", convData, "except:", defBool)
	}
}

func Test_Int(t *testing.T) {
	convData := Int(nil)
	if int(0) != convData {
		t.Error("convData:", convData, "except:", 0)
	}

	dataInt := int(100)
	convData = Int(dataInt)
	if int(dataInt) != convData {
		t.Error("convData:", convData, "except:", dataInt)
	}

	dataInt8 := int8(101)
	convData = Int(dataInt8)
	if int(dataInt8) != convData {
		t.Error("convData:", convData, "except:", dataInt8)
	}

	dataInt16 := int16(102)
	convData = Int(dataInt16)
	if int(dataInt16) != convData {
		t.Error("convData:", convData, "except:", dataInt16)
	}

	dataInt32 := int32(103)
	convData = Int(dataInt32)
	if int(dataInt32) != convData {
		t.Error("convData:", convData, "except:", dataInt32)
	}

	dataInt64 := int64(104)
	convData = Int(dataInt64)
	if int(dataInt64) != convData {
		t.Error("convData:", convData, "except:", dataInt64)
	}

	defInt := int(105)
	convData = IntDef(nil, defInt)
	if defInt != convData {
		t.Error("convData:", convData, "except:", defInt)
	}
}

func Test_Int64(t *testing.T) {
	convData := Int64(nil)
	if int64(0) != convData {
		t.Error("convData:", convData, "except:", 0)
	}

	dataInt := int(100)
	convData = Int64(dataInt)
	if int64(dataInt) != convData {
		t.Error("convData:", convData, "except:", dataInt)
	}

	dataInt8 := int8(101)
	convData = Int64(dataInt8)
	if int64(dataInt8) != convData {
		t.Error("convData:", convData, "except:", dataInt8)
	}

	dataInt16 := int16(102)
	convData = Int64(dataInt16)
	if int64(dataInt16) != convData {
		t.Error("convData:", convData, "except:", dataInt16)
	}

	dataInt32 := int32(103)
	convData = Int64(dataInt32)
	if int64(dataInt32) != convData {
		t.Error("convData:", convData, "except:", dataInt32)
	}

	dataInt64 := int64(104)
	convData = Int64(dataInt64)
	if int64(dataInt64) != convData {
		t.Error("convData:", convData, "except:", dataInt64)
	}

	defInt64 := int64(105)
	convData = Int64Def(nil, defInt64)
	if defInt64 != convData {
		t.Error("convData:", convData, "except:", defInt64)
	}
}

func Test_Float(t *testing.T) {
	convData := Float(nil)
	if float32(0) != convData {
		t.Error("convData:", convData, "except:", 0)
	}

	dataInt := int(100)
	convData = Float(dataInt)
	if float32(dataInt) != convData {
		t.Error("convData:", convData, "except:", dataInt)
	}

	dataInt8 := int8(101)
	convData = Float(dataInt8)
	if float32(dataInt8) != convData {
		t.Error("convData:", convData, "except:", dataInt8)
	}

	dataInt16 := int16(102)
	convData = Float(dataInt16)
	if float32(dataInt16) != convData {
		t.Error("convData:", convData, "except:", dataInt16)
	}

	dataInt32 := int32(103)
	convData = Float(dataInt32)
	if float32(dataInt32) != convData {
		t.Error("convData:", convData, "except:", dataInt32)
	}

	dataInt64 := int64(104)
	convData = Float(dataInt64)
	if float32(dataInt64) != convData {
		t.Error("convData:", convData, "except:", dataInt64)
	}

	defFloat := float32(105)
	convData = FloatDef(nil, defFloat)
	if defFloat != convData {
		t.Error("convData:", convData, "except:", defFloat)
	}
}

func Test_String(t *testing.T) {
	convData := String(nil)
	if "" != convData {
		t.Error("convData:", convData, "except:", "")
	}

	dataString := string("test")
	convData = String(dataString)
	if dataString != convData {
		t.Error("convData:", convData, "except:", dataString)
	}

	defString := string("test default")
	convData = StringDef(nil, defString)
	if defString != convData {
		t.Error("convData:", convData, "except:", dataString)
	}
}

func Test_Time(t *testing.T) {
	convData := Time(nil)
	parsedtime := time.Time{}
	if parsedtime != convData {
		t.Error("convData:", convData, "except:", parsedtime)
	}

	dataTime := string("2013-07-02T23:13:00Z")
	parsedtime, err := time.Parse("2006-01-02T15:04:05Z", dataTime)
	if err != nil {
		panic(err)
	}
	convData = Time(dataTime)
	if parsedtime != convData {
		t.Error("convData:", convData, "except:", dataTime)
	}

	defTime := parsedtime
	convData = TimeDef(nil, parsedtime)
	if defTime != convData {
		t.Error("convData:", convData, "except:", defTime)
	}
}
