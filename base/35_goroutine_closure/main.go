package main

import (
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

var users []User

func init() {
	for i := 0; i < 100; i++ {
		_user := User{
			Name: fmt.Sprintf("name-%d", i),
			Age:  i,
		}
		users = append(users, _user)
	}
}

// for range取址问题
func fn1() {
	s := make([]*User, len(users))
	for _, user := range users {
		log.Printf("user p := %p", &user)
		s = append(s, &user)
	}

	log.Println(s)
}

func main() {
	fn1()
}
