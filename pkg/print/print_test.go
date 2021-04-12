package print

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willshen8/cli-search/pkg/db"
	"github.com/willshen8/cli-search/pkg/parser"
)

func TestPrintResultsUsingOrgTable(t *testing.T) {
	dB := CreateDB()
	ids := []string{"101"}
	actual := captureOutput(func() {
		PrintResults(dB, "dummyTable", ids)
	})
	expected := `----------------------------------------------------------------------
Search Results: Total number of records found =  1
----------------------------------------------------------------------
----------------------------- Result  1 -----------------------------
----------------------------------------------------------------------
_id                  101       
----------------------------------------------------------------------
------------------------ End of Search Result ------------------------
----------------------------------------------------------------------
`
	assert.Equal(t, expected, actual)
}

func TestPrintAllAvailableFieldsInOrgTable(t *testing.T) {
	dB := CreateDB()
	actual := captureOutput(func() {
		PrintAllAvailableFields(dB, "dummyTable")
	})
	expected := `---------------------- Available fields in dummyTable ----------------------
1: _id
---------------------------- End of the list ----------------------------
`
	assert.Equal(t, expected, actual)
}

func TestPrintAllAvailableFieldsInTicketTable(t *testing.T) {
	dB := CreateDB()
	actual := captureOutput(func() {
		PrintAllAvailableFields(dB, "dummyTable")
	})
	expected := `---------------------- Available fields in dummyTable ----------------------
1: _id
---------------------------- End of the list ----------------------------
`
	assert.Equal(t, expected, actual)
}

func TestPrintRelatedOrganizationsFields(t *testing.T) {
	dB := CreateDBForRelatedTables()
	actual := captureOutput(func() {
		PrintRelatedEntities(dB, "organizations", "101")
	})
	expected := `-------------------------- Related  users --------------------------
Result 1: _id = 1                   
-------------------------- Related  tickets --------------------------
Result 1: _id = 81bdd837-e955-4aa4-a971-ef1e3b373c6d
`
	assert.Equal(t, expected, actual)
}

func TestPrintRelatedUsersFields(t *testing.T) {
	dB := CreateDBForRelatedTables()
	actual := captureOutput(func() {
		PrintRelatedEntities(dB, "users", "1")
	})
	expected := `-------------------------- Related  tickets --------------------------
Result 1: _id = 81bdd837-e955-4aa4-a971-ef1e3b373c6d
`
	assert.Equal(t, expected, actual)
}

func TestPrintRelatedTicketsFields(t *testing.T) {
	dB := CreateDBForRelatedTables()
	actual := captureOutput(func() {
		PrintRelatedEntities(dB, "tickets", "81bdd837-e955-4aa4-a971-ef1e3b373c6d")
	})
	expected := `-------------------------- Related  users --------------------------
-------------------------- Related  organizations --------------------------
`
	assert.Equal(t, expected, actual)
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

func CreateDB() db.DB {
	database := db.DB{}
	newOrgRecord_1 := parser.DataRecord{"_id": 101}
	newOrgRecord_2 := parser.DataRecord{"_id": 102}
	dataRecords := []parser.DataRecord{newOrgRecord_1, newOrgRecord_2}
	rows := db.CreateRows(dataRecords)
	newTable := db.CreateTable(rows)
	database["dummyTable"] = newTable
	return database
}

func CreateDBForRelatedTables() db.DB {
	database := db.DB{}
	newOrgRecord_1 := parser.DataRecord{"_id": 101}
	newOrgRecord_2 := parser.DataRecord{"_id": 102}
	dataRecords := []parser.DataRecord{newOrgRecord_1, newOrgRecord_2}
	rows := db.CreateRows(dataRecords)
	newTable := db.CreateTable(rows)
	database["organizations"] = newTable

	newUserRecord_1 := parser.DataRecord{"_id": "1", "organization_id": "101"}
	userDataRecords := []parser.DataRecord{newUserRecord_1}
	usersRows := db.CreateRows(userDataRecords)
	usersTable := db.CreateTable(usersRows)
	database["users"] = usersTable

	newTicketsRecord_1 := parser.DataRecord{"_id": "81bdd837-e955-4aa4-a971-ef1e3b373c6d", "organization_id": 101, "submitter_id": 1, "assignee_id": 40}
	ticketsDataRecords := []parser.DataRecord{newTicketsRecord_1}
	ticketsRows := db.CreateRows(ticketsDataRecords)
	ticketsTable := db.CreateTable(ticketsRows)
	database["tickets"] = ticketsTable
	return database
}
