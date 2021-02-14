package search

import "fmt"

// PrintResults prints out the search results and its related entities
func PrintResults(table string, ids []string, orgMap map[string]map[string]interface{},
	userMap map[string]map[string]interface{}, ticketMap map[string]map[string]interface{}) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Search Results: Total number of records found = ", len(ids))
	for index, val := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", index+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")
		switch table {
		case "organisation":
			for k, v := range orgMap[val] {
				fmt.Printf("%-20v %-10v\n", k, v)
			}
		case "user":
			for k, v := range userMap[val] {
				fmt.Printf("%-20v %-10v\n", k, v)
			}
		case "ticket":
			for k, v := range ticketMap[val] {
				fmt.Printf("%-20v %-10v\n", k, v)
			}
		}

		PrintRelatedEntities(table, val, orgMap, userMap, ticketMap)
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------ End of Search Result ------------------------")
	fmt.Println("----------------------------------------------------------------------")
}

// PrintRelatedEntities takes a lists of ids from search results and returns the related data
// from the other tables, it will only print out the IDs from the other tables
func PrintRelatedEntities(table string, id string, orgMap map[string]map[string]interface{},
	userMap map[string]map[string]interface{}, ticketMap map[string]map[string]interface{}) {
	switch table {
	case "organisation":
		relatedEntities := SearchRelatedEntities(table, id, orgMap, userMap, ticketMap)
		PrintEntity("user", relatedEntities)
		PrintEntity("ticket", relatedEntities)
	case "user":
		relatedEntities := SearchRelatedEntities(table, id, orgMap, userMap, ticketMap)
		PrintEntity("ticket", relatedEntities)
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
