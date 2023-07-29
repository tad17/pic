package pic

import (
	"strings"

	_ "github.com/go-sql-driver/mysql" // для связи с mysql
	"github.com/jmoiron/sqlx"
)

// возвращает список картинок одного альбома
func getAlbum(db *sqlx.DB, name string) ([]Picture, error) {
	cmd := "call getOneSeria(?)"
	var pics []Picture
	err := db.Select(&pics, cmd, name)
	return pics, err
}

// возвращает список альбомов 
func getAlbums(db *sqlx.DB) ([]Album, error) {
	cmd := "call getAlbums()"
	var albums []Album
	err := db.Select(&albums, cmd)
	return albums, err
}

// устанавливает признак удаленной картинки в БД
func deleteImagesDB(db *sqlx.DB, ids []int) error {
	cmd := "UPDATE filemeta SET del = 1 WHERE id = ?"
	for _, id := range ids {
		_, err := db.Exec(cmd, id)
		if err != nil {
			return err
		}
	}
	return nil
}

// перемещает картинки в альбом
func moveImagesDB(db *sqlx.DB, newseria, ids string) error {
	cmd := "UPDATE pic SET seria = ? WHERE file_id = ?"
	// fmt.Printf("перемещаю: %v в %s\n", ids, newseria)
	list := strings.Split(ids, ",")
	for _, id := range list {
		_, err := db.Exec(cmd, newseria, id)
		if err != nil {
			return err
		}
		// fmt.Printf(" ==перемещаю: %s в %s\n", id, newseria)
	}
	return nil
}
