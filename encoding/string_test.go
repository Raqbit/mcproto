package encoding

import (
	"bytes"
	"testing"
)

func TestWriteString(t *testing.T) {
	tests := []struct {
		Value    string
		Expected []byte
	}{
		{Value: "john", Expected: []byte{0x04, 0x6a, 0x6f, 0x68, 0x6e}},
		{Value: " doe ", Expected: []byte{0x05, 0x20, 0x64, 0x6f, 0x65, 0x20}},
		{Value: "😂😂😂", Expected: []byte{0x0c, 0xf0, 0x9f, 0x98, 0x82, 0xf0, 0x9f, 0x98, 0x82, 0xf0, 0x9f, 0x98, 0x82}},
		{Value: "(╯°Д°）╯︵/(.□ . \\)", Expected: []byte{0x1e, 0x28, 0xe2, 0x95, 0xaf, 0xc2, 0xb0, 0xd0, 0x94, 0xc2, 0xb0, 0xef, 0xbc, 0x89, 0xe2, 0x95, 0xaf, 0xef, 0xb8, 0xb5, 0x2f, 0x28, 0x2e, 0xe2, 0x96, 0xa1, 0x20, 0x2e, 0x20, 0x5c, 0x29}},
	}

	for _, test := range tests {
		actual := WriteString(test.Value)

		if bytes.Compare(test.Expected, actual) != 0 {
			// Not equal
			t.Errorf(`Unable to convert "%s": %v != %v`, test.Value, actual, test.Expected)
		}
	}
}

func TestReadString(t *testing.T) {
	tests := []struct {
		Expected string
		Value    []byte
	}{
		{Expected: "john", Value: []byte{0x04, 0x6a, 0x6f, 0x68, 0x6e}},
		{Expected: " doe ", Value: []byte{0x05, 0x20, 0x64, 0x6f, 0x65, 0x20}},
		{Expected: "😂😂😂", Value: []byte{0x0c, 0xf0, 0x9f, 0x98, 0x82, 0xf0, 0x9f, 0x98, 0x82, 0xf0, 0x9f, 0x98, 0x82}},
		{Expected: "(╯°Д°）╯︵/(.□ . \\)", Value: []byte{0x1e, 0x28, 0xe2, 0x95, 0xaf, 0xc2, 0xb0, 0xd0, 0x94, 0xc2, 0xb0, 0xef, 0xbc, 0x89, 0xe2, 0x95, 0xaf, 0xef, 0xb8, 0xb5, 0x2f, 0x28, 0x2e, 0xe2, 0x96, 0xa1, 0x20, 0x2e, 0x20, 0x5c, 0x29}},
	}

	var buff bytes.Buffer

	for _, test := range tests {
		buff.Write(test.Value)

		actual, err := ReadString(&buff, 100)

		if err != nil {
			t.Error(err)
		}

		if actual != test.Expected {
			// Not equal
			t.Errorf(`Unable to convert %v: "%s" != "%s"`, test.Value, actual, test.Expected)
		}

		buff.Reset()
	}
}
