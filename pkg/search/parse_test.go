package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJsonToMapOfMap(t *testing.T) {
	var parseTestData = strings.NewReader(`[
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
	actual, err := ParseJsonToMapOfMap(parseTestData)
	assert.Equal(t, nil, err)
	expectedNumberOfRecords := 2
	assert.Equal(t, expectedNumberOfRecords, len(actual))
	assert.Equal(t, "MegaCorp", actual["101"]["details"])
	assert.Equal(t, false, actual["102"]["shared_tickets"])
	assert.Equal(t, float64(102), actual["102"]["_id"])
	assert.Equal(t, "http://initech.zendesk.com/api/v2/organizations/102.json", actual["102"]["url"])
}

func TestUnmarshallErrorData(t *testing.T) {
	var errorData = strings.NewReader(`[
		{
			"_id": "101"
		  }
	]`)
	_, err := ParseJsonToMapOfMap(errorData)
	assert.Equal(t, nil, err)
}
