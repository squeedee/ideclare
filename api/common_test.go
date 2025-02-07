package api_test

import (
	"encoding/json"
	"github.com/squeedee/ideclare/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	tests := map[string]struct {
		json       string
		expectNull bool
	}{
		"float": {
			json: "3.14",
		},
		"int": {
			json: "3",
		},
		"string": {
			json: "\"hi friend\"",
		},
		"nil": {
			json:       "null",
			expectNull: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v := &api.Value{}

			err := json.Unmarshal([]byte(tc.json), v)
			assert.NoError(t, err)

			r, err := json.Marshal(v)

			assert.NoError(t, err)
			assert.Equal(t, tc.json, string(r))
			assert.Equal(t, name, v.Type())
			assert.Equal(t, tc.expectNull, v.IsNil())
		})
	}
}

func TestConstructors(t *testing.T) {
	tests := map[string]struct {
		constructor func() *api.Value
		expectValue string
		expectNull  bool
		fieldType   string
	}{
		"float": {
			constructor: func() *api.Value {
				return api.NewFloatValue(3.14)
			},
			fieldType:   "float",
			expectValue: "3.14",
		},
		"int": {
			constructor: func() *api.Value {
				return api.NewIntValue(3)
			},
			fieldType:   "int",
			expectValue: "3",
		},
		"string": {
			constructor: func() *api.Value {
				return api.NewStringValue("hi friend")
			},
			fieldType:   "string",
			expectValue: "\"hi friend\"",
		},
		"nil": {
			constructor: func() *api.Value {
				return api.NewNilValue()
			},
			fieldType:   "nil",
			expectValue: "null",
			expectNull:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v := tc.constructor()

			r, err := json.Marshal(v)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectValue, string(r))
			assert.Equal(t, tc.fieldType, v.Type())
			assert.Equal(t, tc.expectNull, v.IsNil())
		})
	}
}
