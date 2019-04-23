package service

import (
	"fmt"
	pb "github.com/eu-ga/quiz/proto"
	"github.com/eu-ga/quiz/server/storage"

	"github.com/stretchr/testify/assert"
	"testing"
)

func init(){
	q1 := &pb.Question{
		Id:1,
		Body:"Question number one.",
		CorrectAnswerId:1,
		Answers:[]*pb.Answer{
			{
				Id:1,
				QuestionId:1,
				Body:"Answer number 1.",
			},{
				Id:2,
				QuestionId:1,
				Body:"Answer number 2.",
			},{
				Id:3,
				QuestionId:1,
				Body:"Answer number 3.",
			},{
				Id:4,
				QuestionId:1,
				Body:"Answer number 4.",
			},

		},
	}
	q2 := &pb.Question{
		Id:2,
		Body:"Question number two.",
		CorrectAnswerId:2,
		Answers:[]*pb.Answer{
			{
				Id:1,
				QuestionId:2,
				Body:"Answer number 1.",
			},{
				Id:2,
				QuestionId:2,
				Body:"Answer number 2.",
			},{
				Id:3,
				QuestionId:2,
				Body:"Answer number 3.",
			},{
				Id:4,
				QuestionId:2,
				Body:"Answer number 4.",
			},

		},
	}
	q3 := &pb.Question{
		Id:3,
		Body:"Question number 3.",
		CorrectAnswerId:3,
		Answers:[]*pb.Answer{
			{
				Id:1,
				QuestionId:3,
				Body:"Answer number 1.",
			},{
				Id:2,
				QuestionId:3,
				Body:"Answer number 2.",
			},{
				Id:3,
				QuestionId:3,
				Body:"Answer number 3.",
			},{
				Id:4,
				QuestionId:3,
				Body:"Answer number 4.",
			},

		},
	}
	q4 := &pb.Question{
		Id:4,
		Body:"Question number 4.",
		CorrectAnswerId:4,
		Answers:[]*pb.Answer{
			{
				Id:1,
				QuestionId:4,
				Body:"Answer number 1.",
			},{
				Id:2,
				QuestionId:4,
				Body:"Answer number 2.",
			},{
				Id:3,
				QuestionId:4,
				Body:"Answer number 3.",
			},{
				Id:4,
				QuestionId:4,
				Body:"Answer number 4.",
			},

		},
	}
	storage.Cache.Questions[q1.Id] = q1
	storage.Cache.Questions[q2.Id] = q2
	storage.Cache.Questions[q3.Id] = q3
	storage.Cache.Questions[q4.Id] = q4
	storage.Cache.PerQuizQuestions = 3
}

func TestGetRandomQuestions(t *testing.T){
	got := getRandomQuestions()
	assert.True(t,
		int64(len(got))==storage.Cache.PerQuizQuestions,
		fmt.Sprintf("Got too many questions. Got: %d, Expected: %d", len(got), storage.Cache.PerQuizQuestions))
	test := make(map[int64]interface{})
	for _,q := range got{
		if _,ok := test[q.Id];ok{
			assert.Fail(t, fmt.Sprintf("Duplicated questions. Id: %d",q.Id))
		}else {
			test[q.Id] = q
		}
	}
}

func TestUpdateStatistics(t *testing.T) {
	storage.Cache.PerQuizQuestions = 5
	data := []storage.UserStatistics{}
	el1 := storage.UserStatistics{
		SuccessRate:60,
	}
	el2 := storage.UserStatistics{
		SuccessRate:70,
	}
	el3 := storage.UserStatistics{
		SuccessRate:10,
	}
	data, _ = storage.InsertSort(data,el1)
	data, _ = storage.InsertSort(data,el2)
	data, _ = storage.InsertSort(data,el3)
	storage.Cache.UpdateStatistics(data)
	correctAnswers := int64(0)
	res := UpdateStatistics(correctAnswers, 0)
	assert.True(t, res.SuccessRate==int64(0),fmt.Sprintf("Error: should be %d, got %d",0, res.SuccessRate))
	correctAnswers = int64(0)
	res = UpdateStatistics(correctAnswers, 0)
	assert.True(t, res.SuccessRate==int64(40),fmt.Sprintf("Error: should be %d, got %d",20, res.SuccessRate))
	correctAnswers = int64(1)
	res = UpdateStatistics(correctAnswers, 0)
	assert.True(t, res.SuccessRate==int64(66),fmt.Sprintf("Error: should be %d, got %d",60, res.SuccessRate))
}