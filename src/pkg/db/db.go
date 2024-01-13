package db

import (
	"fmt"
	"log"
	"os"

	cmd "github.com/pablotrianda/til/src/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Til struct {
	gorm.Model
	Title string
	Note  string
}

/*
Init initialize the local db, if not exist this will be created at
`$USER/.config/til`
*/
func Init() {
	databaseFullPath := getDatabasePath()
	if _, err := os.Stat(databaseFullPath); err != nil {
		// Create a new config dir and database
		err := os.MkdirAll(getAppFolder(), 0777)
		check(err)
		os.Create(databaseFullPath)
	}
}

/*
SaveTil Create a new entry in the database
Params:
  - title: Note title
  - note: Note body
*/
func SaveTil(title, note string) {
	db := getDatabaseConnection()
	db.AutoMigrate(&Til{})

	// Save on db
	db.Create(&Til{Title: title, Note: note})
}

/*
Search Seach in the database, on the note file(note) by hashtag
Params:
  - param: param string
*/
func Search(param string) []Til {
	var tils []Til
	db := getDatabaseConnection()
	db.Where("note LIKE ?", "%#"+param+"%").Find(&tils)

	return tils
}

/*
ListAll Return all TIL entries
*/
func ListAll() []Til {
	var tils []Til
	db := getDatabaseConnection()
	db.Order("created_at DESC").Find(&tils)

	return tils
}

/*
getDatabaseConnection Get the connection database, building with the path with the
APP_NAME and DATABASE_NAME constants.
*/
func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(getDatabasePath()), &gorm.Config{})
	check(err, "failed to connect database")
	return db
}

/*
getAppFolder Get the App folder, building with the path using the
APP_NAME constant and the userHomeDir().
*/
func getAppFolder() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("%s/.config/%s", homeDir, cmd.APP_NAME)
}

/*
getDatabasePath Get the database path, building with the path using the
DATABASE_NAME constant and the getAppFolder().
*/
func getDatabasePath() string {
	return fmt.Sprintf("%s/%s", getAppFolder(), cmd.DATABASE_NAME)
}

func check(params ...any) {
	err := params[0]
	var msg string
	if len(params) > 1 {
		msg = params[1].(string)
	}
	if err != nil {
		if msg != "" {
			log.Panic(msg)
		}
		log.Panic(err)
	}
}
