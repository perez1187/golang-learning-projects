package main

import (
	"fmt"
	"time"
)

func main() {
	userch := make(chan string)

	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-userch:
		}

	}
}

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{} // because struct(empty struct) has 0 bytes
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {

running: // we can name for loop
	for {
		select {
		case msg := <-s.userch:
			fmt.Println(msg)
		case <-s.quitch:
			fmt.Println("Server needs to quit")
			break running // this is the same like return
			// but without name eg running the program "hang"
			// return
		default:
		}
	}
}

// what normally user do:
func (s *Server) addUser(user string) {
	s.users[user] = user
}
