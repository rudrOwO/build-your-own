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

// TODO Process the ignore fields from the docs when writing a reader/parser
type leafHeader struct {
	pageType   uint8
	cellsCount uint16
}

type interiorHeader struct {
	leafHeader
	rightmostPointer uint32
}

// ? Make 4 types of cells
// TODO Make 4 pages in combination of the headers and cell arrays
