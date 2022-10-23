package main

import (
	"os"
	"fmt"
	"time"
	"math/rand"


	"github.com/trixky/rubik/parser"
	"github.com/trixky/rubik/models"
	"github.com/trixky/rubik/server"
	"github.com/trixky/rubik/correction"
)


func main() {
	// ----------------------------------------- thervieu
	api_mode := len(os.Args) == 1

	if api_mode {
		server.Start()
	} else {
		verbose, random, correc, plot, sequence := parser.ReadArgs()
		
		if correc == true {
			correction.Correction(plot)
			return
		}
		if random == true {
			source := rand.NewSource(time.Now().UnixNano())
			randGen := rand.New(source)

			sequence = parser.RandomSequence(-1, randGen.Intn(51))
		}

		fmt.Println("sequence:")
		fmt.Println(sequence)
		start := time.Now()
		result := models.SolveSequence(false, verbose, sequence)
		fmt.Println("Solve sequence:")
		fmt.Println(result)
		duration := time.Since(start)
		fmt.Println("Len : ", len(result))
		fmt.Println("Time :", duration.Milliseconds(), "milliseconds")
	}
}
