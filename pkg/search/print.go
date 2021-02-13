package search

import "fmt"

func PrintResults(table map[string]map[string]interface{}, ids []string) {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Search Results: Total number of records found = ", len(ids))
	for key, val := range ids {
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("----------------------------- Result ", key+1, "-----------------------------")
		fmt.Println("----------------------------------------------------------------------")
		for k, v := range table[val] {
			fmt.Printf("%-20v %-10v\n", k, v)
		}
	}
	fmt.Println("------------------------ End of Search Result ------------------------")
	fmt.Println("----------------------------------------------------------------------")
}
