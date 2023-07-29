package pic

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

//Picture работа с картинкаи
type Picture struct {
	ID          int    `json:"id"`
	Url         string `json:"src"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Thumb       string `json:"srct"`
	ThumbWidth  int    `db:"thumb_width" json:"imgtWidth"`
	ThumbHeight int    `db:"thumb_height" json:"imgtHeight"`

	// закрытые поля
	db    *sqlx.DB
	album string  
}

//Hello - приветствие и комментарии к модулю
func Hello() {
	fmt.Printf("pic - пакет для работы с картинками\n")
	fmt.Printf("v0.0.1 - начальная версия\n")
}

func NewPicture(db *sqlx.DB, album string) Picture {
	pic := Picture {db: db, album: album}
	return pic
}