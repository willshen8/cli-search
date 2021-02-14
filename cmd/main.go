package main

import (
	"os"

	"github.com/willshen8/zendesk-coding-challenge/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args                     = os.Args[1:]
	app                      = kingpin.New("Zendesk-Search", "Welcome to Zendesk Search!")
	defaultOrganisationsFile = "config/organizations.json"
	defaultUsersFile         = "config/users.json"
	defaultTicketsFile       = "config/tickets.json"

	// config is a command that user can execute followed by the name of 3 files
	config             = app.Command("config", "Config the data source files by specifying the files you want to use.")
	configOrganisation = config.Arg("organisation", "Organisation json file.").Required().String()
	configUser         = config.Arg("user", "User json file.").Required().String()
	configTicket       = config.Arg("ticket", "Ticket json file.").Required().String()

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

	// Process search command
	case query.FullCommand():

		orgFile, err := os.Open(defaultOrganisationsFile)
		search.HandleError(err)
		orgMap, err := search.ParseJsonToMapOfMap(orgFile)
		search.HandleError(err)
		dataBase["organisation"] = orgMap

		userFile, err := os.Open(defaultUsersFile)
		search.HandleError(err)
		usersMap, err := search.ParseJsonToMapOfMap(userFile)
		search.HandleError(err)
		dataBase["user"] = usersMap

		ticketFile, err := os.Open(defaultTicketsFile)
		search.HandleError(err)
		ticketsMap, err := search.ParseJsonToMapOfMap(ticketFile)
		search.HandleError(err)
		dataBase["ticket"] = ticketsMap

		searchResults, err := search.Search(dataBase[*queryTable], *queryTable, *queryField, *queryValue)
		search.HandleError(err)
		search.PrintResults(*queryTable, searchResults, orgMap, usersMap, ticketsMap)
	}
}
