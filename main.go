package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("my-app", "My Example CLI Application With Bash Completion")
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
