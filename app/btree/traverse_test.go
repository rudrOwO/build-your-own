package btree

import "testing"

func TestGetRootPageOffset(t *testing.T) {
	schemaRootOffset := getRootPageOffset("sqlite_schema")

	if schemaRootOffset != 100 {
		t.Errorf(`Test Failed for TestGetRootPageOffset
	offset found: %d
	`, schemaRootOffset)
	}
}
