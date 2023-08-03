package pic

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Это скорее интеграционное тестированиеЮ чем модульное
// Сначала создадим тестовую БД - test-sea аналог боевой

// Выполняется один раз при прогоне
func TextMain(m *testing.M) {
	db = createTestDB()
	os.Exit(m.Run())
}

func createTestDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "itman:X753951x@(xigmanas:3306)")
	if err != nil {
		fmt.Printf("не удалось полключиться к БД: %v\n", err)
		os.Exit(0)
	}
	return db

}
func Test_getAlbum(t *testing.T) {
	type args struct {
		db   *sqlx.DB
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []Picture
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAlbum(tt.args.db, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAlbum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAlbum() = %v, want %v", got, tt.want)
			}
		})
	}
}
