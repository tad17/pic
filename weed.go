package pic

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	//"log"
	"github.com/linxGnu/goseaweedfs"
)

var sw *goseaweedfs.Seaweed

func init() {
	masterURL := "http://192.168.0.105:9333"
	filer := []string{}
	sw, _ = goseaweedfs.NewSeaweed(
		masterURL, // master server
		filer,     // ну нету у меня filer
		8096,      // размер chunk
		&http.Client{Timeout: 5 * time.Minute})

	// просто проверим, что weed запущен
	_, err := sw.Status()
	if err != nil {
		panic(err)
	}
}

// загружает файл filename на weed сервер и возвращает url загруженного файла
func upload(filename string) (string, error) {
	fid, err := sw.Submit(filename, "", "")
	if err != nil {
		return "", err
	}
	return fid.FileURL, nil
}

// удаляет файл из weed
func delete(url string) error {
	fid, err := getFID(url)
	if err != nil {
		return err
	}

	err = sw.DeleteFile(fid, nil)
	return err
}

// загружает файл из weed в filename
// filename должен содержать полный путь
func download(url string, filename string) error {
	fid, err := getFID(url)
	if err != nil {
		return err
	}

	_, err = sw.Download(fid, nil, func(r io.Reader) (err error) {
		data, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(filename, os.O_CREATE, 0666)
		if err != nil {
			return err
		}

		_, err = file.Write(data)
		if err != nil {
			return err
		}

		err = file.Close()
		return err
	})

	return err
}

func getFID(url string) (string, error) {
	items := strings.Split(url, "/")
	if len(items) < 2 {
		return "", fmt.Errorf("url должен быть в формате: <http://address/fid.ext>")
	}
	return items[len(items)-1], nil
}

// возвращает имя файла из url без номера volume
// т.е. http://192.168.0.105:9090/1,15df3bf5a2b6cb => 15df3bf5a2b6cb
func getFilename(url string) (string, error) {
	fn, err := getFID(url)
	if err != nil {
		return "", err
	}

	items := strings.Split(fn, ",")
	return items[len(items)-1], nil
}
