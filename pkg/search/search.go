package search

import (
	"fmt"
	"sort"

	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/db"
)

var (
	ORGANISATION = "organisation"
	USER         = "user"
	TICKET       = "ticket"
)

// search takes any table, a field name and a value and returns a slice of IDs of the records found
func Search(m map[string]map[string]interface{}, table string, field string, value string) ([]string, error) {
	switch table {
	case ORGANISATION:
		if _, found := OrgMap[field]; !found {
			return nil, errors.NewError(errors.ErrInvalidSearchField, field)
		}
	case USER:
		if _, found := UserMap[field]; !found {
			return nil, errors.NewError(errors.ErrInvalidSearchField, field)
		}
	case TICKET:
		if _, found := TicketMap[field]; !found {
			return nil, errors.NewError(errors.ErrInvalidSearchField, field)
		}
	default:
		return nil, errors.NewError(errors.ErrInvalidTable, table)
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
	case ORGANISATION:
		for _, foreignKey := range db.OrganisationEnity.ForeignKeys {
			foundUsers, err := Search(dataBase[USER], USER, foreignKey, id) // search user table first
			userIds = append(userIds, foundUsers...)
			errors.HandleError(err)
			foundTickets, err := Search(dataBase[TICKET], TICKET, foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	case USER:
		for _, foreignKey := range db.UserEntity.ForeignKeys {
			foundTickets, err := Search(dataBase[TICKET], TICKET, foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	}
	// store the 3 sets of results into the result map
	result[USER] = userIds
	result[TICKET] = ticketIds
	return result
}
