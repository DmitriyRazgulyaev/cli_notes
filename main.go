package main

import (
	"cli_notes/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()

}
