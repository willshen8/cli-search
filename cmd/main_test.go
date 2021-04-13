package main

import (
	"os"
	"testing"

	"github.com/willshen8/cli-search/internal/copy"
	"github.com/willshen8/cli-search/internal/errors"
)

func TestMain(m *testing.M) {
	os.Args = []string{"query", "organizations", "_id", "101"}
	setup()
	exitCode := m.Run()
	tearDown()
	os.Exit(exitCode)
}

func setup() {
	err := os.Mkdir("config", 0777)
	errors.HandleError(err)
	err = copy.CopyFile("../config/organizations.json", "config/organizations.json")
	errors.HandleError(err)
	err = copy.CopyFile("../config/users.json", "config/users.json")
	errors.HandleError(err)
	err = copy.CopyFile("../config/tickets.json", "config/tickets.json")
	errors.HandleError(err)
}

func tearDown() {
	err := os.RemoveAll("config")
	errors.HandleError(err)
}
