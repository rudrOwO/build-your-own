package api

import (
	"encoding/binary"
	"log"

	btree "github.com/rudrowo/sqlite/internal/btree"
)

func ReadPageSize() uint16 {
	dbHeader := make([]byte, btree.SQLITE3_HEADER_SIZE)
	_, err := dbFile.Read(dbHeader)
	if err != nil {
		log.Fatal(err)
	}

	pageSize := binary.BigEndian.Uint16(dbHeader[16:18])
	return pageSize
}
