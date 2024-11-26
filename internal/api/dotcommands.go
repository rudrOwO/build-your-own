package api

import (
	"encoding/binary"
	"log"
)

func ReadPageSize() uint16 {
	dbHeader := make([]byte, 100)
	_, err := dbFile.Read(dbHeader)
	if err != nil {
		log.Fatal(err)
	}

	pageSize := binary.BigEndian.Uint16(dbHeader[16:18])
	return pageSize
}
