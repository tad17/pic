package pic

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/estebangarcia21/subprocess"
	//"github.com/jmoiron/sqlx"
)

// Picture - работа с картинками
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

// Hello - приветствие и комментарии к модулю
func Hello() {
	fmt.Printf("pic - пакет для работы с картинками\n")
	fmt.Printf("v0.0.1 - начальная версия\n")
}

func NewPicture(album string) Picture {
	pic := Picture{Album: album}
	return pic
}

// возвращает имя файла без расширения
// т.е. из test/filename.jpg получаем filename
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

// cmdFfmpeg формирует команду ffmpeg для последующего исполнения
// ffmpeg - увеличивает gif в 2 раза и сразу преобразует в .webp
// filename - сам файл gif
// Возвращает команду ffmpeg и выходной файл webp
func cmdFfmpeg(src string) (string, string, error) {
	if !strings.HasSuffix(src, ".gif") {
		return "", "", fmt.Errorf("файл не .gif")
	}
	webp := strings.ReplaceAll(src, ".gif", "-2x.webp")
	cmd := fmt.Sprintf("ffmpeg -i \"%s\" -vf scale=iw*2:ih*2 -loop 0 \"%s\"", src, webp)
	return cmd, webp, nil
}

type WebpFile struct {
	filename string
	width    string
	height   string
}

// конвертация из .gif в .webp с изменением размера файла
// с использованием программы ffmpeg.
// filename - исходный файл gif (полный путь)
// Возвращает описание выходного файла WebpFile
// Выходной файл находится рядом с исходным
func convertFile(filename string) (WebpFile, error) {
	webpfile := WebpFile{}

	cmd, webp, err := cmdFfmpeg(filename)
	if err != nil {
		return webpfile, err
	}

	s := subprocess.New(cmd, subprocess.HideStderr)

	if err := s.Exec(); err != nil {
		return webpfile, fmt.Errorf("(%s) ffmpeg: %v", filename, err)
	}

	if s.ExitCode() != 0 {
		return webpfile, fmt.Errorf("(%s): ffmpeg exit code: %d", filename, s.ExitCode())
	}

	// из вывода команды ffmpeg получим новый размер файла webp
	w, h := getSize(s.StderrText())
	webpfile = WebpFile{
		filename: webp,
		width:    w,
		height:   h,
	}
	// fid := upload(newname)
	// println(fid)
	return webpfile, nil
}

// Экспорт функции конвертации
func Convert(filename string) error {
	webp, err := convertFile(filename)
	if err != nil {
		return err
	}

	log.Printf("[%sx%s] %s\n", webp.width, webp.height, webp.filename)

	// все сделано, удалим исходный и выходной файл
	if err := deleteFile(filename); err != nil {
		return err
	}

	if err := deleteFile(webp.filename); err != nil {
		return err
	}

	return nil
}

func deleteFile(filename string) error {
	err := os.Remove(filename)
	return err
}

// получает размер файла из вывода команды ffmpeg
// если пустая строка - бросаем панику
func getSize(out string) (string, string) {
	if out == "" {
		panic("пустая строка вывода ffmpeg")
	}
	lines := strings.Split(out, "\n")
	for _, s := range lines {
		if strings.Contains(s, "Stream #0:0: Video: webp") {
			// fmt.Printf("== %s\n", s)
			re := regexp.MustCompile(`(\d*)x(\d*)`)
			matches := re.FindAllStringSubmatch(s, -1)
			// fmt.Printf("matches: %v\n", matches)
			return matches[0][1], matches[0][2]
		}
	}
	return "0", "0"
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}
	return nil
}
