package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	flag.Parse()

	file := openFile(*csvFile)
	lines := readLines(file)

	correct := 0
	problems := parseLines(lines)
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if problem.isCorrectAnswer(answer) {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

// Read the csv file lines
func readLines(reader io.Reader) [][]string {
	r := csv.NewReader(reader)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines
}

// Open the given file name.
func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		exit(fmt.Sprintf("Can not open the CSV file %s", name))
	}

	return file
}

// Print a message and exit the program.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
