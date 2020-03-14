package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/PaesslerAG/gval"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "A file of format blah")
	flag.Parse()
	// _ = fileName
	dat, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatalf("Could not read in file %s\n", *fileName)
	}
	// fmt.Print(string(dat))

	// r := csv.NewReader(strings.NewReader(string(dat)))
	r := csv.NewReader(strings.NewReader(string(dat)))

	numCorrect := 0
	numProblems := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		numProblems++
		if err != nil {
			log.Fatal()
		}

		value, err := gval.Evaluate(record[0], "")
		// fmt.Print(value)

		fVal, err := strconv.ParseFloat(record[1], 64)
		if value.(float64) == fVal {
			numCorrect++
		} else {
			fmt.Printf("%v\n", record)
		}
	}
	fmt.Printf("Total Num Problems: %d, Total Correct %d\n", numProblems, numCorrect)
	// fmt.Println("Getting back go...ing with 1.14")
}
