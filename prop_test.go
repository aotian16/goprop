// prop_test.go
package goprop

import (
	"testing"
)

func TestNewProp(t *testing.T) {
	prop := NewProp()

	if prop == nil {
		t.Errorf("NewProp() return nil")
	}
}

func TestGet(t *testing.T) {
	prop := NewProp()

	prop["key"] = "value"

	_, e := prop.Get("errorkey")

	if e == nil {
		t.Errorf("Get() should return error when key not exists")
	}

	v1, e1 := prop.Get("key")

	if e1 != nil {
		t.Errorf("Get() return nil when key exists")
	}

	if v1 != "value" {
		t.Errorf("Get() return error value")
	}
}

func TestAtoi(t *testing.T) {
	prop := NewProp()

	prop["key"] = "123"

	_, e := prop.Atoi("errorkey")

	if e == nil {
		t.Errorf("Atoi() should return error when key not exists")
	}

	v1, e1 := prop.Atoi("key")

	if e1 != nil {
		t.Errorf("Atoi() return nil when key exists")
	}

	if v1 != 123 {
		t.Errorf("Atoi() return error value")
	}

	prop["key"] = "-123"

	v2, e2 := prop.Atoi("key")

	if e2 != nil {
		t.Errorf("Atoi() return nil when key exists")
	}

	if v2 != -123 {
		t.Errorf("Atoi() return error value")
	}

	prop["key"] = "0"

	v3, e3 := prop.Atoi("key")

	if e3 != nil {
		t.Errorf("Atoi() return nil when key exists")
	}

	if v3 != 0 {
		t.Errorf("Atoi() return error value")
	}

	prop["key"] = "xxx"

	_, e4 := prop.Atoi("key")

	if e4 == nil {
		t.Errorf("Atoi() should return error when Atoi error")
	}
}

func TestParseBool(t *testing.T) {
	prop := NewProp()

	prop["key"] = "True"

	_, e := prop.ParseBool("errorkey")

	if e == nil {
		t.Errorf("ParseBool() should return error when key not exists")
	}

	v1, e1 := prop.ParseBool("key")

	if e1 != nil {
		t.Errorf("ParseBool() return nil when key exists")
	}

	if v1 != true {
		t.Errorf("ParseBool() return error value")
	}

	prop["key"] = "False"

	v2, e2 := prop.ParseBool("key")

	if e2 != nil {
		t.Errorf("ParseBool() return nil when key exists")
	}

	if v2 != false {
		t.Errorf("ParseBool() return error value")
	}

	prop["key"] = "xxx"

	_, e3 := prop.ParseBool("key")

	if e3 == nil {
		t.Errorf("ParseBool() should return error when ParseBool error")
	}
}

func TestParseFloat(t *testing.T) {
	prop := NewProp()

	prop["key"] = "123.456"

	_, e := prop.ParseFloat("errorkey")

	if e == nil {
		t.Errorf("ParseFloat() should return error when key not exists")
	}

	v1, e1 := prop.ParseFloat("key")

	if e1 != nil {
		t.Errorf("ParseFloat() return nil when key exists")
	}

	if v1 != 123.456 {
		t.Errorf("ParseFloat() return error value")
	}

	prop["key"] = "-123.456"

	v2, e2 := prop.ParseFloat("key")

	if e2 != nil {
		t.Errorf("ParseFloat() return nil when key exists")
	}

	if v2 != -123.456 {
		t.Errorf("ParseFloat() return error value")
	}

	prop["key"] = "0.0"

	v3, e3 := prop.ParseFloat("key")

	if e3 != nil {
		t.Errorf("ParseFloat() return nil when key exists")
	}

	if v3 != 0.0 {
		t.Errorf("ParseFloat() return error value")
	}

	prop["key"] = "xxx"

	_, e4 := prop.ParseFloat("key")

	if e4 == nil {
		t.Errorf("ParseFloat() should return error when ParseFloat error")
	}
}

func TestParseInt(t *testing.T) {
	prop := NewProp()

	prop["key"] = "123"

	_, e := prop.ParseInt("errorkey")

	if e == nil {
		t.Errorf("ParseInt() should return error when key not exists")
	}

	v1, e1 := prop.ParseInt("key")

	if e1 != nil {
		t.Errorf("ParseInt() return nil when key exists")
	}

	if v1 != 123 {
		t.Errorf("ParseInt() return error value")
	}

	prop["key"] = "-123"

	v2, e2 := prop.ParseInt("key")

	if e2 != nil {
		t.Errorf("ParseInt() return nil when key exists")
	}

	if v2 != -123 {
		t.Errorf("ParseInt() return error value")
	}

	prop["key"] = "0"

	v3, e3 := prop.ParseInt("key")

	if e3 != nil {
		t.Errorf("ParseInt() return nil when key exists")
	}

	if v3 != 0 {
		t.Errorf("ParseInt() return error value")
	}

	prop["key"] = "xxx"

	_, e4 := prop.ParseInt("key")

	if e4 == nil {
		t.Errorf("ParseInt() should return error when ParseInt error")
	}
}

func TestParseUint(t *testing.T) {
	prop := NewProp()

	prop["key"] = "123"

	_, e := prop.ParseUint("errorkey")

	if e == nil {
		t.Errorf("ParseUint() should return error when key not exists")
	}

	v1, e1 := prop.ParseUint("key")

	if e1 != nil {
		t.Errorf("ParseUint() return nil when key exists")
	}

	if v1 != 123 {
		t.Errorf("ParseUint() return error value")
	}

	prop["key"] = "-123"

	_, e2 := prop.ParseUint("key")

	if e2 == nil {
		t.Errorf("ParseUint() should return error when ParseUint error")
	}

	prop["key"] = "0"

	v3, e3 := prop.ParseUint("key")

	if e3 != nil {
		t.Errorf("ParseUint() return nil when key exists")
	}

	if v3 != 0 {
		t.Errorf("ParseUint() return error value")
	}

	prop["key"] = "xxx"

	_, e4 := prop.ParseUint("key")

	if e4 == nil {
		t.Errorf("ParseUint() should return error when ParseUint error")
	}
}
