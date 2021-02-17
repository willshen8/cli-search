package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willshen8/cli-search/pkg/file"
)

func TestSearchSuccess(t *testing.T) {
	var testData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	},
	{
	  "_id": 102,
	  "url": "http://initech.zendesk.com/api/v2/organizations/102.json",
	  "external_id": "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
	  "name": "Nutralab",
	  "domain_names": [
		"trollery.com",
		"datagen.com",
		"bluegrain.com",
		"dadabase.com"
	  ],
	  "created_at": "2016-04-07T08:21:44 -10:00",
	  "details": "Non profit",
	  "shared_tickets": false,
	  "tags": [
		"Cherry",
		"Collier",
		"Fuentes",
		"Trevino"
	  ]
	}
  ]`)
	organisationMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, ORGANISATION, "external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
}

func TestSearchTicketTableSuccess(t *testing.T) {
	var testTicketData = strings.NewReader(`[
		{
			"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			"url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			"external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
			"created_at": "2016-04-28T11:19:34 -10:00",
			"type": "incident",
			"subject": "A Catastrophe in Korea (North)",
			"description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
			"priority": "high",
			"status": "pending",
			"submitter_id": 38,
			"assignee_id": 24,
			"organization_id": 116,
			"tags": [
			  "Ohio",
			  "Pennsylvania",
			  "American Samoa",
			  "Northern Mariana Islands"
			],
			"has_incidents": false,
			"due_at": "2016-07-31T02:37:50 -10:00",
			"via": "web"
		}
	  ]`)
	ticketsMap, err := file.ParseJsonToMapOfMap(testTicketData)
	assert.Equal(t, nil, err)
	actual, err := Search(ticketsMap, "ticket", "submitter_id", "38")
	expectedResult := []string{"436bf9b0-1147-4c0a-8439-6f79833bff5b"}
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, nil, err)
}

func TestSearchInvalidFieldInTicketsTable(t *testing.T) {
	var testTicketData = strings.NewReader(`[
		{
			"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			"url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			"external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
			"created_at": "2016-04-28T11:19:34 -10:00",
			"type": "incident",
			"subject": "A Catastrophe in Korea (North)",
			"description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
			"priority": "high",
			"status": "pending",
			"submitter_id": 38,
			"assignee_id": 24,
			"organization_id": 116,
			"tags": [
			  "Ohio",
			  "Pennsylvania",
			  "American Samoa",
			  "Northern Mariana Islands"
			],
			"has_incidents": false,
			"due_at": "2016-07-31T02:37:50 -10:00",
			"via": "web"
		}
	  ]`)
	ticketsMap, err := file.ParseJsonToMapOfMap(testTicketData)
	assert.Equal(t, nil, err)
	actual, err := Search(ticketsMap, "ticket", "invalidField", "101")
	expectedResult := []string(nil)
	expectedErr := ErrInvalidSearchField
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedErr, err)
}

func TestSearchInvalidFieldInOrgTable(t *testing.T) {
	var testData = strings.NewReader(`[
		{
			"_id": 1,
			"url": "http://initech.zendesk.com/api/v2/users/1.json",
			"external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
			"name": "Francisca Rasmussen",
			"alias": "Miss Coffey",
			"created_at": "2016-04-15T05:19:46 -10:00",
			"active": true,
			"verified": true,
			"shared": false,
			"locale": "en-AU",
			"timezone": "Sri Lanka",
			"last_login_at": "2013-08-04T01:03:27 -10:00",
			"email": "coffeyrasmussen@flotonic.com",
			"phone": "8335-422-718",
			"signature": "Don't Worry Be Happy!",
			"organization_id": 119,
			"tags": [
			  "Springville",
			  "Sutton",
			  "Hartsville/Hartley",
			  "Diaperville"
			],
			"suspended": true,
			"role": "admin"
		}
  ]`)
	userMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(userMap, "user", "invalidField", "")
	expectedResult := []string(nil)
	expectedErr := ErrInvalidSearchField
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedErr, err)
}

func TestSearchInvalidFieldInTicketTable(t *testing.T) {
	var testData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	}
  ]`)
	organisationMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, ORGANISATION, "invalidField", "")
	expectedResult := []string(nil)
	expectedErr := ErrInvalidSearchField
	assert.Equal(t, expectedResult, actual)
	assert.Equal(t, expectedErr, err)
}
func TestSearchByID(t *testing.T) {
	var testData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	}
  ]`)
	organisationMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, ORGANISATION, "_id", "101")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
}

func TestSearchByEachField(t *testing.T) {
	var testData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	}
  ]`)
	organisationMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	for key := range OrgMap {
		actual, err := Search(organisationMap, ORGANISATION, key, "")
		assert.NotNil(t, actual)
		assert.Equal(t, nil, err)

	}
}

