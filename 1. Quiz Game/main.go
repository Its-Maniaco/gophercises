package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("csv", "problems.csv", "Pass the file name containing the problems with extension")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	// Since file will not be too big, we will read it at once.
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf(err.Error())
	}
	problems := parseLine(lines)
	correctAnsCount := 0
	for i, p := range problems {
		fmt.Println("Problem #", i+1, " : ", p.question, "= ")
		var ans string
		fmt.Scanf("%s\n", &ans) //reads input (here) in format - 'string enter', then it will read new input
		if ans == p.answer {
			correctAnsCount++
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
