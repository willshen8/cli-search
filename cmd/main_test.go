package main

import (
	"os"
	"testing"

	"gopkg.in/alecthomas/kingpin.v2"
)

func TestMain(m *testing.M) {
	os.Args = []string{"query", "organisation", "_id", "101"}
	os.Exit(m.Run())
	kingpin.MustParse(app.Parse(os.Args))
}
