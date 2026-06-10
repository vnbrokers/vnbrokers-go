package numberic

import (
	"bytes"
	"strconv"
)

var (
	jsonNull  = []byte(`null`)
	jsonEmpty = []byte(`""`)
	jsonNaN   = []byte(`"NaN"`)
)

type NaNFloat64 float64

func (f *NaNFloat64) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, jsonNull) ||
		bytes.Equal(b, jsonEmpty) ||
		bytes.Equal(b, jsonNaN) {
		*f = 0
		return nil
	}

	if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	v, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}

	*f = NaNFloat64(v)
	return nil
}
