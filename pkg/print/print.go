package print

import (
	"fmt"

	"github.com/willshen8/cli-search/pkg/db"
	"github.com/willshen8/cli-search/pkg/search"
)

// PrintResults prints out the search results and its related entities
func PrintResults(database db.DB, table string, ids []string) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Search Results: Total number of records found = ", len(ids))
	for index, val := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", index+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")

		for k, v := range database[table][val] {
			fmt.Printf("%-20v %-10v\n", k, v)
		}
		PrintRelatedEntities(database, table, val)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------ End of Search Result ------------------------")
	fmt.Println("----------------------------------------------------------------------")
}

// PrintRelatedEntities takes a lists of ids from search results and returns the related data
// from the other tables, it will only print out the IDs from the other tables
func PrintRelatedEntities(database db.DB, table string, id string) {
	switch table {
	case "organizations":
		relatedEntities := search.SearchRelatedEntities(database, table, id)
		PrintEntity("users", relatedEntities)
		PrintEntity("tickets", relatedEntities)
	case "users":
		relatedEntities := search.SearchRelatedEntities(database, table, id)
		PrintEntity("tickets", relatedEntities)
	}
}

// PrintEntity prints all the related data entities in other tables given by m which stores all the
// ids of related entity
func PrintEntity(table string, m map[string][]string) {
	fmt.Println("-------------------------- Related ", table, "--------------------------")
	for i, v := range m[table] {
		fmt.Printf("Result %-1v: _id = %-20v\n", i+1, v)
	}
}

// PrintAllAvailableFields prints all the available fields in a table
func PrintAllAvailableFields(table string) {
	fmt.Println("---------------------- Available fields in ", table, "----------------------")
	switch table {
	case "organizations":
		for i, v := range search.OrgFields {
			fmt.Printf("%v: %-0v\n", i+1, v)
		}
	case "users":
		for i, v := range search.UserFields {
			fmt.Printf("%v: %-0v\n", i+1, v)
		}
	case "tickets":
		for i, v := range search.TicketFields {
			fmt.Printf("%v: %-0v\n", i+1, v)
		}
	}
	fmt.Println("---------------------------- End of the list ----------------------------")
}
