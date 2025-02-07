package api

import (
	"encoding/json"
	"fmt"
)

// Value represents a scalar of int, float or string in json/yaml
type Value struct {
	stringValue *string
	intVal      *int
	floatVal    *float64
}

func NewIntValue(val int) *Value {
	return &Value{
		intVal: &val,
	}
}

func NewFloatValue(val float64) *Value {
	return &Value{
		floatVal: &val,
	}
}

func NewStringValue(val string) *Value {
	return &Value{
		stringValue: &val,
	}
}

func NewNilValue() *Value {
	return &Value{}
}

func (d *Value) IsNil() bool {
	return d.stringValue == nil && d.intVal == nil && d.floatVal == nil
}

func (d *Value) MarshalJSON() ([]byte, error) {
	if d.stringValue != nil {
		return json.Marshal(d.stringValue)
	}

	if d.intVal != nil {
		return json.Marshal(*d.intVal)
	}

	if d.floatVal != nil {
		return json.Marshal(*d.floatVal)
	}

	return json.Marshal(nil)
}

func (d *Value) UnmarshalJSON(data []byte) error {
	// Handle null value
	if string(data) == "null" {
		return nil
	}

	// Try as number first
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		// Check if it's an integer by looking for decimal point
		s := string(data)
		if s[0] != '"' && len(s) > 0 { // ensure it's not a string
			hasDecimal := false
			for _, c := range s {
				if c == '.' {
					hasDecimal = true
					break
				}
			}
			if !hasDecimal {
				i := int(f)
				d.intVal = &i
				return nil
			}
		}
		d.floatVal = &f
		return nil
	}

	s := ""
	d.stringValue = &s
	// Try as string
	if err := json.Unmarshal(data, d.stringValue); err == nil {
		return nil
	}

	return fmt.Errorf("value must be either a number or string")
}

func (d *Value) Type() string {
	if d.stringValue != nil {
		return "string"
	}
	if d.intVal != nil {
		return "int"
	}
	if d.floatVal != nil {
		return "float"
	}
	return "nil"
}
