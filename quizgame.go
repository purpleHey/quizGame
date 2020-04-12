package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	fileName := flag.String("csv", "problems.csv", "A file of format blah")
	timeLimit := flag.Int("tLimit", 15, "The number of seconds you have to take the quiz")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Could not read in file %s\n", *fileName)
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided .csv file")
	}
	numCorrect := 0

	problems := parseLines(lines)
	fmt.Println(problems)

	fmt.Printf("You have %d seconds to complete the quiz\n", *timeLimit)
	fmt.Printf("Type enter to start the quiz\n")
	reader := bufio.NewReader(os.Stdin)
	char, _ := reader.ReadString('\n')
	if len(char) > 0 {

		// Print the current time
		fmt.Println(time.Now())
		timer1 := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		numAttempted := 0

		reader := bufio.NewReader(os.Stdin)
		for _, p := range problems {
			fmt.Printf("%s = ?: ", p.q)
			answerCh := make(chan string)
			go func() {
				var answer string
				answer, _ = reader.ReadString('\n')
				answerCh <- answer
			}()
			select {
			case <-timer1.C:
				// Time's up
				fmt.Println(time.Now())
				fmt.Printf("\nScore: %d - %d out of %d (%d not attempted)\n",
					numCorrect/numAttempted*100, numCorrect, numAttempted, len(problems)-numAttempted)
				return
			case answer := <-answerCh:
				numAttempted++
				if strings.TrimSuffix(answer, "\n") == p.a {
					fmt.Printf("Correct\n")
					numCorrect++
				} else {
					fmt.Printf("Dooh! Correct answer is %s\n", p.a)
				}
			}
		}
	}
	// fmt.Println("Getting back go...ing with 1.14")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
func parseLines(lines [][]string) []mathProblem {
	ret := make([]mathProblem, len(lines))

	for i, line := range lines {
		ret[i] = mathProblem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type mathProblem struct {
	q string
	a string
}
