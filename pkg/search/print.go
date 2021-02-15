package search

import (
	"fmt"
)

// PrintResults prints out the search results and its related entities
func PrintResults(table string, ids []string, dataBase map[string]map[string]map[string]interface{}) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Search Results: Total number of records found = ", len(ids))
	for index, val := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", index+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")

		for k, v := range dataBase[table][val] {
			fmt.Printf("%-20v %-10v\n", k, v)
		}
		PrintRelatedEntities(table, val, dataBase)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------ End of Search Result ------------------------")
	fmt.Println("----------------------------------------------------------------------")
}

// PrintRelatedEntities takes a lists of ids from search results and returns the related data
// from the other tables, it will only print out the IDs from the other tables
func PrintRelatedEntities(table string, id string, dataBase map[string]map[string]map[string]interface{}) {
	switch table {
	case ORGANISATION:
		relatedEntities := SearchRelatedEntities(table, id, dataBase)
		PrintEntity(USER, relatedEntities)
		PrintEntity(TICKET, relatedEntities)
	case USER:
		relatedEntities := SearchRelatedEntities(table, id, dataBase)
		PrintEntity(TICKET, relatedEntities)
	}
}

// PrintEntity prints all the related data entities in other tables given by m which stores all the
// ids of related entity
func PrintEntity(table string, m map[string][]string) {
	fmt.Println("-------------------------- Related ", table, "--------------------------")
	for i, v := range m[table] {
		fmt.Printf("Result %-1v ID: %-20v\n", i+1, v)
	}
}
