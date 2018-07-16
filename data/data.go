package data

import (
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type Thread struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Topic     string    `json:"topic"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type IndexData struct {
	Threads []Thread
	Users   []User
}

type TestData struct {
	Title string
	Users []User
}

func (t Thread) NumReplies() int {
	return 7
}

func ThreadList() []Thread {
	return []Thread{
		Thread{1, "1", "First Topic", 1, time.Now()},
		Thread{2, "2", "Second Topic", 2, time.Now()},
		Thread{3, "3", "Third Topic", 3, time.Now()},
	}
}

func UserByEmail(email string) (User, error) {
	return User{ID: 1, Name: "Petya", Email: "petya@google.com"}, nil
}

var Users []User

func CreateUser(user User) {
	Users = append(Users, user)
}
