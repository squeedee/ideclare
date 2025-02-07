package job

import (
	"fmt"
	"sigs.k8s.io/yaml"
	"strconv"
)

// DefaultValue can store either a string or float64 value
type DefaultValue struct {
	strVal   *string
	intVal   *int
	floatVal *float64
}

// IsNull returns true if the value is null
func (d *DefaultValue) IsNull() bool {
	return d.strVal == nil && d.intVal == nil && d.floatVal == nil
}

// IsInt returns true if the value is an integer
func (d *DefaultValue) IsInt() bool {
	return d.intVal != nil
}

// MarshalJSON implements yaml.Marshaler
func (d *DefaultValue) MarshalJSON() ([]byte, error) {
	if d.strVal != nil {
		return yaml.Marshal(d.strVal)
	}

	if d.intVal != nil {
		return yaml.Marshal(*d.intVal)
	}

	if d.floatVal != nil {
		return yaml.Marshal(*d.floatVal)
	}

	return yaml.Marshal(nil)
}

// UnmarshalJSON implements yaml.Unmarshaler
func (d *DefaultValue) UnmarshalJSON(data []byte) error {
	// Handle null value
	if string(data) == "null" {
		return nil
	}

	// Try as number first
	var f float64
	if err := yaml.Unmarshal(data, &f); err == nil {
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

	// Try as string
	if err := yaml.Unmarshal(data, d.strVal); err == nil {
		return nil
	}

	return fmt.Errorf("value must be either a number or string")
}

// String returns the string representation
func (d *DefaultValue) String() string {
	if d.strVal != nil {
		return *d.strVal
	}
	if d.intVal != nil {
		return strconv.Itoa(*d.intVal)
	}
	if d.floatVal != nil {
		return strconv.FormatFloat(*d.floatVal, 'f', -1, 64)
	}
	return "null"
}

// Float64 returns the float64 value and whether the conversion was successful
func (d *DefaultValue) Float64() (float64, error) {
	if d.floatVal != nil {
		return *d.floatVal, nil
	}
	if d.intVal != nil {
		return float64(*d.intVal), nil
	}
	return 0, fmt.Errorf("value is not a number")
}

// First quandry; How do we verify input digests without making that the job's responsibility.
// In my mind, this is a vital lift. Job author says "I need this input, and I can trust that the system will capture its signature"
// Maybe the difference between "Env Vars" and these "digestable inputs" is that they are "Externals".

const Digestable_Value = "Value"
const Digestable_OciRegistry = "OCIRegistry"
const Digestable_UrlBody = "URLBody"

type Mapping struct {
}

type Input struct {
	// Name the input. This is also used as the default mapping to Environment variables, see Mapping.
	Name string `json:"name"`

	// Description is a free text definition of the input.
	Description string `json:"description"`

	// Type is a free text definition
	Type string `json:"type"`

	// Optional is true if this value does not need to be provided
	Optional bool `json:"optional"`

	// Default value that can be either string or number
	Default DefaultValue `json:"default"`

	// Mapping defines how the input is passed into the container
	Mapping Mapping `json:"mapping"`
}

type Inputs []Input

type Spec struct {
	Inputs Inputs `json:"env"`
}
