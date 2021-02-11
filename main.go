package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("Zendesk-Search", "Welcome to Zendesk Search!")

	list      = app.Command("list", "List all possible fields in a table.")
	listTable = list.Arg("table", "Name of the table.").Required().String()
	listField = list.Arg("field", "Name of the field.").Required().String()

	search      = app.Command("search", "Search a field in a table.")
	searchTable = search.Arg("table", "Name of the table.").Required().String()
	searchField = search.Arg("field", "Name of the field.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Post message
	case list.FullCommand():
		println(*listTable)

	// Register user
	case search.FullCommand():
		println(*searchTable)
	}
}

// Configs contains the arguments parsed from command line
type Config struct {
	table string
	field string
}
