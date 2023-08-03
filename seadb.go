package pic

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// SeaDB - картинки в БД
type SeaDB struct {
	db *sqlx.DB
}

func NewSeaDB(db *sqlx.DB) (*SeaDB, error) {
	if db == nil {
		return nil, fmt.Errorf("db = nil")
	}
	seadb := SeaDB{db: db}
	return &seadb, nil
}

// Возвращает список url-ов маленьких gif-ов из БД
func (sea *SeaDB) SmallGifs() ([]string, error) {
	cmd := `
		SELECT 
	    	newurl
		FROM
	    	filemeta
		JOIN 
			pic ON filemeta.id = pic.file_id 
		WHERE
    		ext = '.gif' 
    		and del = 0 
    		and filemeta.converted = 0 
    		and pic.width < 400
	`
	var urls []string
	err := sea.db.Select(&urls, cmd)
	if err != nil {
		return nil, err
	}

	return urls, nil
}
