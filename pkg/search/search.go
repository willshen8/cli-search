package search

import (
	"fmt"
	"sort"

	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/db"
)

var (
	ORGANIZATIONS = "organizations"
	USERS         = "users"
	TICKETS       = "tickets"
)

// search takes any table, a field name and a value and returns a slice of IDs of the records found
func Search(database db.DB, table string, field string, value string) ([]string, error) {
	if _, found := database[table]; !found {
		return nil, errors.NewError(errors.ErrInvalidTable, table)
	}

	var result []string
	for k, v := range database[table] {
		if fmt.Sprintf("%v", v[field]) == value {
			result = append(result, string(k))
		}
	}
	sort.Strings(result)
	return result, nil
}

// SearchRelatedEntities takes the id of the table and search for related entities in the other two tables
// and then store the results in a map of slice
func SearchRelatedEntities(database db.DB, table string, id string) map[string][]string {
	var result = make(map[string][]string)
	var userIds, ticketIds []string
	switch table {
	case ORGANIZATIONS:
		for _, foreignKey := range db.OrganizationsEnity.ForeignKeys {
			foundUsers, err := Search(database, USERS, foreignKey, id) // search user table first
			userIds = append(userIds, foundUsers...)
			errors.HandleError(err)
			foundTickets, err := Search(database, TICKETS, foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	case USERS:
		for _, foreignKey := range db.UsersEntity.ForeignKeys {
			foundTickets, err := Search(database, TICKETS, foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	}
	// store the 3 sets of results into the result map
	result[USERS] = userIds
	result[TICKETS] = ticketIds
	return result
}
