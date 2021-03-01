package db

import (
	"fmt"
	"os"

	"github.com/willshen8/cli-search/internal/errors"
	"github.com/willshen8/cli-search/pkg/parser"
)

// type DB stores database tables
type DB map[string]Table

// Table contains rows of data
type Table map[string]Row

// Row is the basic data container that
type Row map[string]interface{}

// CreateDB returns a pointer a new DB created
func CreateDB() *DB {
	return &DB{}
}

// CreateTable makes a new table with a name from the json file suffix and its data - rows
func CreateTable(rows []Row) Table {
	newTable := Table{}
	for _, row := range rows {
		newTable[string(fmt.Sprintf("%v", row["_id"]))] = row
	}
	return newTable
}

// CreateRows takes a slice of data and returns a slice of rows
func CreateRows(dataSlice []parser.DataRecord) []Row {
	var rows []Row
	for _, v := range dataSlice {
		row := Row{}
		for key, val := range v {
			row[key] = val
		}
		rows = append(rows, row)
	}
	return rows
}

func CreateTableFromJsonFile(file string) (Table, error) {
	newTable := Table{}
	openedFile, err := os.Open(file)
	if err != nil {
		return newTable, errors.NewError(err, fmt.Sprintf("Error opening file %s", file))
	}
	data, err := parser.ReadJsonFile(openedFile)
	if err != nil {
		return newTable, errors.NewError(err, fmt.Sprintf("Error reading file %s", file))
	}
	dataRows, err := parser.UnmarshalData(data)
	if err != nil {
		return newTable, errors.NewError(err, "Error unmarshal data into data records.")
	}
	rows := CreateRows(dataRows)
	createdTable := CreateTable(rows)
	return createdTable, nil
}
