package main

import (
	"os"

	"github.com/tobiashort/monkey/repl"
)

func main() {
	repl.Start(os.Stdout, os.Stdin)
}
