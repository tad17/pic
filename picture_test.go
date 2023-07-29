package pic

import (
	//"log"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewPicture(t *testing.T) {
	// для тестирования базу данных не учитываем
	p := NewPicture("проверка")
	assert.Equal(t, p.Album, "проверка", "должны быть одинаковыми")
}

func TestbaseName(t *testing.T) {
	base := baseName("c:\\Проверка файла с каталогом\\test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("/Проверка файла с каталогом/test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("/test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("test.jpg")
	assert.Equal(t, base, "test", "должны быть одинаковыми")

	base = baseName("test.jpg")
	assert.NotEqual(t, base, "неверное имя", "должны быть разными")
}