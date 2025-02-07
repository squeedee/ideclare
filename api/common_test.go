package api_test

import (
	"encoding/json"
	"github.com/squeedee/ideclare/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MarshalUnmarshalFloat(t *testing.T) {
	v := &api.Value{}

	err := json.Unmarshal([]byte("3.14"), v)
	assert.NoError(t, err)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "3.14", string(r))
	assert.Equal(t, "float", v.Type())
	assert.False(t, v.IsNull())
}

func Test_MarshalUnmarshalInt(t *testing.T) {
	v := &api.Value{}

	err := json.Unmarshal([]byte("3"), v)
	assert.NoError(t, err)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "3", string(r))
	assert.Equal(t, "int", v.Type())
	assert.False(t, v.IsNull())
}

func Test_MarshalUnmarshalString(t *testing.T) {
	v := &api.Value{}

	err := json.Unmarshal([]byte("\"all the things\""), v)
	assert.NoError(t, err)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "\"all the things\"", string(r))
	assert.Equal(t, "string", v.Type())
	assert.False(t, v.IsNull())
}

func Test_MarshalUnmarshalNull(t *testing.T) {
	v := &api.Value{}

	err := json.Unmarshal([]byte("null"), v)
	assert.NoError(t, err)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "null", string(r))
	assert.Equal(t, "nil", v.Type())
	assert.True(t, v.IsNull())
}

// ----------------------------------------

func TestValue_NewFloat(t *testing.T) {
	v := api.NewFloatValue(3.14)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "3.14", string(r))
	assert.Equal(t, "float", v.Type())
	assert.False(t, v.IsNull())
}

func TestValue_NewInt(t *testing.T) {
	v := api.NewIntValue(3)

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "3", string(r))
	assert.Equal(t, "int", v.Type())
	assert.False(t, v.IsNull())
}

func TestValue_NewString(t *testing.T) {
	v := api.NewStringValue("all the things")

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "\"all the things\"", string(r))
	assert.Equal(t, "string", v.Type())
	assert.False(t, v.IsNull())
}

func TestValue_NewNull(t *testing.T) {
	v := api.NewNilValue()

	r, err := json.Marshal(v)

	assert.NoError(t, err)
	assert.Equal(t, "null", string(r))
	assert.Equal(t, "nil", v.Type())
	assert.True(t, v.IsNull())
}
