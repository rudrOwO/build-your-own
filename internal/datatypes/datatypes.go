package datatypes

// Based on: https://www.sqlite.org/fileformat2.html#varint
func VARINT(bytes []byte) (uint64, int) {
	var result uint64
	for i, b := range bytes {
		result <<= 7
		result |= uint64(b & 0x7f)
		if b&0x80 == 0 {
			return result, i + 1
		}
	}
	return result, 0
}

// Based on: https://www.sqlite.org/fileformat.html#record_format
func INTEGER[T int8 | int16 | int32 | int64](serialType uint64, bytes []byte) T {
	return 8
}
