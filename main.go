package main

import (
	"fmt"
	"os"

	"github.com/willshen8/zendesk-coding-challenge/pkg/search"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args                     = os.Args[1:]
	app                      = kingpin.New("Zendesk-Search", "Welcome to Zendesk Search!")
	defaultOrganisationsFile = "./config/organizations.json"
	defaultUserFile          = "./config/users.json"
	defaultTicketsFile       = "./config/tickets.json"

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
		orgFile, _ := os.Open(defaultOrganisationsFile)
		organisations, err := search.ParseJsonOrgs(orgFile)
		search.HandleError(err)

		userFile, _ := os.Open(defaultUserFile)
		users, err := search.ParseJsonUsers(userFile)
		search.HandleError(err)

		ticketFile, _ := os.Open(defaultTicketsFile)
		tickets, err := search.ParseJsonTickets(ticketFile)
		search.HandleError(err)

		fmt.Println("organisations", len(organisations))
		fmt.Println("users", len(users))
		fmt.Println("tickets", len(tickets))
	}
}

// Configs contains the arguments parsed from command line
type Config struct {
	organization_file *os.File
	user_file         *os.File
	ticket_file       *os.File
}

// Search is a structure that contains the specific search item parsed from command line
// type Search struct {
// 	table string
// 	field string
// 	value string
// }
