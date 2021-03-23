package main

import (
	"os"

	"github.com/willshen8/cli-search/internal/copy"
	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/db"
	"github.com/willshen8/cli-search/pkg/parser"
	"github.com/willshen8/cli-search/pkg/print"
	"github.com/willshen8/cli-search/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args               = os.Args[1:]
	app                = kingpin.New("CLI-Search", "Welcome to CLI Search!")
	configDir          = "config"
	defaultOrgFile     = "config/organizations.json"
	defaultUsersFile   = "config/users.json"
	defaultTicketsFile = "config/tickets.json"

	// config is a command that user can execute followed by the name of 3 files
	config               = app.Command("config", "Config the data source files by specifying the files you want to use.")
	configOrgJsonFile    = config.Arg("organisation json file", "Organisation json file.").Required().String()
	configUserJsonFile   = config.Arg("users json file", "User json file.").Required().String()
	configTicketJsonFile = config.Arg("tickets json file", "Ticket json file.").Required().String()

	list      = app.Command("list", "List all available data fields in a table.")
	listTable = list.Arg("table", "(Required) Name of the table [organsaition|user|ticket].").Required().String()

	// query is a command that user uses to query a table
	query      = app.Command("query", "Search a specific field in a table.")
	queryTable = query.Arg("table", "(Required) Name of the table.").Required().String()
	queryField = query.Arg("field", "(Required) Name of the field.").Required().String()
	queryValue = query.Arg("value", "(Optional) Value of the field searching for.").String()
)

func main() {
	// database is a three level map with each level contains table -> row -> value
	switch kingpin.MustParse(app.Parse(args)) {
	// Process config command
	case config.FullCommand():
		err := copy.CopyFile(*configOrgJsonFile, defaultOrgFile)
		errors.HandleError(err)
		err = copy.CopyFile(*configUserJsonFile, defaultUsersFile)
		errors.HandleError(err)
		err = copy.CopyFile(*configTicketJsonFile, defaultTicketsFile)
		errors.HandleError(err)

	case list.FullCommand():
		print.PrintAllAvailableFields(*listTable)

	// Process search command
	case query.FullCommand():
		database := db.DB{}
		filesNames, err := parser.GetFileNamesInDir(configDir)
		errors.HandleError(err)

		for _, file := range filesNames {
			createdTable, err := db.CreateTableFromJsonFile(configDir + "/" + file)
			errors.HandleError(err)
			fileName := parser.GetFileSuffix(file)
			database[fileName] = createdTable
		}
		searchResults, err := search.Search(database, *queryTable, *queryField, *queryValue)
		errors.HandleError(err)
		print.PrintResults(database, *queryTable, searchResults)
	}
}
