package search

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseOrganisationJsonSuccessfully(t *testing.T) {
	input := strings.NewReader(
		`[{"_id": 101, 
			"url":"http://initech.zendesk.com/api/v2/organizations/101.json",
			"external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
			"domain_names":["kage.com", "ecratic.com", "endipin.com", "zentix.com"],
			"created_at": "2016-05-21T11:10:28 -10:00",
			"details": "MegaCorp",
			"shared_tickets": false,
			"tags": ["Fulton","West","Rodriguez","Farley"]
		}]`)
	organisations, err := ParseJsonOrgs(input)
	expected := []Organisation{Organisation{ID: 101,
		Url:            "http://initech.zendesk.com/api/v2/organizations/101.json",
		External_Id:    "9270ed79-35eb-4a38-a46f-35725197ea8d",
		Domain_names:   []string{"kage.com", "ecratic.com", "endipin.com", "zentix.com"},
		Created_at:     "2016-05-21T11:10:28 -10:00",
		Details:        "MegaCorp",
		Shared_tickets: false,
		Tags:           []string{"Fulton", "West", "Rodriguez", "Farley"},
	}}
	assert.Equal(t, expected, organisations)
	assert.Equal(t, nil, err)
}
