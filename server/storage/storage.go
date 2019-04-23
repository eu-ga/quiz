package storage

import(
	pb "github.com/eu-ga/quiz/proto"
	"sync"
)

//simulation fo the data storage
var Cache DataStore

type dataStore interface{
	AddUser(pb.User)(*pb.User,bool)
	GetUser(int64)(*pb.User)
	AddQuestion(pb.Question)(*pb.Question,bool)
	SaveSolution(pb.Solution)(bool)
}

//In-memory caching system
type DataStore struct{
	sync.RWMutex
	PerQuizQuestions int64
	Users map[int64]*User
	Questions map[int64]*pb.Question
	Statistics []UserStatistics
}


type User struct{
	pb.User
	Statistics UserStatistics
}

type UserStatistics struct{
	CorrectAnswers int64
	SuccessRate int64
	UserId int64
}

func (u *User) update(us *User){
	u.Name = us.Name
	u.Statistics = us.Statistics
}

func (ds *DataStore)AddUser(user *User)(*User,bool){
	ds.Lock()
	defer ds.Unlock()
	isExists := false
	if u := ds.GetUser(user.Id);u != nil{
		u.update(user)
	}else{
		ds.Users[user.Id]=user
		isExists = true
	}
	return user,isExists
}

func (ds *DataStore)GetUser(id int64)*User{
	if u,ok:=ds.Users[id];ok{
		return u
	}
	return nil
}

func (ds *DataStore)AddQuestion(question pb.Question)(*pb.Question,bool){
	ds.Lock()
	defer ds.Unlock()
	return nil,false
}

func (ds *DataStore)SaveSolution(solution pb.Solution)(bool){
	ds.Lock()
	defer ds.Unlock()
	return false
}

func (ds *DataStore)UpdateStatistics(data []UserStatistics){
	ds.Lock()
	defer ds.Unlock()
	ds.Statistics = data
}