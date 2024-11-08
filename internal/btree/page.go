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
	LeafHeader struct {
		PageType   uint8
		CellsCount uint16
	}
	interiorHeader struct {
		LeafHeader
		rightmostPointer uint32
	}
	recordHeader struct {
		columnTypes []uint64
		headerSize  uint64
	}
)

// Cells
type (
	interiorTableCell struct {
		leftChildPointer uint32
		rowId            uint64
	}
	leafTableCell struct {
		// TODO
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
		Header       LeafHeader
	}
)

func (l *LeafTablePage) loadFromBuffer(fileBuffer []byte) {
	l.Header.PageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.Header.CellsCount = binary.BigEndian.Uint16(fileBuffer[3:5])

	// TODO load Cells
}

func (l *interiorTablePage) loadFromBuffer(fileBuffer []byte) {
	l.header.PageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.header.CellsCount = binary.BigEndian.Uint16(fileBuffer[3:5])
	// Bytes ignored => [5:8]
	l.header.rightmostPointer = binary.BigEndian.Uint32(fileBuffer[8:12])

	l.cellPointers = make([]uint16, l.header.CellsCount)
	l.cells = make([]interiorTableCell, l.header.CellsCount)

	for i, j := 0, 12; i < int(l.header.CellsCount); {
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
