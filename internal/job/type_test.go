package job

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper functions remain unchanged at the top
func stringPtr(s string) *string { return &s }
func intPtr(i int) *int { return &i }
func float64Ptr(f float64) *float64 { return &f }

func TestDefaultValue_IsNull(t *testing.T) {
	tests := map[string]struct {
		name string
		d    DefaultValue
		want bool
	}{
		"null value": {
			d:    DefaultValue{},
			want: true,
		},
		"string value": {
			d:    DefaultValue{strVal: stringPtr("test")},
			want: false,
		},
		"int value": {
			d:    DefaultValue{intVal: intPtr(42)},
			want: false,
		},
		"float value": {
			d:    DefaultValue{floatVal: float64Ptr(3.14)},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.IsNull())
		})
	}
}

func TestDefaultValue_MarshalJSON(t *testing.T) {
	tests := map[string]struct {
		d       DefaultValue
		want    string
		wantErr bool
	}{
		"string value": {
			d:    DefaultValue{strVal: stringPtr("test")},
			want: `"test"`,
		},
		"int value": {
			d:    DefaultValue{intVal: intPtr(42)},
			want: `42`,
		},
		"float value": {
			d:    DefaultValue{floatVal: float64Ptr(3.14)},
			want: `3.14`,
		},
		"null value": {
			d:    DefaultValue{},
			want: `null`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(tt.d)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestDefaultValue_UnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		json    string
		want    DefaultValue
		wantErr bool
	}{
		"string value": {
			json: `"test"`,
			want: DefaultValue{strVal: stringPtr("test")},
		},
		"int value": {
			json: `42`,
			want: DefaultValue{intVal: intPtr(42)},
		},
		"float value": {
			json: `3.14`,
			want: DefaultValue{floatVal: float64Ptr(3.14)},
		},
		"null value": {
			json: `null`,
			want: DefaultValue{},
		},
		"invalid value": {
			json:    `{"invalid": "json"}`,
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			var got DefaultValue
			err := json.Unmarshal([]byte(tt.json), &got)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDefaultValue_String(t *testing.T) {
	tests := map[string]struct {
		d    DefaultValue
		want string
	}{
		"string value": {
			d:    DefaultValue{strVal: stringPtr("test")},
			want: "test",
		},
		"int value": {
			d:    DefaultValue{intVal: intPtr(42)},
			want: "42",
		},
		"float value": {
			d:    DefaultValue{floatVal: float64Ptr(3.14)},
			want: "3.14",
		},
		"null value": {
			d:    DefaultValue{},
			want: "null",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.String())
		})
	}
}

func TestDefaultValue_Float64(t *testing.T) {
	tests := map[string]struct {
		d       DefaultValue
		want    float64
		wantErr bool
	}{
		"float value": {
			d:    DefaultValue{floatVal: float64Ptr(3.14)},
			want: 3.14,
		},
		"int value": {
			d:    DefaultValue{intVal: intPtr(42)},
			want: 42,
		},
		"string value": {
			d:       DefaultValue{strVal: stringPtr("test")},
			wantErr: true,
		},
		"null value": {
			d:       DefaultValue{},
			wantErr: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.d.Float64()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
