package input

import (
	"log"
	"os"
)

func readArg() (input string) {
	if len(os.Args) == 2 {
		input = os.Args[1]
	} else if len(os.Args) > 2 {
		log.Fatal("only one argument is required")
	}

	return
}
