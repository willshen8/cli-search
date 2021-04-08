package search

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/db"
)

// search takes any table, a field name and a value and returns a slice of IDs of the records found
func Search(database db.DB, table string, field string, value string) ([]string, error) {
	if _, found := database[table]; !found {
		return nil, errors.NewError(errors.ErrInvalidTable, table)
	}

	availableFields := make(map[string]bool)
	var randomRow db.Row
	for _, v := range database[table] {
		randomRow = v
		break
	}
	for key := range randomRow { // build a map to check whether a field exists
		availableFields[key] = true
	}
	if _, found := availableFields[field]; !found {
		return nil, errors.NewError(errors.ErrInvalidSearchField, field)
	}

	var result []string
	for k, v := range database[table] {
		if v[field] != nil && reflect.TypeOf(v[field]).Kind() == reflect.Slice { // check for field that is a slice
			s := reflect.ValueOf(v[field])
			for i := 0; i < s.Len(); i++ {
				actualValue := s.Index(i).Interface().(string)
				if actualValue == value {
					fmt.Println("match = ", actualValue)
					result = append(result, string(k))
				}
			}
		} else if fmt.Sprintf("%v", v[field]) == value {
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
	case "organizations":
		for _, foreignKey := range db.OrganizationsEntity.ForeignKeys {
			foundUsers, err := Search(database, "users", foreignKey, id) // search user table first
			userIds = append(userIds, foundUsers...)
			errors.HandleError(err)
			foundTickets, err := Search(database, "tickets", foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	case "users":
		for _, foreignKey := range db.UsersEntity.ForeignKeys {
			foundTickets, err := Search(database, "tickets", foreignKey, id) // then search ticket table
			ticketIds = append(ticketIds, foundTickets...)
			errors.HandleError(err)
		}
	}
	// store the 3 sets of results into the result map
	result["users"] = userIds
	result["tickets"] = ticketIds
	return result
}
