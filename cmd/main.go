package main

import (
	"os"

	"github.com/willshen8/zendesk-coding-challenge/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args                     = os.Args[1:]
	app                      = kingpin.New("Zendesk-Search", "Welcome to Zendesk Search!")
	defaultOrganisationsFile = "../config/organizations.json"
	defaultUsersFile         = "../config/users.json"
	defaultTicketsFile       = "../config/tickets.json"

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
	switch kingpin.MustParse(app.Parse(args)) {
	// Process config command
	case config.FullCommand():

	// Process search command
	case query.FullCommand():
		switch *queryTable {
		case "organisation":
			orgFile, _ := os.Open(defaultOrganisationsFile)
			orgMap, err := search.ParseJsonToMapOfMap(orgFile)
			search.HandleError(err)
			searchResults, err := search.Search(orgMap, *queryField, *queryValue)
			search.HandleError(err)
			search.PrintResults(orgMap, searchResults)
		case "ticket":
			userFile, _ := os.Open(defaultTicketsFile)
			usersMap, err := search.ParseJsonToMapOfMap(userFile)
			search.HandleError(err)
			searchResults, err := search.Search(usersMap, *queryField, *queryValue)
			search.HandleError(err)
			search.PrintResults(usersMap, searchResults)
		case "user":
			ticketFile, _ := os.Open(defaultUsersFile)
			ticketsMap, err := search.ParseJsonToMapOfMap(ticketFile)
			search.HandleError(err)
			searchResults, err := search.Search(ticketsMap, *queryField, *queryValue)
			search.HandleError(err)
			search.PrintResults(ticketsMap, searchResults)
		}
	}
}