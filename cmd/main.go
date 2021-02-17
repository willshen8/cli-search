package main

import (
	"os"

	"github.com/willshen8/cli-search/pkg/file"
	"github.com/willshen8/cli-search/pkg/print"
	"github.com/willshen8/cli-search/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args               = os.Args[1:]
	app                = kingpin.New("CLI-Search", "Welcome to CLI Search!")
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
	dataBase := make(map[string]map[string]map[string]interface{}, 3)
	switch kingpin.MustParse(app.Parse(args)) {
	// Process config command
	case config.FullCommand():
		err := file.CopyFile(*configOrgJsonFile, defaultOrgFile)
		search.HandleError(err)
		err = file.CopyFile(*configUserJsonFile, defaultUsersFile)
		search.HandleError(err)
		err = file.CopyFile(*configTicketJsonFile, defaultTicketsFile)
		search.HandleError(err)

	case list.FullCommand():
		print.PrintAllAvailableFields(*listTable)

	// Process search command
	case query.FullCommand():
		dataBase, err := file.ParseFileAndStoreInDb(search.ORGANISATION, defaultOrgFile, dataBase)
		search.HandleError(err)
		dataBase, err = file.ParseFileAndStoreInDb(search.USER, defaultUsersFile, dataBase)
		search.HandleError(err)
		dataBase, err = file.ParseFileAndStoreInDb(search.TICKET, defaultTicketsFile, dataBase)
		search.HandleError(err)
		searchResults, err := search.Search(dataBase[*queryTable], *queryTable, *queryField, *queryValue)
		search.HandleError(err)
		print.PrintResults(*queryTable, searchResults, dataBase)
	}
}
