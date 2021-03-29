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
	id := []string{"101"}

	actual := captureOutput(func() {
		PrintResults(dB, "dummyTable", id)
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInOrgTable(t *testing.T) {
	dB := CreateDB()
	actual := captureOutput(func() {
		PrintAllAvailableFields(dB, "organizations")
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInUserTable(t *testing.T) {
	dB := CreateDB()
	actual := captureOutput(func() {
		PrintAllAvailableFields(dB, "users")
	})
	assert.NotEmpty(t, actual)
}

func TestPrintAllAvailableFieldsInTicketTable(t *testing.T) {
	dB := CreateDB()
	actual := captureOutput(func() {
		PrintAllAvailableFields(dB, "tickets")
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

func CreateDB() db.DB {
	database := db.DB{}
	newDataRecord_1 := parser.DataRecord{"_id": 101, "url": "http://initech.zendesk.com/api/v2/organizations/101.json"}
	newDataRecord_2 := parser.DataRecord{"_id": 102, "url": "http://initech.zendesk.com/api/v2/organizations/102.json"}
	dataRecords := []parser.DataRecord{newDataRecord_1, newDataRecord_2}
	rows := db.CreateRows(dataRecords)
	newTable := db.CreateTable(rows)
	database["dummyTable"] = newTable
	return database
}
