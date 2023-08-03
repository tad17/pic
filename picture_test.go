package pic

import (
	//"log"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPicture(t *testing.T) {
	// для тестирования базу данных не учитываем
	p := NewPicture("альбом")
	assert.Equal(t, p.Album, "альбом", "должны быть одинаковыми")
}

func TestBaseName(t *testing.T) {
	base := baseName("c:\\Проверка файла с каталогом\\test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("/каталог файла/test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("/test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("test.jpg")
	assert.NotEqual(t, base, "неверное имя", "должны быть разными")
}

func TestCmdFfmpeg(t *testing.T) {
	// правильный файл
	cmd, out, err := cmdFfmpeg("testdir\\test.gif")
	log.Printf("cmd: %s\n", cmd)

	assert.Nil(t, err)
	assert.Equal(t, out, "testdir\\test-2x.webp", "должны быть одинаковыми")
	assert.Contains(t, cmd, "ffmpeg", "команда должна сформироваться")

	// не верный файл
	cmd, out, err = cmdFfmpeg("test.jpg")
	assert.Error(t, err, "должна быть ошибка")
	assert.Equal(t, out, "", "должен быть пустой строкой")
	assert.Equal(t, cmd, "", "должен быть пустой строкой")
}

func TestGetSize(t *testing.T) {
	out := `
	Stream #0:0: Video: webp <какой-то текст> 1200x1980 <еще какой-то текст>
	`
	w, h := getSize(out)
	assert.Equal(t, w, "1200", "должны быть 1200")
	assert.Equal(t, h, "1980", "должны быть 1980")
}

func TestConvertFile(t *testing.T) {
	// предварительно скопируем тестовый файл во временный каталог
	tmpgif := "test.gif"
	require.FileExists(t, tmpgif)

	// создадим временный каталог
	tmpdir := t.TempDir()
	filename := filepath.Join(tmpdir, tmpgif)

	// скопируем исходный файл в этот каталог
	err := copyFile(tmpgif, filename)
	require.NoError(t, err)
	require.FileExists(t, filename)

	// собственно конвертация
	webp, err := convertFile(filename)
	assert.NoError(t, err)
	require.FileExists(t, webp.filename)

	// скопируем обратно, что бы посмотреть, что получилось
	// err = copyFile(webp.filename, "d:/projects/pic/test.webp")
	// assert.NoError(t, err)
}

func TestConvert(t *testing.T) {
	// проверяем, что после конвертации файлы удаляются
	tmpgif := "test.gif"
	require.FileExists(t, tmpgif)

	err := Convert(tmpgif)
	assert.NoError(t, err)
}
