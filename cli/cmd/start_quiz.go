package cmd

import (
	"bufio"
	"context"
	"fmt"
	d "github.com/eu-ga/quiz/cli/display"
	pb "github.com/eu-ga/quiz/proto"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

func StartCmd(Client pb.QuizClient) *cobra.Command {
	inputValues, err := Client.PerQuizQuestions(context.Background(), &pb.NumberOfQuestions{})
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	c := &cobra.Command{
		Use:   "start",
		Short: "Start the quiz",
		Long: `Quiz have ` + fmt.Sprint(inputValues.Questions) + ` questions. Answer them all by typing in number 
of the proper answer. After last question you will get results. Good luck!`,
		Run: func(cmd *cobra.Command, args []string) {
			user := pb.User{
				Id: time.Now().Unix(),
			}
			r, err := Client.GetQuestions(context.Background(), &user)
			if err != nil {
				fmt.Println("Sorry. Cannot get questions. Please connect your manager.")
				log.Fatal(err.Error())
				return
			}
			anw := getAnswers(user.Id, r.Questions)
			result, err := Client.SendAnswers(context.Background(), &anw)
			if err != nil {
				fmt.Println("Sorry. Cannot get results. Please connect your manager.")
				return
			} else {
				d.DisplayResult(result)
			}
		},
	}
	return c
}

func scanAnswer() int64 {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {

		}
		return int64(n)
	}
	if scanner.Err() != nil {
		log.Println("Err0r: ", scanner.Err().Error())
	}
	return -1
}

func getAnswers(userId int64, questions []*pb.Question) pb.PostAnswers {
	var s []*pb.Solution
	for _, q := range questions {
		d.DisplayQuestion(q)
		a := scanAnswer()
		tmp := pb.Solution{
			QuestionId: q.Id,
			AnswerId:   a,
		}
		s = append(s, &tmp)
	}
	result := pb.PostAnswers{
		UserId:   userId,
		Solution: s,
	}
	return result
}
