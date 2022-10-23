package parser

import (
	"os"
	"fmt"
	"time"
	"strings"
    "math/rand"
	"github.com/trixky/rubik/models"
)

// prints help and exits
func help() {
	fmt.Println("help:")
	fmt.Println("options available [-h, --help] [-r, --random] [-v, --verbose]")
	fmt.Println("  -h : prints this message and exits")
	fmt.Println("  -r : creates a random sequence, overrides the sequence if one was given")
	fmt.Println("  -v : writes a more detailed solve (sequence, length of sequence and time per group transformation)")
	fmt.Println("  -c : correction subject (takes ~10 minutes)")
	fmt.Println("  -p : plot for correction, overrides existing png")
	fmt.Println("Example : go run . -v \"F R U2 B' R' L D2\"")
	os.Exit(0)
}


// prints usage and wrong input and exits
func wrongInputUsage(wrong string) {
	fmt.Printf("rubik: wrong input: %s\n", wrong)
	fmt.Println("Valid characters are 'UDFRBL' and can be followed by either ' (reverse) and 2 (twice).")
	fmt.Println("Each move seperated by whitespace(s).")
	fmt.Println("Can't input two sequences.")
	fmt.Println("Example : go run . \"F R U2 B' R' L D2\"")
	os.Exit(0)
}


/*
	checks that the shuffle is valid:
		e.g words are within : [F, F', F2, B, B', B2, ...]
	@argument words: the split shuffle
	return value : true if valid, false otherwise and a message saying what's wrong
*/
func isValidShuffle(words []string) (bool, string) {
	var valid_chars = "UDFRBL"
	var valid_sec_chars = "'2"

	for _, s := range words {
		if len(s) > 2 {
			return false, "sequence: \"" + s + "\" is too long"
		}
		if strings.Contains(valid_chars, string(s[0])) == false {
			return false, "sequence: \"" + s + "\" is not accepted"
		}
		if len(s) == 2 {
			if strings.Contains(valid_sec_chars, string(s[1])) == false {
				return false, "sequence: \"" + s + "\" is not accepted"
			}
		}
	}
	return true, ""
}


/*
	reads arguments and returns info about the execution of the program

	return values 
		verbose, random: boolean depicting those options
		api_mode: false if no api, yes if it has to start
*/
func ReadArgs() (verbose bool, random bool, correction bool, plot bool, sequence []string) {
	sequence_exists := false
	verbose, random, correction, plot = false, false, false, false
	sequence = []string{}

	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			help()
		} else if arg == "-v" || arg == "--verbose" {
			verbose = true
		} else if arg == "-r" || arg == "--random" {
			random = true
		} else if arg == "-c" || arg == "--correction" {
			correction = true
		} else if arg == "-p" || arg == "--plot" {
			plot = true
		} else {
			if (len(arg) > 0 && string(arg[0]) == "-") {
				wrongInputUsage("option: " + arg + " is not an option")
			}
			if sequence_exists == true {
				wrongInputUsage("two sequences: second is \"" + arg + "\"")
			}
			words := strings.Fields(arg)
			valid, which := isValidShuffle(words)
			if valid == false {
				wrongInputUsage(which)
			} else {
				sequence = words
				sequence_exists = true
			}
		}
	}

	if correction {
		return
	}
	if sequence_exists == false && random == false {
		wrongInputUsage("no sequence (written or random)")
	}

	return
}

// returns a random sequence, overrides the one given in argument
func RandomSequence(min int, length int) []string {
	if min > 0 && length < min {
		length = min
	}

	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	sequence := []string{}

	moves := models.GetGroupMoves(0, models.Cube{})
	for i := 0; i < length; i++ {
		sequence = append(sequence, moves[randGen.Intn(len(moves))],)
	}
	return sequence
}
