package main

import (
	"fmt"
	"time"
)

// run - Run the quiz
func run(problems []problem, quizTime time.Duration) int {
	timer := createTimer(quizTime)
	correct := 0 // num of correct answers

problemsLoop: // label
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		// create channel for the answer and set it's size to 1 answer
		answerChan := make(chan string, 1)

		// we will get the answer in goroutine because the fmt.Scanf
		// blocks the code until we hit "Enter"
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			// send the answer to the answer channel
			answerChan <- answer
		}()

		select {
		// wait for the quiz time to the end. when it end it will break the for loop.
		case <-timer:
			break problemsLoop
		// wait for the answer, when we get an answer we will check if it correct or not.
		case answer := <-answerChan:
			if problem.isCorrectAnswer(answer) {
				correct++
			}
		}
	}

	return correct
}

// createTimer - Create the timer for the quiz
func createTimer(quizTime time.Duration) <-chan time.Time {
	timer := time.NewTimer(quizTime * time.Second)
	return timer.C
}
