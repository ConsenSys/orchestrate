package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
	err "gitlab.com/ConsenSys/client/fr/core-stack/pkg.git/types/error"
)

type unmarshalTest struct {
	in     string
	ptr    interface{}
	errMsg string
}

var unmarshalTests = []unmarshalTest{
	{in: `true`, ptr: new(bool)},
	{in: `1`, ptr: new(int)},
	{in: `{"x": 1}`, ptr: new(bool), errMsg: "json: cannot unmarshal object into Go value of type bool"},
}

func TestUnmarshal(t *testing.T) {
	for _, test := range unmarshalTests {
		in := []byte(test.in)
		e := Unmarshal(in, test.ptr)
		if test.errMsg == "" {
			assert.Nil(t, e, "Unmarshal should not error")
		} else {
			e, ok := e.(*err.Error)
			assert.NotNil(t, e, "Unmarshal should error")
			assert.True(t, ok, "Error should be internal format")
			assert.Equal(t, []byte{0x10, 0x00}, e.GetCode(), "Error code should be correct")
			assert.Equal(t, "encoding.json", e.GetComponent(), "Error code should be correct")
			assert.Equal(t, test.errMsg, e.GetMessage(), "Error message should be correct")
		}
	}
}
