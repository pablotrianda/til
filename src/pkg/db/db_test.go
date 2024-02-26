package db_test

import (
	"os"
	"testing"

	cmd "github.com/pablotrianda/til/src/constants"
	"github.com/pablotrianda/til/src/pkg/db"
	"gorm.io/gorm"
)

var Db *gorm.DB

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	tearDown()

	os.Exit(exitVal)
}

func setup() {
	os.Setenv("TIL_TEST", "1")
	db.Init()
	db.SaveTil("test", "example")
}

func tearDown() {
	os.Remove(cmd.DATABASE_TEST_NAME)
}

func TestFindByTitle(t *testing.T) {
	tils := db.FindByTitle("test")
	if len(tils) == 0 {
		t.Error("Result must be 0 got " + string(rune(len(tils))))
	}
}

func TestFail(t *testing.T) {
	t.Error("jajaja")
}
