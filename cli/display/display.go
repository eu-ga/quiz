//display package contains functions for displaying data in command line
package display

import (
	"fmt"
	pb "github.com/eu-ga/quiz/proto"
)

//Display makes single line more visible
func Display(msg string) {
	fmt.Print("###############   ")
	fmt.Print(msg)
	fmt.Println("   ###############")
}

//DisplayList makes list from array of strings
func DisplayList(msgs ...string) {
	for i, msg := range msgs {
		fmt.Print(i + 1)
		fmt.Print(") ")
		fmt.Println(msg)
	}
}

//DisplayResult shows result of the quiz
func DisplayResult(r *pb.Result) {
	success := r.SuccessRate - 1
	if success <= 0 {
		success = 0
	}
	fmt.Println("You have ", r.CorrectAnswers, " correct answers.")
	fmt.Println("You better than ", success, "% users. Congratulations!")
}

//DisplayQuestion displays question in a proper form
func DisplayQuestion(q *pb.Question) {
	Display(q.Body)
	var tmp []string
	for _, a := range q.Answers {
		tmp = append(tmp, a.Body)
	}
	DisplayList(tmp...)
	Display("")
	fmt.Print("Enter your answer: ")
}