func TestSearchWithoutSpecifiedValue(t *testing.T) {
	var testData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	},
	{
	  "_id": 102,
	  "url": "http://initech.zendesk.com/api/v2/organizations/102.json",
	  "external_id": "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
	  "name": "Nutralab",
	  "domain_names": [
		"trollery.com",
		"datagen.com",
		"bluegrain.com",
		"dadabase.com"
	  ],
	  "created_at": "2016-04-07T08:21:44 -10:00",
	  "details": "Non profit",
	  "shared_tickets": false,
	  "tags": [
		"Cherry",
		"Collier",
		"Fuentes",
		"Trevino"
	  ]
	}
  ]`)
	organisationMap, err := file.ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, ORGANISATION, "external_id", "")
	var expectedResultIDs = []string{"101", "102"}
	assert.Equal(t, expectedResultIDs, actual)
	assert.Equal(t, nil, err)
}

func TestSearchUserTable(t *testing.T) {
	var usersData = strings.NewReader(`[
		{
			"_id": 1,
			"url": "http://initech.zendesk.com/api/v2/users/1.json",
			"external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
			"name": "Francisca Rasmussen",
			"alias": "Miss Coffey",
			"created_at": "2016-04-15T05:19:46 -10:00",
			"active": true,
			"verified": true,
			"shared": false,
			"locale": "en-AU",
			"timezone": "Sri Lanka",
			"last_login_at": "2013-08-04T01:03:27 -10:00",
			"email": "coffeyrasmussen@flotonic.com",
			"phone": "8335-422-718",
			"signature": "Don't Worry Be Happy!",
			"organization_id": 119,
			"tags": [
			  "Springville",
			  "Sutton",
			  "Hartsville/Hartley",
			  "Diaperville"
			],
			"suspended": true,
			"role": "admin"
		}
  ]`)
	usersMap, err := file.ParseJsonToMapOfMap(usersData)
	assert.Equal(t, nil, err)
	actual, err := Search(usersMap, "user", "active", "true")
	expectedResults := []string{"1"}
	assert.Equal(t, expectedResults, actual)
	assert.Equal(t, nil, err)

	searchIdResult, err := Search(usersMap, "user", "_id", "1")
	assert.Equal(t, nil, err)
	expectedIdResult := []string{"1"}
	assert.Equal(t, expectedIdResult, searchIdResult)
}

func TestSearchRelatedEntitiesByOrgID(t *testing.T) {
	var usersData = strings.NewReader(`[
		{
			"_id": 24,
			"url": "http://initech.zendesk.com/api/v2/users/1.json",
			"external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
			"name": "Francisca Rasmussen",
			"alias": "Miss Coffey",
			"created_at": "2016-04-15T05:19:46 -10:00",
			"active": true,
			"verified": true,
			"shared": false,
			"locale": "en-AU",
			"timezone": "Sri Lanka",
			"last_login_at": "2013-08-04T01:03:27 -10:00",
			"email": "coffeyrasmussen@flotonic.com",
			"phone": "8335-422-718",
			"signature": "Don't Worry Be Happy!",
			"organization_id": 101,
			"tags": [
			  "Springville",
			  "Sutton",
			  "Hartsville/Hartley",
			  "Diaperville"
			],
			"suspended": true,
			"role": "admin"
		}
  ]`)
	var OrgData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"
	  ]
	}
  ]`)
	var ticketData = strings.NewReader(`[
		{
			"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			"url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			"external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
			"created_at": "2016-04-28T11:19:34 -10:00",
			"type": "incident",
			"subject": "A Catastrophe in Korea (North)",
			"description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
			"priority": "high",
			"status": "pending",
			"submitter_id": 38,
			"assignee_id": 24,
			"organization_id": 101,
			"tags": [
			  "Ohio",
			  "Pennsylvania",
			  "American Samoa",
			  "Northern Mariana Islands"
			],
			"has_incidents": false,
			"due_at": "2016-07-31T02:37:50 -10:00",
			"via": "web"
		  }
  	]`)
	dataBase := make(map[string]map[string]map[string]interface{}, 3)
	organisationMap, _ := file.ParseJsonToMapOfMap(OrgData)
	userMap, _ := file.ParseJsonToMapOfMap(usersData)
	ticketMap, _ := file.ParseJsonToMapOfMap(ticketData)
	dataBase[ORGANISATION] = organisationMap
	dataBase[USER] = userMap
	dataBase[TICKET] = ticketMap

	actualResult := SearchRelatedEntities(ORGANISATION, "101", dataBase)
	expectedNumOfRelatedUsers := 1
	expectedNumOfRelatedTickets := 1
	expectedUserID := "24"
	expectedTicketID := "436bf9b0-1147-4c0a-8439-6f79833bff5b"
	assert.Equal(t, expectedNumOfRelatedUsers, len(actualResult[USER]))
	assert.Equal(t, expectedNumOfRelatedTickets, len(actualResult[TICKET]))
	assert.Equal(t, expectedUserID, actualResult[USER][0])
	assert.Equal(t, expectedTicketID, actualResult[TICKET][0])
}

