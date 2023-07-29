package pic

import (
)

//SeaDB - картинки в БД
type SeaDB struct {
	db *sqlx.DB
}

func NewSeaDB(db *qlx.DB) (*SeaDB, error) {
	seadb := SeaDB{db: db}
	return seadb
}

//Возвращает список url-ов маленьких gif-ов их БД
func (sea *SeaDB)SmallGifs() ([]string, error) {}
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
    		and a.converted = 0 
    		and b.width < 400
		LIMIT 
			10;
	`
	var urls []string
	err := sea.db.Select(&urls, cmd)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

