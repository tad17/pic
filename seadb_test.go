package pic

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db = openDB()

func openDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "itman:X753951x@(xigmanas:3306)/sea")
	if err != nil {
		panic(err)
	}
	return db
}

// создадим временный каталог для тестирования
func TestMain(m *testing.M) {
	os.RemoveAll("test")
	os.Mkdir("test", 0666)
	os.Exit(m.Run())
}

func TestGetSmallGifs(t *testing.T) {
	seadb, err := NewSeaDB(db)
	// если ошибка - дальнейшая проверка бессмыслена
	require.NoError(t, err)

	urls, err := seadb.SmallGifs()
	assert.Nil(t, err)
	assert.Equal(t, len(urls), 10, "должны получить ровно 10 записей")

	// теперь проверяем правильность формата url
	// должны быть в форме: http://192.168.0.105:9091/4,530f8ea86e3e.gif
	for _, u := range urls {
		assert.Equal(t, strings.HasPrefix(u, "http://192.168.0.105:9091/"), true, "должен начинаться с http://192.168.0.105:9091/")
		assert.Equal(t, strings.HasSuffix(u, ".gif"), true, "должен заканчиваться на .gif")
	}

	// проверим загрузку
	url := urls[1]
	// fid := baseName()
	filename, err := getFilename(url)
	assert.Nil(t, err)

	err = download(url, filepath.Join("test", filename))
	assert.Nil(t, err)
}

func BenchmarkDownload(b *testing.B) {
	seadb, err := NewSeaDB(db)
	if err != nil {
		panic(err)
	}
	urls, err := seadb.SmallGifs()
	if err != nil {
		panic(err)
	}

	for _, url := range urls {
		filename, err := getFilename(url)
		if err != nil {
			panic(err)
		}
		err = download(url, filepath.Join("test", filename))
		if err != nil {
			panic(err)
		}
	}
}
