package main

import (
	"fmt"
	"os"

	"github.com/pablotrianda/til/src/pkg/cli"
	"github.com/pablotrianda/til/src/pkg/db"
	"github.com/pablotrianda/til/src/pkg/til"
	"github.com/pablotrianda/til/src/pkg/tui"
)

/*
TIL app, cli app to save notes.
Allowed params:

	-s search by hastag

TIL structure example:
```
# Title
body notes
body notes 2
body notes 3

#hashtag1 ##hashtag2 hashtag3
```
*/
func main() {
	params := cli.HandleArgs(os.Args)
	db.Init()

	if params.Search != "" {
		tui.List(params.Search, db.Search(params.Search))
		return
	}

	if params.List {
		fmt.Println("Listar todo!!") // TODO
		return
	}

	til.NewTil()
}
