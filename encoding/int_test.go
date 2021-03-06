package encoding

import (
	"bytes"
	"testing"
)

func TestWriteInt(t *testing.T) {
	tests := []struct {
		Value    Int
		Expected []byte
	}{
		{Value: 0, Expected: []byte{0x00, 0x00, 0x00, 0x00}},
		{Value: 5, Expected: []byte{0x00, 0x00, 0x00, 0x05}},
		{Value: 2147483647, Expected: []byte{0x7f, 0xff, 0xff, 0xff}},
		{Value: -1, Expected: []byte{0xff, 0xff, 0xff, 0xff}},
		{Value: -2147483648, Expected: []byte{0x80, 0x00, 0x00, 0x00}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		if err := test.Value.Write(&buff); err != nil {
			t.Error(err)
		}

		actual := buff.Bytes()

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf("Unable to convert %d: %v != %v", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}

func TestReadInt(t *testing.T) {
	tests := []struct {
		Expected Int
		Value    []byte
	}{
		{Expected: 0, Value: []byte{0x00, 0x00, 0x00, 0x00}},
		{Expected: 5, Value: []byte{0x00, 0x00, 0x00, 0x05}},
		{Expected: 2147483647, Value: []byte{0x7f, 0xff, 0xff, 0xff}},
		{Expected: -1, Value: []byte{0xff, 0xff, 0xff, 0xff}},
		{Expected: -2147483648, Value: []byte{0x80, 0x00, 0x00, 0x00}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		var actual Int
		if err := actual.Read(&buff); err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf("Unable to convert %v: %d != %d", test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
