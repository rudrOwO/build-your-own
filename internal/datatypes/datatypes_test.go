package datatypes

import (
	"testing"
)

func TestVARINT(t *testing.T) {
	var testVal uint64 = 0b1111101000
	buf := []byte{0b10000111, 0b01101000}
	decodedVal, bytesRead := VARINT(buf)

	if testVal != decodedVal {
		t.Errorf(`Bytes Read = %d
		Decoded Value = %b
		Actual Value = %b`,
			bytesRead, decodedVal, testVal)
	}
}
