package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer'")
	timeLimit := flag.Int("time", 0, "the time limit of the quiz")
	flag.Parse()

	file := openFile(*csvFile)
	lines := readLines(file)

	problems := parseLines(lines)

	correct := run(problems, time.Duration(*timeLimit))
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}

// readLines - Read the csv file lines
func readLines(reader io.Reader) [][]string {
	r := csv.NewReader(reader)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines
}

// openFile - Open the given file name.
func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		exit(fmt.Sprintf("Can not open the CSV file %s", name))
	}

	return file
}

// exit - Print a message and exit the program.
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
