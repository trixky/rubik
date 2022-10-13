package main

import (
	"os"
	"fmt"
	"time"

	"github.com/trixky/rubik/parser"
	"github.com/trixky/rubik/models"
	"github.com/trixky/rubik/server"
)

func main() {
	// ----------------------------------------- thervieu
	api_mode := len(os.Args) == 1

	// tables := getTables()

	if api_mode {
		server.Start()
	} else {
		verbose, random, sequence := parser.ReadArgs()
		
		if random == true {
			sequence = parser.RandomSequence()
		}

		fmt.Println("sequence ", sequence)
		start := time.Now()
		result := models.SolveSequence(verbose, sequence)
		fmt.Println("total solve took ", time.Since(start), " seconds")
		fmt.Println("solveSequence is ", result)
	}
}
