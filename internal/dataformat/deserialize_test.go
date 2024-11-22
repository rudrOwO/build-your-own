package dataformat

import (
	"testing"
)

func TestDeserializeVarint(t *testing.T) {
	var testVal uint64 = 0b1111101000
	buf := []byte{0b10000111, 0b01101000}
	decodedVal, bytesRead := DeserializeVarint(buf)

	if testVal != decodedVal {
		t.Errorf(`Bytes Read = %d
		Decoded Value = %b
		Actual Value = %b`,
			bytesRead, decodedVal, testVal)
	}
}

func TestDeserializeInteger(t *testing.T) {
	var testVal uint64 = 0b1111101000
	//
	buf := []byte{0b11, 0b11101000}
	decodedVal := DeserializeInteger(buf)

	if testVal != decodedVal {
		t.Errorf(`
		Decoded Value = %b 
		Actual Value = %b`,
			decodedVal, testVal)
	}
}
