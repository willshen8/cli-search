package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/db"
	"github.com/willshen8/cli-search/pkg/parser"
)

func SetupDatabase() db.DB {
	database := db.DB{}
	organizationsRecord_1 := parser.DataRecord{"_id": 101, "url": "http://initech.zendesk.com/api/v2/organizations/101.json", "tags": []string{"Fulton", "West", "Rodriguez", "Farley"}}
	organizationsRecord_2 := parser.DataRecord{"_id": 102, "url": "http://initech.zendesk.com/api/v2/organizations/102.json"}
	organizationsDataRecords := []parser.DataRecord{organizationsRecord_1, organizationsRecord_2}
	organizationsRows := db.CreateRows(organizationsDataRecords)
	organizationsTable := db.CreateTable(organizationsRows)
	database["organizations"] = organizationsTable

	userRecord_1 := parser.DataRecord{"_id": 1, "organization_id": 101}
	userRecord_2 := parser.DataRecord{"_id": 2, "organization_id": 102}
	usersDataRecords := []parser.DataRecord{userRecord_1, userRecord_2}
	userRows := db.CreateRows(usersDataRecords)
	userTable := db.CreateTable(userRows)
	database["users"] = userTable

	ticketRecord_1 := parser.DataRecord{"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b", "organization_id": 101, "assignee_id": 1, "submitter_id": 11}
	ticketRecord_2 := parser.DataRecord{"_id": "1a227508-9f39-427c-8f57-1b72f3fab87c", "organization_id": 102, "assignee_id": 22, "submitter_id": 1}
	ticketDataRecords := []parser.DataRecord{ticketRecord_1, ticketRecord_2}
	ticketsRows := db.CreateRows(ticketDataRecords)
	ticketsTable := db.CreateTable(ticketsRows)
	database["tickets"] = ticketsTable
	return database
}

func TestSearchSuccess(t *testing.T) {
	database := SetupDatabase()
	actual, err := Search(database, "organizations", "_id", "101")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
}

func TestSearchInvalidFieldInTicketsTable(t *testing.T) {
	database := SetupDatabase()
	actual, err := Search(database, "organizations", "invalidField", "101")
	expectedResult := []string(nil)
	expectedErr := errors.NewError(errors.ErrInvalidSearchField, "invalidField").Error()
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedErr, err.Error())
}

func TestSearchInvalidTable(t *testing.T) {
	database := SetupDatabase()
	actual, err := Search(database, "invalidTable", "_id", "101")
	expectedResult := []string(nil)
	expectedErr := errors.NewError(errors.ErrInvalidTable, "invalidTable").Error()
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedErr, err.Error())
}

func TestSearchByTagName(t *testing.T) {
	database := SetupDatabase()
	actual, err := Search(database, "organizations", "tags", "West")
	fmt.Println(actual)
	expectedIDs := []string{"101"}
	assert.Equal(t, expectedIDs, actual)
	assert.Equal(t, nil, err)
}

func TestSearchWithoutSpecifiedValue(t *testing.T) {
	database := SetupDatabase()
	actual, err := Search(database, "organizations", "_id", "")
	var expectedResultIDs = []string(nil)
	assert.Equal(t, expectedResultIDs, actual)
	assert.Equal(t, nil, err)
}

func TestSearchRelatedEntitiesByOrgID(t *testing.T) {
	database := SetupDatabase()
	actualResult := SearchRelatedEntities(database, "organizations", "101")
	expectedNumOfRelatedUsers := 1
	expectedNumOfRelatedTickets := 1
	expectedUserID := "1"
	expectedTicketID := "436bf9b0-1147-4c0a-8439-6f79833bff5b"
	fmt.Println(actualResult)
	assert.Equal(t, expectedNumOfRelatedUsers, len(actualResult["users"]))
	assert.Equal(t, expectedNumOfRelatedTickets, len(actualResult["tickets"]))
	assert.Equal(t, expectedUserID, actualResult["users"][0])
	assert.Equal(t, expectedTicketID, actualResult["tickets"][0])
}

func TestSearchRelatedEntitiesByUserID(t *testing.T) {
	database := SetupDatabase()
	actualResult := SearchRelatedEntities(database, "users", "1")
	expectedTicketIDs := []string{"436bf9b0-1147-4c0a-8439-6f79833bff5b", "1a227508-9f39-427c-8f57-1b72f3fab87c"}
	assert.Equal(t, expectedTicketIDs, actualResult["tickets"])
}
