// * A Page here is a Node in the B-Tree.
// Helpful visualizer: https://saveriomiroddi.github.io/SQLIte-database-file-format-diagrams/
package btree

const (
	PAGE_SIZE = 4096

	INTERIOR_INDEX_PAGE_TYPE = 0x02
	LEAF_INDEX_PAGE_TYPE     = 0x0a

	INTERIOR_TABLE_PAGE_TYPE = 0x05
	LEAF_TABLE_PAGE_TYPE     = 0x0d
)

// ! Process the ignore fields from the docs when writing a reader/parser
type leafHeader struct {
	pageType   uint8
	cellsCount uint16
}

type interiorHeader struct {
	leafHeader
	rightmostPointer uint32
}

type tableInteriorCell struct {
	leftChildPointer uint32
	rowId            uint64
}

type tableInteriorPage struct {
	header       interiorHeader
	cellPointers []uint16
	cells        []tableInteriorCell
}

type tableLeafPage struct {
	header       leafHeader
	cellPointers []uint16
	// TODO Add Leaf Cells later
}
