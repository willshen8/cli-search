package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
	organisationMap, err := ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "external_id", "9270ed79-35eb-4a38-a46f-35725197ea8d")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
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
	organisationMap, err := ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "_id", "101")
	var expectedResultID = []string{"101"}
	assert.Equal(t, expectedResultID, actual)
	assert.Equal(t, nil, err)
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
	organisationMap, err := ParseJsonToMapOfMap(testData)
	assert.Equal(t, nil, err)
	actual, err := Search(organisationMap, "external_id", "")
	var expectedResultIDs = []string{"101", "102"}
	assert.Equal(t, expectedResultIDs, actual)
	assert.Equal(t, nil, err)
}
