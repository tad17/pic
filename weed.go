package pic

import (
	"net/http"
	"time"
	"strings"
	"fmt"
	//"log"
	"github.com/linxGnu/goseaweedfs"
)

var sw *goseaweedfs.Seaweed

func init() {
	masterURL := "http://192.168.0.105:9333"
	filer := []string{}
	sw, _ = goseaweedfs.NewSeaweed(
		masterURL, // master server
		filer,     // нету у меня filer
		8096,      // размер chunk
		&http.Client{Timeout: 5 * time.Minute})

	_, err := sw.Status()
	if err != nil {
		panic(err)
	}
}

// загружает файл filename на weed и возвращает url загруженного файла
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

func getFID(url string) (string, error) {
	items := strings.Split(url, "/")
	if len(items) != 2  {
		return "", fmt.Errorf("url должен быть в формате: <address/fid>")
	}
	return items[len(items)-1], nil
}
