package pic

import (
	"fmt"
	"strings"
	"github.com/jmoiron/sqlx"
)

//Picture - работа с картинками
type Picture struct {
	ID          int    `db:"id" json:"id"`
	Album       string `db:"album" json:"album"`
	Url         string `db:"pic_url" json:"src"`
	Width       int    `db:"pic_width" json:"width"`
	Height      int    `db:"pic_height" json:"height"`
	Thumb       string `db:"thumb_url" json:"srct"`
	ThumbWidth  int    `db:"thumb_width" json:"imgtWidth"`
	ThumbHeight int    `db:"thumb_height" json:"imgtHeight"`
}

//Hello - приветствие и комментарии к модулю
func Hello() {
	fmt.Printf("pic - пакет для работы с картинками\n")
	fmt.Printf("v0.0.1 - начальная версия\n")
}

func NewPicture(album string) Picture {
	pic := Picture {Album: album}
	return pic
}

// возвращает имя файла без расширения
func baseName(filename string) string {
	// на всякий случай заменим обратный слэш
	filename = strings.ReplaceAll(filename, "\\", "/")

	// проверка на наличие каталогов
	items := strings.Split(filename, "/")
	if len(items) > 0 {
		// есть каталоги, уберем
		filename = items[len(items)-1]
	} 
	fn := strings.Split(filename, ".")
	return fn[0]
}

