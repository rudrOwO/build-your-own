// * Implemented using https://saveriomiroddi.github.io/SQLIte-database-file-format-diagrams/
package btree

import (
	"encoding/binary"

	"github.com/rudrowo/sqlite/internal/datatypes"
)

const (
	PAGE_SIZE                 = 4096
	SQLITE_SCHEMA_ROOT_OFFSET = 100

	INTERIOR_INDEX_PAGE_TYPE = 0x02
	LEAF_INDEX_PAGE_TYPE     = 0x0a

	INTERIOR_TABLE_PAGE_TYPE = 0x05
	LEAF_TABLE_PAGE_TYPE     = 0x0d
)

// Headers
type (
	leafHeader struct {
		PageType  uint8
		CellCount uint16
	}
	interiorHeader struct {
		pageType         uint8
		cellCount        uint16
		rightmostPointer uint32
	}
	recordHeader struct {
		ColumnTypes []uint64
		HeaderSize  uint64
	}
)

// Cells
type (
	interiorTableCell struct {
		leftChildPointer uint32
		rowId            uint64
	}
	leafTableCell struct {
		Payload struct {
			RecordBody []byte
			recordHeader
		}
		RowId       uint64
		PayloadSize uint64
	}
)

// Pages
type (
	interiorTablePage struct {
		cellPointers []uint16
		cells        []interiorTableCell
		header       interiorHeader
	}
	LeafTablePage struct {
		CellPointers []uint16
		Cells        []leafTableCell
		Header       leafHeader
	}
)

func (l *interiorTablePage) loadFromBuffer(fileBuffer []byte) {
	l.header.pageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.header.cellCount = binary.BigEndian.Uint16(fileBuffer[3:5])
	// Bytes ignored => [5:8]
	l.header.rightmostPointer = binary.BigEndian.Uint32(fileBuffer[8:12])

	l.cellPointers = make([]uint16, l.header.cellCount)
	l.cells = make([]interiorTableCell, l.header.cellCount)

	for i, j := 0, 12; i < int(l.header.cellCount); {
		l.cellPointers[i] = binary.BigEndian.Uint16(fileBuffer[j : j+2])
		// Load cell at i
		{
			ci := l.cellPointers[i]
			l.cells[i].leftChildPointer = binary.BigEndian.Uint32(fileBuffer[ci : ci+4])
			l.cells[i].rowId, _ = datatypes.VARINT(fileBuffer[ci+4:])
		}
		i += 1
		j += 2
	}
}

func (l *LeafTablePage) loadFromBuffer(fileBuffer []byte) {
	l.Header.PageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.Header.CellCount = binary.BigEndian.Uint16(fileBuffer[3:5])

	l.CellPointers = make([]uint16, l.Header.CellCount)

	for i, j := 0, 12; i < int(l.Header.CellCount); {

		i += 1
		j += 2
	}
}
