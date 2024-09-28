package btree

import (
	"encoding/binary"
	"os"
)

const (
	PAGE_SIZE                 = 4096
	SQLITE_SCHEMA_ROOT_OFFSET = 100

	INTERIOR_INDEX_PAGE_TYPE = 0x02
	LEAF_INDEX_PAGE_TYPE     = 0x0a

	INTERIOR_TABLE_PAGE_TYPE = 0x05
	LEAF_TABLE_PAGE_TYPE     = 0x0d
)

type leafHeader struct {
	pageType   uint8
	cellsCount uint16
}

type interiorHeader struct {
	leafHeader
	rightmostPointer uint32
}

type interiorTableCell struct {
	leftChildPointer uint32
	rowId            uint64
}

// A Page here is a Node in the B-Tree.
type interiorTablePage struct {
	header       interiorHeader
	cellPointers []uint16
	cells        []interiorTableCell
}

type LeafTablePage struct {
	header       leafHeader
	cellPointers []uint16
	// TODO Add Leaf Cells later
}

func (l *LeafTablePage) loadPageFromFile(dbFile *os.File, fileOffset int64) error {
	fileBuffer := make([]byte, PAGE_SIZE)
	_, err := dbFile.Seek(fileOffset, 0)
	if err != nil {
		return err
	}
	_, err = dbFile.Read(fileBuffer)
	if err != nil {
		return err
	}

	l.header.pageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.header.cellsCount = binary.BigEndian.Uint16(fileBuffer[3:5])
	// TODO load Cells

	return nil
}

func (l *interiorTablePage) loadPageFromFile(dbFile *os.File, fileOffset int64) error {
	fileBuffer := make([]byte, PAGE_SIZE)
	_, err := dbFile.Seek(fileOffset, 0)
	if err != nil {
		return err
	}
	_, err = dbFile.Read(fileBuffer)
	if err != nil {
		return err
	}

	l.header.pageType = fileBuffer[0]
	// Bytes ignored => [1:3]
	l.header.cellsCount = binary.BigEndian.Uint16(fileBuffer[3:5])
	// Bytes ignored => [5:8]
	l.header.rightmostPointer = binary.BigEndian.Uint32(fileBuffer[8:12])
	l.cellPointers = make([]uint16, l.header.cellsCount)

	for i, ci := 0, 12; i < int(l.header.cellsCount); {
		l.cellPointers[i] = binary.BigEndian.Uint16(fileBuffer[ci : ci+2])
		// TODO Load cell
		i += 1
		ci += 2
	}

	return nil
}
