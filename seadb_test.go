package pic

import (
	//"log"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetSmallGifs(t *testing.T) {
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
