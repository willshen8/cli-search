package search

import (
	"fmt"
	"sort"
)

// search takes any table, a field name and a value and returns a slice of IDs of the records found
func Search(m map[string]map[string]interface{}, table string, field string, value string) ([]string, error) {
	switch table {
	case "organisation":
		if _, found := OrgMap[field]; !found {
			return nil, ErrInvalidSearchField
		}
	case "user":
		if _, found := UserMap[field]; !found {
			return nil, ErrInvalidSearchField
		}
	case "ticket":
		if _, found := TicketMap[field]; !found {
			return nil, ErrInvalidSearchField
		}
	default:
		return nil, ErrInvalidTable
	}

	var result []string
	if value == "" {
		for k := range m {
			result = append(result, string(k))
		}
		sort.Strings(result)
		return result, nil
	}

	for k, v := range m {
		if fmt.Sprintf("%v", v[field]) == value {
			result = append(result, string(k))
		}
	}
	sort.Strings(result)
	return result, nil
}

// SearchRelatedEntities takes the id of the table and search for related entities in the other two tables
// and then store the results in a map of slice
func SearchRelatedEntities(table string, id string, dataBase map[string]map[string]map[string]interface{}) map[string][]string {
	var result = make(map[string][]string)
	var userIds, ticketIds []string
	switch table {
	case "organisation":
		for _, foreignKey := range organisationEnity.foreignKeys {
			foundUsers, err := Search(dataBase["user"], "user", foreignKey, id) // search user table first
			userIds = append(userIds, foundUsers...)
			HandleError(err)
			foundTickets, err := Search(dataBase["ticket"], "ticket", foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			HandleError(err)
		}
	case "user":
		for _, foreignKey := range userEntity.foreignKeys {
			foundTickets, err := Search(dataBase["ticket"], "ticket", foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			HandleError(err)
		}
	}
	// store the 3 sets of results into the result map
	result["user"] = userIds
	result["ticket"] = ticketIds
	return result
}
