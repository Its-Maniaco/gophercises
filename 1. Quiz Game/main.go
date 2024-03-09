package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Pass the file name")
	timeLimit := flag.Int("time", 5, "Set time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll() // Since file will not be too big, we will read it at once.
	if err != nil {
		log.Fatalf(err.Error())
	}
	problems := parseLine(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correctAnsCount := 0
	/*
		We want program to end when time expires, but also dont want to take an ans and then end
		So we run scanf into seperate goroutine so it runs seperately to timer, ques display & program exit.
	*/
	for i, p := range problems {
		fmt.Println("Problem #", i+1, " : ", p.question, "= ")
		ansCh := make(chan string) //to store ans, as we will need to access value outside of anon func
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans) //reads input in format (here) - 'string enter', then it will read new input
			ansCh <- ans
		}()

		select {
		case <-timer.C: //if timer fires
			fmt.Println("Time ran out.")
			fmt.Println("Number of correct ans: ", correctAnsCount)
			return
		case answer := <-ansCh:
			if answer == p.answer {
				correctAnsCount++
			}
		}
	}
	fmt.Println("Number of correct ans: ", correctAnsCount)

}

// We unmarshal into a struct
func parseLine(lines [][]string) []problem {
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}
	return result
}
