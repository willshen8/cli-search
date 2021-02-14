package search

import "fmt"

func PrintResults(table map[string]map[string]interface{}, ids []string) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Search Results: Total number of records found = ", len(ids))
	for index, val := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", index+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")
		for k, v := range table[val] {
			fmt.Printf("%-20v %-10v\n", k, v)
		}
	}
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("------------------------ End of Search Result ------------------------")
	fmt.Println("----------------------------------------------------------------------")
}

// PrintRelatedEntities takes a lists of ids from search results and returns the related data
// from the other tables, it will only print out the IDs from the other tables
// func PrintRelatedEntities(table string, ids []string, org map[string]map[string]interface{},
// 	user map[string]map[string]interface{}, ticket map[string]map[string]interface{}) {
// 	switch table {
// 	case "organisation":

// 	}
// }
