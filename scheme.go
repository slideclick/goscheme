package main

import (
	"os"

	"github.com/slideclick/goscheme/repl"
)

func main() {
	if len(os.Args) < 2 {
		// TODO: REPL
		repl.Start()
		return
	}

}
