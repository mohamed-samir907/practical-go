package main

import "strings"

type problem struct {
	question string
	answer   string
}

// parseLines is used to convert the csv file lines
// to array of problems
func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func (p *problem) isCorrectAnswer(answer string) bool {
	return p.answer == answer
}
