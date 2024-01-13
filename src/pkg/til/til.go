package til

import (
	"log"
	"os"
	"os/exec"
	"strings"

	constans "github.com/pablotrianda/til/src/constants"
	"github.com/pablotrianda/til/src/pkg/db"
)

/*
NewTil create a new til and save in the database
*/
func NewTil() {
	file := createNoteTempFile()
	runEditor(file)
	fullNote := readNoteFromTempFile(file)
	title := getTitleFromNote(fullNote)

	if fileIsEmpty(title) {
		return
	}

	db.SaveTil(title, fullNote)
}

/*
Check if the file title is equal to the TIL_DEFAULT_CONTENT
*/
func fileIsEmpty(file string) bool {
	return strings.Split(file, "\n")[0] == constans.TIL_DEFAULT_CONTENT
}

/*
createNoteTempFile create a new temp file in /tmp directory
*/
func createNoteTempFile() string {
	fileName := "/tmp/_til.md"
	err := os.WriteFile(fileName, []byte(constans.TIL_DEFAULT_CONTENT), 0644)
	check(err)

	return fileName
}

func getTitleFromNote(fullNote string) string {
	return strings.Split(fullNote, "\n")[0]
}

// Run the editor with the name of the new file
func runEditor(file string) {
	editor := os.Getenv("EDITOR")
	var cmd *exec.Cmd
	if editor == "nvim" {
		nvimParams := "+call cursor(1, 3)"
		cmd = exec.Command("nvim", nvimParams, file)
	} else {
		cmd = exec.Command(editor, file)
	}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
	cmd.Wait()
}

func readNoteFromTempFile(file string) string {
	dat, err := os.ReadFile(file)
	check(err)

	return string(dat)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
