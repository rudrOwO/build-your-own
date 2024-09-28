package btree

/*
Traverse a table btree and enumerate list of
starting positions of leaf pages
*/
func LoadAllLeafTablePages(tableName string) (chan LeafTablePage, error) {
	return loadAllLeafTablePages(getRootPageOffset(tableName))
}

func getRootPageOffset(tableName string) int64 {
	if tableName == "sqlite_schema" || tableName == "sqlite_master" {
		return SQLITE_SCHEMA_ROOT_OFFSET
	} else {
		// TODO parse sqlite_schema table and find root offset
		return 0
	}
}

func loadAllLeafTablePages(rootPageOffset int64) (chan LeafTablePage, error) {
	leafTablesQueue := make(chan LeafTablePage, 1)
	// TODO write recursive page loader
	return leafTablesQueue, nil
}
