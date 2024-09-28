package btree

import (
	"os"
	"testing"
)

func TestLoadPage(t *testing.T) {
	dbFile, err := os.Open("../../superheroes.db")
	if err != nil {
		t.Errorf(`Error Opening db file`)
	}
	defer dbFile.Close()

	l := new(interiorTablePage)
	err = l.loadPageFromFile(dbFile, 4096)
	if err != nil {
		t.Errorf(`Error Loading Page`)
	}

	t.Logf("\n%+v\n", *l)
}
