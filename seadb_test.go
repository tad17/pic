package pic

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db *sqlx.DB

func openDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "itman:X753951x@(xigmanas:3306)/sea-test")
	if err != nil {
		panic(err)
	}
	return db
}

// Первоначальная установка окружения
func TestMain(m *testing.M) {
	// Каталог test должен присутствовать и в нем должны находиться:
	// sea-test.sql - скрипт для создания тестовой БД и заполнения ее тестовыми данными (mock)
	// test.gif - тестовый gif (mock)

	//os.RemoveAll("test")
	//os.Mkdir("test", 0666)
	db = openDB()
	os.Exit(m.Run())
}

func TestGetSmallGifs(t *testing.T) {
	// db открыта на глобальном уровне
	seadb, err := NewSeaDB(db)
	// если ошибка - дальнейшая проверка бессмыслена
	require.NoError(t, err)

	urls, err := seadb.SmallGifs()
	assert.Nil(t, err)

	// из тестовой БД должны получить 3 записи
	assert.Equal(t, len(urls), 3, "должны получить ровно 3 записи")
}

// Проверка загрузки из weed
func TestDownload(t *testing.T) {
	// для удобства - теперь вместо assert(t...) пишем просто assert(..)
	assert := assert.New(t)
	url := "http://192.168.0.105:9090/1,15df3bf5a2b6cb"
	tmpdir := t.TempDir()

	// формируем полное имя файла с путем.
	filename := filepath.Join(tmpdir, "1,15df3bf5a2b6cb")
	err := download(url, filename)
	assert.Nil(err)
	assert.FileExists(filename)
}

// Работает, но не понятно как интерпретировать результаты
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