func TestSearchRelatedEntitiesByUserID(t *testing.T) {
	var usersData = strings.NewReader(`[
		{
			"_id": 24,
			"url": "http://initech.zendesk.com/api/v2/users/1.json",
			"external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
			"name": "Francisca Rasmussen",
			"alias": "Miss Coffey",
			"created_at": "2016-04-15T05:19:46 -10:00",
			"active": true,
			"verified": true,
			"shared": false,
			"locale": "en-AU",
			"timezone": "Sri Lanka",
			"last_login_at": "2013-08-04T01:03:27 -10:00",
			"email": "coffeyrasmussen@flotonic.com",
			"phone": "8335-422-718",
			"signature": "Don't Worry Be Happy!",
			"organization_id": 101,
			"tags": [
			  "Springville",
			  "Sutton",
			  "Hartsville/Hartley",
			  "Diaperville"
			],
			"suspended": true,
			"role": "admin"
		}
  	]`)
	var OrgData = strings.NewReader(`[
	{
	  "_id": 101,
	  "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
	  "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
	  "name": "Enthaze",
	  "domain_names": [
		"kage.com",
		"ecratic.com",
		"endipin.com",
		"zentix.com"
	  ],
	  "created_at": "2016-05-21T11:10:28 -10:00",
	  "details": "MegaCorp",
	  "shared_tickets": false,
	  "tags": [
		"Fulton",
		"West",
		"Rodriguez",
		"Farley"]
		}
  	]`)
	var ticketData = strings.NewReader(`[
		{
			"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
			"url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
			"external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
			"created_at": "2016-04-28T11:19:34 -10:00",
			"type": "incident",
			"subject": "A Catastrophe in Korea (North)",
			"description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
			"priority": "high",
			"status": "pending",
			"submitter_id": 38,
			"assignee_id": 24,
			"organization_id": 101,
			"tags": [
			  "Ohio",
			  "Pennsylvania",
			  "American Samoa",
			  "Northern Mariana Islands"
			],
			"has_incidents": false,
			"due_at": "2016-07-31T02:37:50 -10:00",
			"via": "web"
		  },
		  {
			"_id": "1a227508-9f39-427c-8f57-1b72f3fab87c",
			"url": "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
			"external_id": "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
			"created_at": "2016-04-14T08:32:31 -10:00",
			"type": "incident",
			"subject": "A Catastrophe in Micronesia",
			"description": "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
			"priority": "low",
			"status": "hold",
			"submitter_id": 24,
			"assignee_id": 38,
			"organization_id": 112,
			"tags": [
			  "Puerto Rico",
			  "Idaho",
			  "Oklahoma",
			  "Louisiana"
			],
			"has_incidents": false,
			"due_at": "2016-08-15T05:37:32 -10:00",
			"via": "chat"
		  }
  	]`)

	dataBase := make(map[string]map[string]map[string]interface{}, 3)
	organisationMap, _ := file.ParseJsonToMapOfMap(OrgData)
	userMap, _ := file.ParseJsonToMapOfMap(usersData)
	ticketMap, _ := file.ParseJsonToMapOfMap(ticketData)
	dataBase[ORGANISATION] = organisationMap
	dataBase[USER] = userMap
	dataBase[TICKET] = ticketMap
	actualResult := SearchRelatedEntities(USER, "24", dataBase)
	expectedTicketIDs := []string{"436bf9b0-1147-4c0a-8439-6f79833bff5b", "1a227508-9f39-427c-8f57-1b72f3fab87c"}
	assert.Equal(t, expectedTicketIDs, actualResult["ticket"])
}
