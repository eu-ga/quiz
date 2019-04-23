//storage package is a dumb implementation of data storage functionality
package storage

import (
	"github.com/eu-ga/go_utils"
	pb "github.com/eu-ga/quiz/proto"
	"sync"
)

//simulation fo the data storage
var Cache DataStore

type dataStore interface {
	AddUser(pb.User) (*pb.User, bool)
	GetUser(int64) *pb.User
	AddQuestion(pb.Question) (*pb.Question, bool)
	SaveSolution(pb.Solution) bool
}

//In-memory caching system
type DataStore struct {
	sync.RWMutex
	PerQuizQuestions int64
	Users            map[int64]*User
	Questions        map[int64]*pb.Question
	Statistics       []UserStatistics
}

type User struct {
	pb.User
	Statistics UserStatistics
}

type UserStatistics struct {
	CorrectAnswers int64
	SuccessRate    int64
	UserId         int64
}

func (u *User) update(us *User) {
	u.Name = us.Name
	u.Statistics = us.Statistics
}

//AddUser puts user info into storage for statistics
func (ds *DataStore) AddUser(user *User) (*User, bool) {
	ds.Lock()
	defer ds.Unlock()
	isExists := false
	if u := ds.GetUser(user.Id); u != nil {
		u.update(user)
	} else {
		ds.Users[user.Id] = user
		isExists = true
	}
	return user, isExists
}

//GetUser checks if user exists and returns it
func (ds *DataStore) GetUser(id int64) *User {
	if u, ok := ds.Users[id]; ok {
		return u
	}
	return nil
}

//AddQuestion adds new questions into the storage
func (ds *DataStore) AddQuestion(question pb.Question) (*pb.Question, bool) {
	ds.Lock()
	defer ds.Unlock()
	if _, ok := ds.Questions[question.Id]; !ok {
		ds.Questions[question.Id] = &question
		return &question, true
	}
	return nil, false
}

//SaveSolution puts result of the quiz into storage,
//so app can use it for the response for users
func (ds *DataStore) SaveSolution(solution pb.Solution) bool {
	ds.Lock()
	defer ds.Unlock()
	return false
}

//UpdateStatistics keeps statistics up to date, and users will get proper data in the response
func (ds *DataStore) UpdateStatistics(data []UserStatistics) {
	ds.Lock()
	defer ds.Unlock()
	ds.Statistics = data
}

//LoadQuestions loads questions from questions.json file into memory
func (ds *DataStore) LoadQuestions() {
	filePath, err := go_utils.GetCurrentDir()
	if err != nil {
		panic(err.Error())
	}
	var questions []pb.Question
	if err := readConf(filePath+"/questions.json", &questions); err != nil {
		panic(err.Error())
	}
	for _, q := range questions {
		ds.AddQuestion(q)
	}

}
