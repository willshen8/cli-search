package main

import (
	"os"

	"github.com/willshen8/zendesk-coding-challenge/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args               = os.Args[1:]
	app                = kingpin.New("Zendesk-Search", "Welcome to Zendesk Search!")
	defaultOrgFile     = "config/organizations.json"
	defaultUsersFile   = "config/users.json"
	defaultTicketsFile = "config/tickets.json"

	// config is a command that user can execute followed by the name of 3 files
	config               = app.Command("config", "Config the data source files by specifying the files you want to use.")
	configOrgJsonFile    = config.Arg("organisation json file", "Organisation json file.").Required().String()
	configUserJsonFile   = config.Arg("users json file", "User json file.").Required().String()
	configTicketJsonFile = config.Arg("tickets json file", "Ticket json file.").Required().String()

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
		err := search.CopyFile(*configOrgJsonFile, defaultOrgFile)
		search.HandleError(err)
		err = search.CopyFile(*configUserJsonFile, defaultUsersFile)
		search.HandleError(err)
		err = search.CopyFile(*configTicketJsonFile, defaultTicketsFile)
		search.HandleError(err)

	// Process search command
	case query.FullCommand():
		dataBase, err := search.ParseFileAndStoreInDb(search.ORGANISATION, defaultOrgFile, dataBase)
		search.HandleError(err)
		dataBase, err = search.ParseFileAndStoreInDb(search.USER, defaultUsersFile, dataBase)
		search.HandleError(err)
		dataBase, err = search.ParseFileAndStoreInDb(search.TICKET, defaultTicketsFile, dataBase)
		search.HandleError(err)
		searchResults, err := search.Search(dataBase[*queryTable], *queryTable, *queryField, *queryValue)
		search.HandleError(err)
		search.PrintResults(*queryTable, searchResults, dataBase)
	}
}
