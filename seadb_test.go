package pic

import (
	//"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jmoiron/sqlx"
)

var db = openDB()

func openDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "itman:X753951x@(xigmanas:3306)/sea")
	if err != nil {
		panic(err)
	}
	return db
}

func TestGetSmallGifs(t *testing.T) {
	seadb, err := NewSeaDB(db)
	assert.Nil(t, err)

	urls, err := seadb.SmallGifs()
	assert.Nil(t, err)
	assert.Equal(t, len(urls), 10, "должны получить 10 записей")
}
