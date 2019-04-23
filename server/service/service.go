//service package provides functionality for RCP requests
//that configured in ./proto/quiz.proto
package service

import (
	"context"
	pb "github.com/eu-ga/quiz/proto"
	"github.com/eu-ga/quiz/server/storage"
	"log"
	"math/rand"
)

type Service struct {
	Cache storage.DataStore
}

//GetQuestions RCP request handler. Returns random set of questions
func (s *Service) GetQuestions(ctx context.Context, user *pb.User) (*pb.Response, error) {
	var res = pb.Response{
		Questions: getRandomQuestions(),
		UserId:    user.Id,
	}
	return &res, nil
}

//SendAnswers RCP request handler that updates statistics for the proper response
// with user success rate(how he performed quiz comparing to the others)
func (s *Service) SendAnswers(ctx context.Context, answers *pb.PostAnswers) (*pb.Result, error) {
	var res pb.Result
	correctAnswers, err := checkAnswers(answers)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	} else {
		UpdateStatistics(correctAnswers, answers.UserId)
		res.CorrectAnswers = correctAnswers
		res.SuccessRate = storage.Cache.Users[answers.UserId].Statistics.SuccessRate
	}
	return &res, nil
}

func checkAnswers(answer *pb.PostAnswers) (int64, error) {
	var counter int64
	for _, a := range answer.Solution {
		corAnswer := storage.Cache.Questions[a.QuestionId].CorrectAnswerId
		if corAnswer == a.AnswerId {
			counter++
		}
	}
	return counter, nil
}

//getRandomQuestions dumb function that pretend to be question selector
func getRandomQuestions() []*pb.Question {
	var res []*pb.Question
	l := len(storage.Cache.Questions)
	n := int64(rand.Intn(l-1) + 1)
	amount := int64(len(storage.Cache.Questions))
	for i := n; i <= amount; i++ {
		q := storage.Cache.Questions[i]
		if res = append(res, q); int64(len(res)) == storage.Cache.PerQuizQuestions {
			return res
		}
		if i == amount {
			i = 0
		}
	}
	return res
}

//UpdateStatistics sets percents of correct answers for sorted storage
//and returns number of correct answers and value how user compared to others that have taken the quiz
func UpdateStatistics(correctAnswers, userId int64) *pb.Result {
	oneQuestion := 100 / storage.Cache.PerQuizQuestions
	rate := oneQuestion * correctAnswers
	tmp := storage.UserStatistics{
		UserId:         userId,
		CorrectAnswers: correctAnswers,
		SuccessRate:    rate,
	}
	data, i := storage.InsertSort(storage.Cache.Statistics, tmp)
	storage.Cache.UpdateStatistics(data)
	res := pb.Result{
		CorrectAnswers: tmp.CorrectAnswers,
	}
	if i > 0 {
		per := int64((int64(i) + 1)) * 100
		r := per / int64(len(data))
		res.SuccessRate = r
	} else {
		res.SuccessRate = 0
	}
	return &res
}
