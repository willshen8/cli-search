package search

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintResultsUsingOrgTable(t *testing.T) {
	dB := SetupTest()
	id := []string{"101"}

	actual := captureOutput(func() {
		PrintResults(ORGANISATION, id, dB)
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInOrgTable(t *testing.T) {
	actual := captureOutput(func() {
		PrintAllAvailableFields(ORGANISATION)
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInUserTable(t *testing.T) {
	actual := captureOutput(func() {
		PrintAllAvailableFields(USER)
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInTicketTable(t *testing.T) {
	actual := captureOutput(func() {
		PrintAllAvailableFields(TICKET)
	})
	assert.NotEmpty(t, actual)
}

func TestPrintResultsWithUserTable(t *testing.T) {
	dB := SetupTest()
	id := []string{"1"}

	actual := captureOutput(func() {
		PrintResults(USER, id, dB)
	})
	assert.NotEmpty(t, actual)
}

// captureOutput is a test utility function that captures the output to standardout
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func SetupTest() map[string]map[string]map[string]interface{} {
	var orgData = strings.NewReader(`[
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
	organisationMap, _ := ParseJsonToMapOfMap(orgData)
	userMap, _ := ParseJsonToMapOfMap(usersData)
	ticketsMap, _ := ParseJsonToMapOfMap(ticketData)
	dataBase[ORGANISATION] = organisationMap
	dataBase[USER] = userMap
	dataBase[TICKET] = ticketsMap
	return dataBase
}
