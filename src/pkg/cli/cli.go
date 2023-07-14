package cli

import (
	"flag"
)

/*
Cli is an struct to handle and put on a only place the params data
field:
- Search: Search keyword
- List: List day entries
*/
type Cli struct {
	Search string
	List   bool
}

/*
HandleArgs put on a one place all the info passed through the params
Params:
- args: []String with params (os.Args)
Return:
- Cli: Cli struct with the params loadedd and they default values
*/
func HandleArgs(args []string) Cli {
	cli := Cli{}

	flag.StringVar(&cli.Search, "s", "", "Search keyword")
	flag.BoolVar(&cli.List, "l", false, "List day entries")
	flag.Parse()

	return cli
}
