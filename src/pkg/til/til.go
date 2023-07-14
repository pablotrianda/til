package til

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

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

	db.SaveTil(title, fullNote)
}

/*
createNoteTempFile create a new temp file in /tmp directory
*/
func createNoteTempFile() string {
	fileName := "/tmp/_til.md"
	err := ioutil.WriteFile(fileName, []byte("#  "), 0644)
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
