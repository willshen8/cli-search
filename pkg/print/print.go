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
	for index, id := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", index+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")

		for k, v := range database[table][id] {
			fmt.Printf("%-20v %-10v\n", k, v)
		}
		PrintRelatedEntities(database, table, id)
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
		PrintEntity("organizations", relatedEntities)
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
func PrintAllAvailableFields(database db.DB, table string) {
	var availableFields []string
	var randomRow db.Row
	for _, v := range database[table] {
		randomRow = v
		break
	}
	for key := range randomRow { // build a map to check whether a field exists
		availableFields = append(availableFields, key)
	}

	fmt.Println("---------------------- Available fields in", table, "----------------------")
	for i, v := range availableFields {
		fmt.Printf("%v: %-0v\n", i+1, v)
	}

	fmt.Println("---------------------------- End of the list ----------------------------")
}
