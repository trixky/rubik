package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"math/rand"

	"github.com/trixky/rubik/parser"
	"github.com/trixky/rubik/models"
	"github.com/trixky/rubik/server"
)

func doCorrection() {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	arrTime := []float64{}
	arrSolLen := []float64{}

	for i := 0; i < 50; i++ {
		sequence := parser.RandomSequence(0, randGen.Intn(6))
		// fmt.Println("sequence ", i, ":", sequence)
		start := time.Now()
		result := models.SolveSequence(false, false, sequence)
		// fmt.Println("Solve ", result)
		duration := time.Since(start)
		// fmt.Println("Len of", len(result), "that took", duration.Milliseconds(), "milliseconds")
		// fmt.Println()

		arrSolLen = append(arrSolLen, float64(len(result)))
		arrTime = append(arrTime, float64(duration.Milliseconds()))
	}
	sumLen, sumTime := 0.0, 0.0
	for i := range arrTime {
		sumLen += arrSolLen[i]
		sumTime += arrTime[i]
	} 
	fmt.Println("50 solutions of size between 0 and 5:")
	fmt.Println("average solution length =", sumTime / 50, "moves")
	fmt.Println("average solution time   =", sumLen / 50, "ms")
	fmt.Println()


	arrTime = []float64{}
	arrSolLen = []float64{}

	for i := 0; i < 50; i++ {
		sequence := parser.RandomSequence(5, randGen.Intn(21))
		// fmt.Println("sequence ", i, ":", sequence)
		start := time.Now()
		result := models.SolveSequence(false, false, sequence)
		// fmt.Println("Solve ", result)
		duration := time.Since(start)
		// fmt.Println("Len of", len(result), "that took", duration.Milliseconds(), "milliseconds")
		// fmt.Println()

		arrSolLen = append(arrSolLen, float64(len(result)))
		arrTime = append(arrTime, float64(duration.Milliseconds()))
	}
	sumLen, sumTime = 0.0, 0.0
	for i := range arrTime {
		sumLen += arrSolLen[i] 
		sumTime += arrTime[i] 
	} 
	fmt.Println("50 solutions of size between 5 and 20:")
	fmt.Println("average solution length =", sumLen / 50, "moves")
	fmt.Println("average solution time   =", sumTime / 50, "ms")
	fmt.Println()


	arrTime = []float64{}
	arrSolLen = []float64{}

	for i := 0; i < 50; i++ {
		sequence := parser.RandomSequence(20, randGen.Intn(100))
		// fmt.Println("sequence ", i, ":", sequence)
		start := time.Now()
		result := models.SolveSequence(false, false, sequence)
		// fmt.Println("Solve ", result)
		duration := time.Since(start)
		// fmt.Println("Len of", len(result), "that took", duration.Milliseconds(), "milliseconds")
		// fmt.Println()

		arrSolLen = append(arrSolLen, float64(len(result)))
		arrTime = append(arrTime, float64(duration.Milliseconds()))
	}
	sumLen, sumTime = 0.0, 0.0
	for i := range arrTime {
		sumLen += arrSolLen[i] 
		sumTime += arrTime[i] 
	} 
	fmt.Println("50 solutions of size between 20 and 100:")
	fmt.Println("average solution length =", sumLen / 50, "moves")
	fmt.Println("average solution time   =", sumTime / 50, "ms")
	fmt.Println()

	superflips := []string{
		"U R2 F B R B2 R U2 L B2 R U' D' R2 F R' L B2 U",
	"F2 F2 L2 F2 U' B2 U2 F2 R' D' L1F' U' LBD' LFL' F' L" ,
	"L2 D2 L D2 R' F2 R2 B2 U F' D L' B' R D R F' U2 R U" ,
	"D' L2 F2 L R2 U B U' R' D' F B' D2 R' D2 R U2 R F2 U" ,
	"B2 U2 B2 L D2 U2 L2 B L' B2 D' F D' L U L2 B2 D B' R" ,
	"L F R U' D L' F D' F L2 F D R2 B2 D2 F U2 D2 F' U" ,
	"R2 B U2 B' U2 L2 F B2 L' B D' B U B2 D2 B' R' F' D " ,
	"L D2 F' L2 B D' U L2 U L U' R' U F' L2 U2 B U2 L2 U" ,
	"U L U' R' F L U B D' R2 B2 R F B2 D' B2 L2 D2 R2 D" ,
	"B' L2 U R2 L' F2 B U' R D' B L U' B D2 B' L2 B2 R2 F" ,
	"U2 F2 L2 U' F2 U R2 U2 F2 L' U2 B' D' L2 U L U2 F' R " ,
	"L B2 U2 R2 B2 D2 R U B' F2 D' F2 L R2 F' U2 B' U' R' D" ,
	"R D' R2 F' U B D F R U L' U2 B D' L2 U' L2 F2 U' F" ,
	"B2 F2 L F2 L2 B2 R2 F2 D L U' B' U' R2 U F2 D' L B' " ,
	"R B' L D' F D2 B' U' B R' U2 R2 L2 B2 D L2 U R2 U2 R" ,
	"D2 R2 B' D2 B2 D2 L2 B2 F' D' L' U2 R U' B' D' F L B2 R" ,
	"D2 R' U2 L' B2 D2 F2 B' R D2 B U' L2 D R U' B' F U F" ,
	"B U' B2 R2 D2 L' B' R F' D B L' F2 R' D2 R' L U2 F2 R" ,
	"L2 U2 B L2 B' L2 R2 U2 B R' B R2 F' U B F' L' D' F' U" ,
	"U2 F2 L F2 L R D2 R' D F2 U F' L' B' U R D F L U" ,
	"F2 D' F2 U2 B2 D' U2 F2 D L' U2 F D' R' F U B D' L " ,
	"F D2 B' R2 U2 F2 R2 L' D U' L R U L B U2 L' B2 U2 " ,
	"D2 L B2 L U2 L' F2 L2 D2 U2 L' F2 B R2 D' F' L' U' B' D" ,
	"L2 U2 B U' R' U D' L F' U' B U' L2 F2 L2 D L2 U L2 B" ,
	"L2 R' U2 F2 L D2 B2 D2 B2 F' D2 R' D U R' U2 B L F' " ,
	"D' R2 F2 U2 F B2 L U' F' R' U' B' D' R U2 D2 F2 L B2 L" ,
	"D2 L2 F' R2 F2 L2 F L2 D2 F2 D2 R' F2 D' U' F' U' B' F2 R" ,
	"F' D R2 D2 R' U2 B F' U L' F2 U' R2 B2 U' F2 U' F R " ,
	"F' R' F2 U R' B D' R' L2 F' L U' D2 F L F' R' U2 R' B" ,
	"B2 F2 L2 D2 F2 U2 R' B2 U B' R2 B R' F2 L' U F D' B' " ,
	"R D2 U2 F2 L' B2 R' U R B D' L' U' B' F' D' U L F2 L" ,
	"B L2 B2 D2 U2 B' R2 U2 D F' L U L' U B U' L R F2 U" ,
	"D2 F2 U2 R2 U2 R2 F' R' D' F' D' B' U F' R U' L B2 U' " ,
	"R' U R2 B' U' F D2 L B' U' R' B' U2 L2 U' L2 D' R2 D' B" ,
	"F' L2 F D2 B' U2 F' L' B2 D' U' R D F D2 L' D2 U2 B2 F" ,
	"B2 L U' R' D B2 U' R2 B' L' F' L' D R2 U2 D2 B2 D2 B D" ,
	"D L2 F2 R2 D B2 D' L2 U2 L B' R2 F D' L R' F' U' B F" ,
	"F L2 F' D2 B2 D2 B R2 D' L2 R' F L U' B2 L2 B F2 U' L" ,
	"L' U R2 U B D2 R2 D F L' U' B' D F2 U2 D2 R' L' B2 " ,
	"L' F U' F2 B U' L' F D' R' U D2 F2 D2 L F2 R D2 R' L" ,
	"F2 B' U' D' R' L2 F R' F2 B' L D R B2 D2 L2 U2 R' B2 " ,
	"B L' D2 B R2 U L' F' R' D F' B L2 U B2 U2 R2 U2 D F" ,
	"D F2 B' U' L F U' R2 D' L U R L2 B' R2 U2 R2 B2 D2 B" ,
	"U' R L B' U D B U2 R' F' D B R' B2 L' B2 D2 F2 U2 L" ,
	"U' L2 B2 U L2 B2 U2 L B D R U B L2 F D B2 F2 U2 " ,
	"R U D2 F' B R' F R' D' R' F' D L U2 R2 D2 F2 R B2 U" ,
	"D F B U R' F' B2 U2 R2 B U B2 D2 R F2 D2 F2 D2 L F" ,
	"R2 F2 D2 B2 L D2 F2 L2 F' D' F R U' L' U2 B D' L F R" ,
	"D2 B2 R2 B2 L' R U2 F2 D F L2 U' F2 D2 R2 B' L' U2 F' U" ,
	"U' L' F U2 R B2 D' F' B L2 D' U2 R' D2 R U2 L' F2 B2 L" ,
	"R2 B2 F R2 B D2 B2 L U2 B' L' U' L F D' F' L' F R F" ,
	"F2 R B2 R2 F2 L' R2 B2 F' D L U' B F D' L' B2 D U B" ,
	"B2 L2 R2 U2 F2 B D2 B' L B' D F B2 U' R D2 R' B2 U " ,
	"R D2 L' U2 D2 L D2 L B' D' F' U' B D L B2 D' F' L' F" ,
	"U2 L2 D2 B2 L2 U2 R' D2 B2 R B R2 D2 L2 U B R D L B" ,
	"R2 F D2 B' F2 U2 B' L2 D' L2 R2 D' B' D2 L R' U' B D " ,
	"D2 B2 F2 R U2 L2 B2 D L' U B2 F' R B L' D2 L' B2 L' B" ,
	"R F2 L2 F2 U2 R2 F2 L' U' F D2 L R D U' R2 F' U B' U" ,
	"R U2 L' D2 F L' D' B' R' U' D2 F2 L2 B' D2 F2 D2 R2 U2 L" ,
	"F' D2 B2 L2 D2 U2 B2 L2 U' B R F U' L D' B2 R2 F R D" ,
	"F R2 F2 D2 F' L2 F' L2 U F2 L2 R' U' F' L' D' B' R U B" ,
	"F2 R2 D U R2 U' F2 U' F' D U' R' D U' F L' R' F D' " ,
	"D2 U2 L2 B2 R' U2 R2 B R U' F2 L2 R2 B L2 U L' R U F" ,
	"F2 U2 L U2 L' R2 F2 L' R D F' U' L' F D2 R' B R' F' D" ,
	"D F R2 B L D R2 B R' U' L' U F2 U' L2 U' F2 U B2 " ,
	"B F2 L2 F' U2 B2 L2 F' D F' D L2 R D L' U2 B D2 U " ,
	"L2 R2 B2 U2 B L2 B' D2 U' B2 L' U2 L' B F U F2 R' B " ,
	"R' U2 B2 L F2 U2 B2 F' D2 R' B U2 F' D' U2 L' U' B' D' R" ,
	"D2 B2 F R2 U2 B2 F' U2 L' F2 R' U B' F' D U B U' R U" ,
	"L U' B D2 L' D' R U2 F' L2 D' B' R2 U' F2 U B2 L2 U2 L" ,
	"R2 B2 D L2 B2 L2 D' L B F L' D' F U F L' D' R2 U' R" ,
	"L2 F2 U2 L2 U L2 D B' D2 L R' F D' L' U' L B R F L" ,
	"L U2 L' F U' B R D F U' F2 U2 L' U2 D2 L' D2 L2 U2 " ,
	"D2 L2 F2 L' B2 F' R2 D L U B U2 F' D R B2 L' U2 L' R" ,
	"F D B R' L' F' L' D R2 U' L F2 L' U2 B2 D2 F2 R' F2 R" ,
	"D L' B' U2 R' D B' U R U2 B2 F' L B2 L2 U' B2 L2 F2 " ,
	"B2 F2 L B2 U2 L2 F2 L2 B2 F' D R' F R' D2 R' D' U B' R" ,
	"B R' D2 F D' R2 D' L F R2 U R' D F B D2 F2 B U2 D" ,
	"U' B2 U L2 U' B2 D2 L D2 F D R B L D' B2 F' R2 F " ,
	"D R L B' U F D2 B2 L' D' F' D' F U L2 U L2 F2 U D" ,
	"D2 R2 D2 R' U2 F2 U2 B2 U' R2 B2 F' L B U' B2 R' D' R2 F" ,
	"D' L2 D U2 L2 U2 R2 F2 R' U2 F' L' D' L U' L B' F U2 " ,
	"D' F R' B2 D' B' U L D' F' R B2 U2 L2 F2 R2 D' R2 U' D" ,
	"F U2 R2 B' L2 R2 F R2 D' B2 L' F D R' U R2 B' U L' F" ,
	"L2 U2 F2 U L2 U' L2 D' L F' L2 R B R2 U' B D2 R D' F" ,
	"U2 B2 D2 L2 U2 L2 R B F' L D2 F' D' F U' B' U R' D R" ,
	"R U2 L2 U2 B2 U2 F2 U2 L U L D2 B' L' F' D' L U' R " ,
	"D2 B2 R' L2 F2 L D' L' F2 U2 F2 B R' B2 L' U' F' U' D' B" ,
	"U2 L2 F2 U2 B2 F2 L' F2 U B' D' L' F' U B L2 B2 R U' " ,
	"B' L F2 D2 F U' L' U D' R' F L2 F L2 U2 F B R2 D2 F" ,
	"U' D' B' R2 B2 D L2 F R L' F L' D2 F' B2 D2 F2 U' L2 " ,
	"B L2 U R F' D' B' L' U R F' U B2 U2 L2 F D2 B R2 U" ,
	"L B' R F U' B2 L' D F D' R' U R' U2 L' F2 D2 R' B2 " ,
	"F2 D' F2 L2 D' L2 B2 L' B F2 L' D U L F R' B D U' F" ,
	"F U' B D2 R B' L' B D' F U' F2 B2 R D2 B2 R D2 L' B" ,
	"B2 D2 B' L2 R2 U2 B' D' L' R2 U' R2 F' R' F' D L' U' F U" ,
	"R2 B' U R B' L' D F' U' L D2 L F' B' L2 D2 F B2 U2 R" ,
	"U2 R' B2 L F2 D2 F2 R F' D' R' B D2 R2 U2 B2 D F2 R B'",
	"L2 F2 U2 L2 B2 D2 F R' U L R2 F' U2 F' U F2 U2 B L2 R",
	"D2 U2 F2 R U2 R2 D2 L' D2 F U F L' D2 F' L B R' D' F'",
	"F2 U F' B D R U F' R2 F U L R2 B2 L R U2 F2 B2 R",
	}
	arrTime = []float64{}
	arrSolLen = []float64{}

	for _, val := range superflips {
		start := time.Now()
		result := models.SolveSequence(false, false, strings.Fields(val))
		// fmt.Println("Solve ", result)
		duration := time.Since(start)
		// fmt.Println("Len of", len(result), "that took", duration.Milliseconds(), "milliseconds")
		// fmt.Println()

		arrSolLen = append(arrSolLen, float64(len(result)))
		arrTime = append(arrTime, float64(duration.Milliseconds()))
	}
	sumLen, sumTime = 0.0, 0.0
	for i := range arrTime {
		sumLen += arrSolLen[i] 
		sumTime += arrTime[i] 
	} 
	fmt.Println("101 Superflips:")
	fmt.Println("average solution length =", sumLen / float64(len(superflips)), "moves")
	fmt.Println("average solution time   =", sumTime / float64(len(superflips)), "ms")
	fmt.Println()

}

func main() {
	// ----------------------------------------- thervieu
	api_mode := len(os.Args) == 1

	if api_mode {
		server.Start()
	} else {
		verbose, random, correction, sequence := parser.ReadArgs()
		
		if correction == true {
			doCorrection()
			os.Exit(0)
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
