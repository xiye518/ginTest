package gormModel

import (
	"testing"
)

func TestCreatTable(t *testing.T) {
	db:=CreatTable()
	t.Log(db.Error)
	t.Log(db.Value)
	t.Log(db.RowsAffected)
	t.Log(db)
}
