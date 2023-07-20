// message communication with For Select
//59:00

package main

import (
	"fmt"
	"time"
)

// he use this patterm all the time
type Message struct {
	From    string
	Payload string // normaly it will be slice of bytes
}

type Server struct {
	msgch  chan Message
	quitch chan struct{} // zero memory allocation
}

func (s *Server) StartAndListen() {

free:
	for {
		select {

		case msg := <-s.msgch:
			fmt.Printf("receiving meesage from: %s payload: %s \n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a gracefull shutdown")
			// logic for the server shutdown
			break free // without name, is break for this case statement
		default:
		}

	}
	fmt.Println("the server is shut down")
}

func graceFullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Joy",
		Payload: payload,
	}

	msgch <- msg
}

func main() {

	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "yo yo")
	}() // this is what we return (?)

	go func() {
		time.Sleep(5 * time.Second)
		graceFullQuitServer(s.quitch)
	}()

	select {} // never use it on production
}
