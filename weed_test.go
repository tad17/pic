package pic

import (
	//"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	url, err := upload("weed_test.go")
	assert.Nil(t, err)
	//log.Printf("url: %s\n", url)

	// проверка на обработку ошибок (перепутал url и fid)
	err = delete("4,23456")
	assert.NotNil(t, err)

	// проверка на корректный url
	err = delete(url)
	assert.Nil(t, err)
}

func TestGetFID(t *testing.T) {
	fn, err := getFID("http://192.168.0.105:9091/6,23441c89622148.gif")
	assert.Nil(t, err)
	assert.Equal(t, fn, "6,23441c89622148.gif", "должны совпадать")
}

func TestGetFilename(t *testing.T) {
	fn, err := getFilename("http://192.168.0.105:9091/6,23441c89622148.gif")
	assert.Nil(t, err)
	assert.Equal(t, fn, "23441c89622148.gif", "должны совпадать")
}
